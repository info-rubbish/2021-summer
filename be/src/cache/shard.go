package cache

import (
	"container/list"
	"sync"
)

type shard struct {
	max  int
	p    int
	kv   map[uint64]uint64
	list []*list.Element
	lru  *list.List
	lock *sync.RWMutex
}

type element struct {
	value interface{}
	hash  uint64
}

func newShard(m int) *shard {
	return &shard{
		max:  m,
		p:    0,
		kv:   make(map[uint64]uint64, m),
		list: make([]*list.Element, m),
		lru:  list.New(),
		lock: &sync.RWMutex{},
	}
}

func (s *shard) set(keyHash uint64, value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if v, ok := s.kv[keyHash]; ok {
		s.list[v].Value.(*element).value = value
		s.lru.MoveToBack(s.list[v])
		return
	}
	if s.lru.Len() >= s.max {
		s._delete(s.lru.Front())
	}
	l := s._findEmpty()
	s.list[l] = s.lru.PushBack(&element{
		value: value,
		hash:  keyHash,
	})
	s.kv[keyHash] = l
}

func (s *shard) get(keyHash uint64) (interface{}, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	v, ok := s.kv[keyHash]
	if !ok {
		return nil, false
	}
	s.lru.MoveToBack(s.list[v])
	return s.list[v].Value.(*element).value, true

}

func (s *shard) del(keyHash uint64) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	v, ok := s.kv[keyHash]
	if !ok {
		return false
	}
	s._delete(s.list[v])
	return true
}

func (s *shard) len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.lru.Len()
}

func (s *shard) _delete(e *list.Element) {
	s.lru.Remove(e)
	s.list[s.kv[e.Value.(*element).hash]] = nil
	delete(s.kv, e.Value.(*element).hash)
}

func (s *shard) _findEmpty() uint64 {
	for ; s.p < s.max; s.p++ {
		if s.list[s.p] == nil {
			return uint64(s.p)
		}
	}
	s.p = 0
	return s._findEmpty()
}
