package main

import "fmt"

type Cache struct {
	Data map[string]int
	Keys []string
	Max  int
}

func (c *Cache) set(k string, v int) {
	if c.Data == nil {
		c.Data = make(map[string]int)
	}
	//fmt.Println(len(c.Keys), c.Max)
	if len(c.Keys) >= c.Max {
		delete(c.Data, c.Keys[0])
		c.Keys[0] = ""
	}
	c.Data[k] = v
	c.Keys = append(c.Keys, k)
}

func (c *Cache) get(k string) int {
	v := c.Data[k]
	index := 0
	for i := 0; i < len(c.Keys); i++ {
		if c.Keys[i] == k {
			index = i
		}
	}
	c.Keys[index] = ""
	c.Keys = append(c.Keys, k)
	return v
}

func main() {
	c := Cache{
		Max: 2,
	}
	c.set("a", 1)
	c.set("b", 2)
	c.set("c", 3)

	fmt.Println(c.get("a"))
	fmt.Println(c.get("b"))
	fmt.Println(c.get("c"))
}
