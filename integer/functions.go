package integer

import "fmt"

type FilterFunc struct {
	Comparison int
}

func (f *FilterFunc) Larger(i interface{}) bool {
	return i.(int) > f.Comparison
}

func (f *FilterFunc) Smaller(i interface{}) bool {
	return i.(int) < f.Comparison
}

type EachFunc struct {
	PrintFmt string
}

func (f *EachFunc) Println(k, v int) {
	fmt.Printf(f.PrintFmt, k, v)
}
