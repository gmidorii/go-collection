package main

import (
	"fmt"

	"github.com/midorigreen/go-collection/collection"
	"github.com/midorigreen/go-collection/collection/integer"
)

func main() {
	sampleList()
	sampleFuncList()
}

func sampleList() {
	fmt.Println("-- sampleList --")
	var list collection.Collection
	list = integer.CreateList(5)
	for i := 0; i < 5; i++ {
		list = append(list.(integer.List), i)
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

	// Collect
	mapped := list.Collect(func(i interface{}) interface{} {
		return i.(int) * 10
	})
	fmt.Printf("%v\n", mapped)
}

func sampleFuncList() {
	fmt.Println("-- sampleFuncList --")
	var list collection.Collection
	list = integer.CreateList(5)
	for i := 0; i < 5; i++ {
		list = append(list.(integer.List), i)
	}
	// Filter
	fFn := integer.FilterFunc{Comparison: 2}
	filtered := list.Filter(fFn.Larger)
	fmt.Printf("%v\n", filtered)

	// Each
	eFn := integer.EachFunc{PrintFmt: "index: %d, value: %d\n"}
	list.Each(eFn.Println)
}
