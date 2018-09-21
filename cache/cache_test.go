package cache

import (
	"testing"
	"time"
)

//
type SomeStruct struct {
	Id uint64
}

//
func TestCacheGetSet(t *testing.T) {
	c := MemoryCacheNew()
	c.Set("1000", &SomeStruct{1000})
	c.Set("2000", &SomeStruct{2000})
	c.Set("3000", &SomeStruct{3000})

	if !c.Has("1000") {
		t.Errorf("Not exists 1000 key")
	}
	if !c.Has("2000") {
		t.Errorf("Not exists 2000 key")
	}
	if !c.Has("3000") {
		t.Errorf("Not exists 3000 key")
	}
	if c.Has("4000") {
		t.Errorf("Not exists 4000 key")
	}

	p := c.Get("1000")
	if p == nil {
		t.Errorf("P is not nil")
	}

	if s, ok := p.(*SomeStruct); ok {
		if s.Id != 1000 {
			t.Errorf("1000 != ID struct ")
		}
	} else {
		t.Errorf("Error convert pointer to SomeStruct{}")
	}
}

//
func TestCacheGetSetExp(t *testing.T) {
	c := MemoryCacheNew()
	c.Set("1000", &SomeStruct{1000})
	c.SetExp("2000", &SomeStruct{2000}, 3)

	if !c.Has("1000") {
		t.Errorf("Not exists 1000 key")
	}
	if !c.Has("2000") {
		t.Errorf("Not exists 2000 key")
	}

	time.Sleep(time.Second * 5)

	if !c.Has("1000") {
		t.Errorf("Not exists 1000 key")
	}
	if c.Has("2000") {
		t.Errorf("Exists 2000 key")
	}
}
