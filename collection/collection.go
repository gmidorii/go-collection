package collection

type Collection interface {
	Filter(func(interface{}) bool) Collection
	Each(func(int, int))
	Collect(func(interface{}) interface{}) Collection
}
