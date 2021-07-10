package cache

import (
	"container/list"
	"reflect"
	"sync"
)

type Store struct {
	shadeMask uint64
	shards    []*shard
	queue     *list.List
	lock      *sync.RWMutex
}

func New(sm, lm uint64) *Store {
	sm = 1 << sm
	lm = 1 << lm
	r := &Store{
		shadeMask: sm - 1,
		shards:    make([]*shard, sm),
		queue:     list.New(),
		lock:      &sync.RWMutex{},
	}
	for i := range r.shards {
		r.shards[i] = newShard(int(lm))
	}
	return r
}

func (s *Store) Set(key string, value interface{}) {
	h := hash(key)
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Ptr {
		// 轉成pointer
		ptrv := reflect.New(rv.Type())
		ptrv.Elem().Set(rv)
		s.shards[h&(s.shadeMask)].set(h, ptrv.Interface())
		return
	}
	s.shards[h&(s.shadeMask)].set(h, value)
}

func (s *Store) Get(key string) (interface{}, bool) {
	h := hash(key)
	return s.shards[h&s.shadeMask].get(h)
}

func (s *Store) Del(key string) bool {
	h := hash(key)
	return s.shards[h&s.shadeMask].del(h)
}

func (s *Store) Len() int {
	r := 0
	for _, v := range s.shards {
		r += v.len()
	}
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
