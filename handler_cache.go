package handler_cache

type Cache struct {
	key   string
	value any
}

var items = make(map[string]any)

func New() Cache {
	return Cache{
		key: "",
	}
}

func (c *Cache) Set(key string, value any) {
	items[key] = value
}

func (c Cache) Get(key string) any {
	return items[key]
}

func (c *Cache) Delete(key string) {
	delete(items, key)
}
