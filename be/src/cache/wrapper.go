package cache

import (
	"errors"
	"fmt"
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
		// 如果不是pointer會panic
		ptr.Elem().Set(rv.Elem())
		return nil
	}
	fmt.Printf("\nLoad '%s' from db \n", key)
	v, err := db(key)
	if err != nil {
		return err
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		ptr.Elem().Set(rv)

		// 轉成pointer
		ptrv := reflect.New(rv.Type())
		ptrv.Elem().Set(rv)
		s.Set(key, ptrv.Interface())
	} else {
		ptr.Elem().Set(rv.Elem())
		s.Set(key, v)
	}
	return nil
}
