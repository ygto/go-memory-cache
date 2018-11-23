package memory

import (
	"sync"
)

type memoryCache struct {
	Data map[string]string
	mux  sync.Mutex
}

func (c *memoryCache) Set(key string, val string) {
	defer c.mux.Unlock()
	c.mux.Lock()
	c.Data[key] = val
}
func (c *memoryCache) Get(key string) (string, bool) {
	defer c.mux.Unlock()
	c.mux.Lock()
	val, ok := c.Data[key]
	return val, ok
}

func (c *memoryCache) Del(key string) {
	defer c.mux.Unlock()
	c.mux.Lock()
	delete(c.Data,key)
}
func NewCache() *memoryCache {
	c := memoryCache{}
	c.Data = make(map[string]string)
	return &c
}
/*
func main() {

	c := NewCache()
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			key := fmt.Sprintf("name_%d", i)
			val := fmt.Sprintf("%d", i)
			c.Set(key, val)
			wg.Done()
		}(i)
	}
	wg.Wait()
	wg = sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			key := fmt.Sprintf("name_%d", i)
			if data, ok := c.Get(key); ok {
				fmt.Println(data)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}*/
