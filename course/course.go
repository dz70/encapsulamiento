package main

import "fmt"

// Course es una estructura
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIds []uint
	Classes map[uint]string
}

// ChangePrice es un metodo para estructura course
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

// PrintClasses es un metodo para estructura course
func (c Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
