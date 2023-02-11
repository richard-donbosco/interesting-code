package main

import "fmt"

const (
	a int = iota
	b
	c
	d
	_
	f
)

func main() {
	fmt.Printf("%v, %v, %v, %v, %v", a, b, c, d, f)
}
