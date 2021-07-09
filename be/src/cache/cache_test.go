package cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestKeyValue(t *testing.T) {
	s := New(2, 5)
	s.Set("1", 2)
	v, ok := s.Get("1")
	if !ok {
		t.Fatal("should exist")
	}
	t.Log(v)
}

func TestNil(t *testing.T) {
	s := New(2, 5)
	s.Set("1", 2)
	s.Del("1")
	s.Set("1", 2)
	s.Del("2")
	s.Set("1", 2)
}

func TestFill(t *testing.T) {
	s := New(3, 5)
	for i := 0; i < 10000; i++ {
		s.Set(strconv.Itoa(i), i+1)
	}
	t.Log(s.Len())
}

func TestGC(t *testing.T) {
	s := New(2, 5)
	for i := 0; i < 17; i++ {
		s.Set(strconv.Itoa(i), i+1)
	}
	t.Log(s.Len())
	for i := 0; i < 17; i++ {
		if v, ok := s.Get(strconv.Itoa(i)); ok && v != i+1 {
			panic(fmt.Sprintf("should=i+1 (%d!=%d+1)", v, i))
		}
	}
	if _, ok := s.Get(strconv.Itoa(1)); !ok {
		t.Fatal("should exist")
	}
	s.Del(strconv.Itoa(1))
	if _, ok := s.Get(strconv.Itoa(1)); ok {
		t.Fatal("should not exist")
	}
}

func TestDataRace(t *testing.T) {
	s := New(5, 10)
	for i := 0; i < 10000; i++ {
		go func(i int) {
			s.Set(strconv.Itoa(i), i+1)
		}(i)
		go func(i int) {
			s.Get(strconv.Itoa(i))
		}(i)
		go func(i int) {
			s.Del(strconv.Itoa(i))
		}(i)
	}
	time.Sleep(3 * time.Second)
}

func TestLen(t *testing.T) {
	s := New(1, 10)
	for i := 0; i < 3000; i++ {
		go func(i int) {
			s.Set(strconv.Itoa(i), i+1)
		}(i)
		// go func(i int) {
		// 	s.Get([]byte(strconv.Itoa(i)))
		// 	// if ok && v != i+1 {
		// 	// 	panic(fmt.Sprintf("should=i+1 (%d!=%d+1)", v, i))
		// 	// }
		// }(i)
	}
	time.Sleep(3 * time.Second)
	t.Log(s.Len())
	if s.Len() > int(1<<11) {
		t.Fatalf("should not larger than max (%d)", s.Len())
	}
}

// old 1000000	      1721 ns/op	     187 B/op	       3 allocs/op
// new 3502290	     363.1 ns/op	      31 B/op	       3 allocs/op
// end 1000000	      1261 ns/op	     175 B/op	       4 allocs/op
// sha 1330687	     905.5 ns/op	      90 B/op	       4 allocs/op
func BenchmarkDataW(b *testing.B) {
	s := New(5, 10)
	for i := 0; i < b.N; i++ {
		s.Set(strconv.Itoa(i), i+1)
	}
}

// old 12986787	        99.95 ns/op	       0 B/op	       0 allocs/op
// new 5215490 	        222.2 ns/op	      16 B/op	       1 allocs/op
// end 14813652	        80.06 ns/op	       0 B/op	       0 allocs/op
// sha 15560730	        77.58 ns/op	       0 B/op	       0 allocs/op
func BenchmarkDataR(b *testing.B) {
	s := New(5, 10)
	for i := 0; i < 1000; i++ {
		s.Set(strconv.Itoa(i), i+1)
	}
	for i := 0; i < b.N; i++ {
		s.Get(strconv.Itoa(20))
	}
}

// old 1000000	      2194 ns/op	     187 B/op	       3 allocs/op
// new  928842	      1608 ns/op	     136 B/op	       8 allocs/op
// end 1250223	       922.2 ns/op	     105 B/op	       5 allocs/op
// sha 1227085	       892.1 ns/op	     105 B/op	       5 allocs/op
func BenchmarkDataRWD(b *testing.B) {
	s := New(5, 10)
	for i := 0; i < b.N; i++ {
		s.Set(strconv.Itoa(i), i+1)
		s.Get(strconv.Itoa(i))
		s.Del(strconv.Itoa(i))
	}
}

func BenchmarkRandDataRWD(b *testing.B) {
	s := New(5, 10)
	for i := 1; i < b.N; i++ {
		r := rand.Intn(i)
		d := rand.Intn(i)
		// println(r, ":", d)
		s.Set(strconv.Itoa(r), 0)
		s.Get(strconv.Itoa(rand.Intn(i)))
		s.Del(strconv.Itoa(d))
	}
}
