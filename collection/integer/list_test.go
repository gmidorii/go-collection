package integer

import "testing"

var filterDatas = []struct {
	in  List
	fn  func(interface{}) bool
	out List
}{
	{[]int{0, 1, 2, 3, 4, 5, 6, 7}, func(i interface{}) bool { return i.(int) > 3 }, []int{4, 5, 6, 7}},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7}, func(i interface{}) bool { return i.(int) <= 3 }, []int{0, 1, 2, 3}},
}

var collectDatas = []struct {
	in  List
	fn  func(interface{}) interface{}
	out List
}{
	{[]int{0, 1, 2, 3}, func(i interface{}) interface{} { return i.(int) * 10 }, []int{0, 10, 20, 30}},
	{[]int{0, 1, 2, 3}, func(i interface{}) interface{} { return i.(int) + 10 }, []int{10, 11, 12, 13}},
}

func TestFilter(t *testing.T) {
	for _, v := range filterDatas {
		actual := v.in.Filter(v.fn)
		for i, a := range actual.(List) {
			if v.out[i] != a {
				t.Fatalf("error not expected value e: %d, a: %d", v.out[i], a)
			}
		}
	}
}

func TestCollect(t *testing.T) {
	for _, v := range collectDatas {
		actual := v.in.Collect(v.fn)
		for i, a := range actual.(List) {
			if v.out[i] != a {
				t.Fatalf("error not expected value e: %d, a: %d", v.out[i], a)
			}
		}
	}
}
