package integer

type Collection interface {
	Filter(func(interface{}) bool) Collection
	Each(func(int, int))
	Map(func(interface{}) interface{}) Collection
}
