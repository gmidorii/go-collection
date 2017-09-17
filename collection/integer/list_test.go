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
