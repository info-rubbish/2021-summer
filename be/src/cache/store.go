package cache

import (
	"container/list"
	"sync"
)

type Store struct {
	shadeMax uint64
	Max      int
	shades   []*shade
	list     []*list.Element
	lru      *list.List
	queue    *list.List
	p        uint64
	lock     *sync.RWMutex
}

type shade struct {
	kv   map[uint64]uint64
	lock *sync.RWMutex
}

type element struct {
	value interface{}
	hash  uint64
}

func New(sm, lm uint64) *Store {
	sm = 1 << sm
	lm = 1 << lm
	r := &Store{
		shadeMax: sm,
		Max:      int(sm * lm),
		shades:   make([]*shade, sm),
		list:     make([]*list.Element, sm*lm),
		lru:      list.New(),
		queue:    list.New(),
		p:        0,
		lock:     &sync.RWMutex{},
	}
	for i := range r.shades {
		r.shades[i] = &shade{
			kv:   make(map[uint64]uint64, lm),
			lock: &sync.RWMutex{},
		}
	}
	return r
}

func (s *Store) Set(key string, value interface{}) {
	keyHash := hash(key)
	shadeHash := keyHash & (s.shadeMax - 1)
	s.lock.Lock()
	s.shades[shadeHash].lock.Lock()
	if v, ok := s.shades[shadeHash].kv[keyHash]; ok {
		// s.lock.Lock()
		s.list[v].Value.(*element).value = value
		s.lru.MoveToBack(s.list[v])
		s.lock.Unlock()
		s.shades[shadeHash].lock.Unlock()
		return
	}
	if s.lru.Len() >= s.Max {
		front := s.lru.Front()
		s.lru.Remove(front)
		frontHash := front.Value.(*element).hash
		frontShadeHash := frontHash & (s.shadeMax - 1)
		if shadeHash != frontShadeHash {
			s.shades[frontShadeHash].lock.Lock()
			s.list[s.shades[frontShadeHash].kv[frontHash]] = nil
			delete(s.shades[frontShadeHash].kv, frontHash)
			s.shades[frontShadeHash].lock.Unlock()
		} else {
			s.list[s.shades[frontShadeHash].kv[frontHash]] = nil
			delete(s.shades[frontShadeHash].kv, frontHash)
		}
	}
	l := s.findEmpty()
	s.shades[shadeHash].kv[keyHash] = l
	s.list[l] = s.lru.PushBack(&element{
		value: value,
		hash:  keyHash,
	})
	s.shades[shadeHash].lock.Unlock()
	s.lock.Unlock()
}

func (s *Store) Get(key string) (interface{}, bool) {
	keyHash := hash(key)
	shadeHash := keyHash & (s.shadeMax - 1)
	s.lock.Lock()
	s.shades[shadeHash].lock.RLock()
	offset, ok := s.shades[shadeHash].kv[keyHash]
	if !ok {
		s.shades[shadeHash].lock.RUnlock()
		s.lock.Unlock()
		return nil, false
	}
	s.lru.MoveToBack(s.list[offset])
	r := s.list[offset].Value.(*element).value
	s.shades[shadeHash].lock.RUnlock()
	s.lock.Unlock()
	return r, true
}

func (s *Store) Del(key string) bool {
	keyHash := hash(key)
	shadeHash := keyHash & (s.shadeMax - 1)
	s.lock.Lock()
	s.shades[shadeHash].lock.Lock()
	v, ok := s.shades[shadeHash].kv[keyHash]
	if !ok {
		s.shades[shadeHash].lock.Unlock()
		s.lock.Unlock()
		return false
	}
	delete(s.shades[shadeHash].kv, keyHash)
	s.lru.Remove(s.list[v])
	s.list[v] = nil
	// reuse empty
	s.queue.PushBack(v)
	s.shades[shadeHash].lock.Unlock()
	s.lock.Unlock()
	return true
}

func (s *Store) Len() int {
	s.lock.RLock()
	r := s.lru.Len()
	s.lock.RUnlock()
	return r
}

// from bigcache
const (
	// offset64 FNVa offset basis. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	offset64 = 14695981039346656037
	// prime64 FNVa prime value. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	prime64 = 1099511628211
)

func hash(key string) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= prime64
	}
	return hash
}

func (s *Store) findEmpty() uint64 {
	if s.queue.Len() > 0 {
		f := s.queue.Front().Value.(uint64)
		s.queue.Remove(s.queue.Front())
		return f
	}
	for ; s.p < uint64(s.Max); s.p++ {
		if s.list[s.p] == nil {
			return s.p
		}
	}
	s.p = 0
	return s.findEmpty()
}

func (s *shade) set(key string, value interface{}) {

}

// type Store struct {
// 	shadeMax uint64
// 	listMax  uint64
// 	shades   []*shade
// 	lru      *list.List
// 	lock     *sync.RWMutex
// }

// type shade struct {
// 	list []*list.Element
// 	lock *sync.RWMutex
// }

// type element struct {
// 	value interface{}
// 	hash  uint64
// }

// // m is max length of cache
// func New(sm, lm uint64) *Store {
// 	sm = 1 << sm
// 	lm = 1 << lm
// 	r := &Store{
// 		shadeMax: sm,
// 		listMax:  lm,
// 		shades:   make([]*shade, sm),
// 		lru:      list.New(),
// 		lock:     &sync.RWMutex{},
// 	}
// 	for i := range r.shades {
// 		r.shades[i] = &shade{
// 			list: make([]*list.Element, lm),
// 			lock: &sync.RWMutex{},
// 		}
// 	}
// 	return r
// }

// func (s *Store) Set(key string, value interface{}) bool {
// 	h := hashString(key)
// 	r, c := s.find(h)
// 	shade := s.shades[r]
// 	shade.lock.Lock()
// 	defer shade.lock.Unlock()
// 	s.lock.Lock()
// 	defer s.lock.Unlock()
// 	item := shade.list[c]
// 	if item != nil {
// 		// f := s.lru.Front()
// 		// s.lru.Remove(f)
// 		// fr, fc := s.find(f.Value.(*element).hash)
// 		// if fr != r {
// 		// 	s.shades[fr].lock.Lock()
// 		// 	s.shades[fr].list[fc] = nil
// 		// 	s.shades[fr].lock.Unlock()
// 		// } else {
// 		// 	s.shades[fr].list[fc] = nil
// 		// }
// 		return false
// 	}
// 	if s.lru.Len() >= int(s.shadeMax*s.listMax) {
// 		f := s.lru.Front()
// 		s.lru.Remove(f)
// 		fr, fc := s.find(f.Value.(*element).hash)
// 		if fr != r {
// 			s.shades[r].lock.Lock()
// 			s.shades[fr].list[fc] = nil
// 			s.shades[r].lock.Unlock()
// 		} else {
// 			s.shades[fr].list[fc] = nil
// 		}
// 	}
// 	shade.list[c] = s.lru.PushBack(&element{
// 		value: value,
// 		hash:  h,
// 	})
// 	return true
// 	// s.lock.Lock()
// 	// defer s.lock.Unlock()
// 	// if check, ok := s.kv[key]; ok {
// 	// 	check.Value = value
// 	// 	s.list.MoveToBack(check)
// 	// 	return
// 	// }
// 	// if s.list.Len() >= s.max {
// 	// 	s.list.Remove(s.list.Front())
// 	// }
// 	// s.kv[key] = s.list.PushBack(value)
// }

// func (s *Store) Put(key string, value interface{}) bool {
// 	h := hashString(key)
// 	r, c := s.find(h)
// 	shade := s.shades[r]
// 	shade.lock.Lock()
// 	defer shade.lock.Unlock()
// 	s.lock.Lock()
// 	defer s.lock.Unlock()
// 	item := shade.list[c]
// 	if item == nil {

// 		return false
// 	}
// 	item.Value.(*element).value = value
// 	s.lru.MoveToBack(item)
// 	return true
// }

// func (s *Store) Get(key string) (interface{}, bool) {
// 	h := hashString(key)
// 	r, c := s.find(h)
// 	shade := s.shades[r]
// 	shade.lock.RLock()
// 	defer shade.lock.RUnlock()
// 	item := shade.list[c]
// 	if item != nil && item.Value.(*element).hash == h {
// 		s.lock.Lock()
// 		defer s.lock.Unlock()
// 		s.lru.MoveToBack(item)
// 		return item.Value.(*element).value, true
// 	}
// 	return nil, false
// 	// s.lock.Lock()
// 	// defer s.lock.Unlock()
// 	// value, ok := s.kv[key]
// 	// if !ok {
// 	// 	return nil, false
// 	// }
// 	// s.list.MoveToBack(value)
// 	// return value.Value, true
// }

// func (s *Store) Del(key string) bool {
// 	r, c := s.find(hashString(key))
// 	shade := s.shades[r]
// 	shade.lock.Lock()
// 	defer shade.lock.Unlock()
// 	item := shade.list[c]
// 	if item != nil {
// 		s.lock.Lock()
// 		defer s.lock.Unlock()
// 		s.lru.Remove(item)
// 		shade.list[c] = nil
// 		return true
// 	}
// 	return false
// 	// s.lock.Lock()
// 	// defer s.lock.Unlock()
// 	// value, ok := s.kv[key]
// 	// if !ok {
// 	// 	return false
// 	// }
// 	// s.list.Remove(value)
// 	// return true
// }

// func (s *Store) DelOldest() {
// 	s.lock.Lock()
// 	defer s.lock.Unlock()
// 	f := s.lru.Front()
// 	if f == nil {
// 		return
// 	}
// 	s.lru.Remove(f)
// 	r, c := s.find(f.Value.(*element).hash)
// 	s.shades[r].lock.Lock()
// 	defer s.shades[r].lock.Unlock()
// 	s.shades[r].list[c] = nil
// }

// func (s *Store) Len() int {
// 	s.lock.RLock()
// 	defer s.lock.RUnlock()
// 	return s.lru.Len()
// }

// func hashByteArray(b []byte) uint32 {
// 	h := fnv.New32a()
// 	h.Write(b)
// 	return h.Sum32()
// }

// func hashUint32(u uint32) uint32 {
// 	h := fnv.New32a()
// 	h.Write([]byte(strconv.Itoa(int(u))))
// 	return h.Sum32()
// }

// func (s *Store) find(h uint64) (uint64, uint64) {
// 	return h & (s.shadeMax - 1), hashString(strconv.Itoa(int(h))) & (s.listMax - 1)
// }

// // func (s *Store)delete(r,c uint32) {

// // }
