package cache

import (
	"errors"
	"reflect"
)

var (
	ErrMustBePointer = errors.New("must be pointer")
)

type Cache interface {
	Set(string, interface{})
	Get(string) (interface{}, bool)
	Del(string) bool
}

type FromDB func(string) (interface{}, error)
type Wrapper struct {
	Cache
}

func NewWrapper(v Cache) *Wrapper {
	return &Wrapper{
		v,
	}
}

// pointer must be a pointer
func (s *Wrapper) Load(pointer interface{}, key string, db FromDB) error {
	ptr := reflect.ValueOf(pointer)
	if ptr.Kind() != reflect.Ptr {
		return ErrMustBePointer
	}
	if v, ok := s.Get(key); ok {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Ptr {
			ptr.Elem().Set(rv)
		} else {
			ptr.Elem().Set(rv.Elem())
		}
		return nil
	}
	v, err := db(key)
	if err != nil {
		return err
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		ptr.Elem().Set(rv)
	} else {
		ptr.Elem().Set(rv.Elem())
	}
	s.Set(key, v)
	return nil
}
