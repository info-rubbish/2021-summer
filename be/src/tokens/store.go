package tokens

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrNotFind error = errors.New("key not find in store")
	ErrExpired error = errors.New("key Expired")
)

type Store struct {
	store *sync.Map
	n     uint
	ttl   time.Duration
	close chan struct{}
}

type token struct {
	expired time.Time
	data    interface{}
}

// ttl token time to live
// gct gc time
// n token len
func NewStore(ttl time.Duration, gct time.Duration, n uint) *Store {
	s := &Store{
		store: &sync.Map{},
		n:     n,
		ttl:   ttl,
		close: make(chan struct{}),
	}
	go s.gc(gct)
	return s
}

func (s *Store) Close() {
	close(s.close)
	s = nil
}

// data must be a pointer
// go gc lagggggg
func (s *Store) NewToken(data interface{}) (string, error) {
	tokenString, err := GenerateToken(s.n)
	if err != nil {
		return "", err
	}
	s.store.Store(tokenString, &token{
		expired: time.Now().Add(s.ttl),
		data:    data,
	})
	return tokenString, nil
}

func (s *Store) DestroyToken(tokenKey interface{}) error {
	_, ok := s.store.LoadAndDelete(tokenKey)
	if !ok {
		return ErrNotFind
	}
	return nil
}

// return data
func (s *Store) GetToken(tokenKey interface{}) (interface{}, error) {
	value, ok := s.store.Load(tokenKey)
	if !ok {
		return nil, ErrNotFind
	}
	if value.(*token).expired.Before(time.Now()) {
		s.DestroyToken(tokenKey)
		return nil, ErrExpired
	}
	return value.(*token).data, nil
}

// return new token
func (s *Store) ReNewToken(tokenKey interface{}) (string, error) {
	value, ok := s.store.LoadAndDelete(tokenKey)
	if !ok {
		return "", ErrNotFind
	}
	if value.(*token).expired.Before(time.Now()) {
		return "", ErrExpired
	}
	newTokenString, err := s.NewToken(value.(*token).data)
	if err != nil {
		return "", err
	}
	return newTokenString, nil
}

func (s *Store) gc(g time.Duration) {
	tick := time.NewTicker(g)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			expired := time.Now().Add(s.ttl)
			s.store.Range(func(key, value interface{}) bool {
				if value.(*token).expired.Before(expired) {
					s.DestroyToken(key)
				}
				return true
			})
		case <-s.close:
			return
		}
	}
}
