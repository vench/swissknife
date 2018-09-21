package cache

import (
	"time"
)

//
type MemoryCacheItem struct {
	createTime      time.Time
	secondsDuration float64
	Data            interface{}
}

//
type MemoryCache struct {
	mapCache map[string]*MemoryCacheItem
}

//
func MemoryCacheNew() *MemoryCache {
	item := MemoryCache{}
	item.mapCache = make(map[string]*MemoryCacheItem)
	return &item
}

//
func (m *MemoryCache) Get(key string) interface{} {
	if item, ok := m.mapCache[key]; ok {
		if item.secondsDuration == 0 ||
			time.Now().Sub(item.createTime).Seconds() < item.secondsDuration {
			return item.Data
		}
		delete(m.mapCache, key)
	}
	return nil
}

//
func (m *MemoryCache) Has(key string) bool {
	if m.Get(key) == nil {
		return false
	}
	return true
}

//
func (m *MemoryCache) Set(key string, data interface{}) {
	if !m.Has(key) {
		item := MemoryCacheItem{}
		item.Data = data
		m.mapCache[key] = &item
	}
}

//
func (m *MemoryCache) SetExp(key string, data interface{}, seconds float64) {
	if !m.Has(key) {
		item := MemoryCacheItem{time.Now(), seconds, data}
		m.mapCache[key] = &item
	} else {
		m.mapCache[key].createTime = time.Now()
		m.mapCache[key].secondsDuration = seconds
	}

}
