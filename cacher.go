package cacher

type Cache struct {
	storage map[string]interface{}
}

func New() Cache {
	return Cache{make(map[string]interface{})}
}

func (c Cache) Get(key string) (interface{}, bool) {
	value, ok := c.storage[key]
	return value, ok
}

func (c *Cache) Set(key string, val interface{}) {
	c.storage[key] = val
}

func (c *Cache) Delete(key string) {
	delete(c.storage, key)
}
