package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	ci := &cacheItem{
		key:   key,
		value: value,
	}
	_, found := cache.items[key]
	if !found {
		cache.items[key] = cache.queue.PushFront(ci)
		if cache.queue.Len() > cache.capacity {
			delete(cache.items, cache.queue.Back().Value.(*cacheItem).key)
			cache.queue.Remove(cache.queue.Back())
		}

		return false
	}

	cache.items[key].Value.(*cacheItem).value = value
	cache.queue.MoveToFront(cache.items[key])

	return true
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	_, found := cache.items[key]

	if found {
		cache.queue.MoveToFront(cache.items[key])
		return cache.items[key].Value.(*cacheItem).value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
