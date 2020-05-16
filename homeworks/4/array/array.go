package array

type Array interface {
	Add(item interface{})
	Set(item interface{}, index int) bool
	Get(index int) interface{}
	GetExisted(index int) (interface{}, bool)
	Len() int
	Cap() int
	Remove(index int) (interface{}, bool)
}
