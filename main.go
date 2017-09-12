package main

import (
	"fmt"

	"github.com/midorigreen/go-collection/integer"
)

func main() {
	sampleList()
}

func sampleList() {
	list := integer.CreateList(5)
	for i := 0; i < 5; i++ {
		list = append(list, i)
	}
	// Filter
	filtered := list.Filter(func(i interface{}) bool {
		return i.(int) > 2
	})
	fmt.Printf("%v\n", filtered)

	// Each
	list.Each(func(i, v int) {
		fmt.Printf("index: %d, value: %d\n", i, v)
	})

	// Map
	mapped := list.Map(func(i interface{}) interface{} {
		return i.(int) * 10
	})
	fmt.Printf("%v\n", mapped)
}
