package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")
	a := 10
	b := &a
	c := b
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", &a)
	fmt.Printf("%v",*c)

}
