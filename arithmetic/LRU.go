package arithmetic

const length = 5

type LRUCache struct {
	Map       map[int]int
	MaxLength int
}

func NewLRUCache(l int) *LRUCache {
	return &LRUCache{Map: make(map[int]int), MaxLength: l}
}

func (c *LRUCache) Put(key, value int) {

}
func (c *LRUCache) Get(key int) int {
	if _, ok := c.Map[key]; !ok {
		return -1
	}

	return c.Map[key]
}
