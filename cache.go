package cacheUdemy

type Cache struct {
	data map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) SetKey(key string, value interface{}) {
	c.data[key] = value
}

func (c Cache) GetKey(key string) interface{} {
	return c.data[key]
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
