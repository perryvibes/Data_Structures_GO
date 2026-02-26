package main

import "fmt"

type Clothing struct {
	sku   int
	stock int
	name  string
	price float32
	size  rune
}

func ModifyAttribute(c *Clothing) {
	if c.stock < 20 {
		c.price *= 1.25
	}
}

func main() {

	jacket := Clothing{421, 16, "Leather Jacket", 120.4, 'L'}
	fmt.Println(jacket)
	ModifyAttribute(&jacket)
	fmt.Println(jacket)

}
