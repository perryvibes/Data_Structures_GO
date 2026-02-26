package main

import "fmt"

type Clothing struct {
	sku   int
	stock int
	name  string
	price float32
	size  rune
}

func Initiate(_sku int, _stock int, _name string, _price float32, _size rune) Clothing {
	var c Clothing
	c.sku = _sku
	c.stock = _stock
	c.name = _name
	c.price = _price
	c.size = _size
	return c
}

func ModifyAttribute(c *Clothing) {
	if c.stock < 20 {
		c.price *= 1.25
	}
}

func main() {

	var jacket Clothing = Initiate(421, 16, "Leather Jacket", 120.4, 'L')
	fmt.Println(jacket)
	ModifyAttribute(&jacket)
	fmt.Println(jacket)

}
