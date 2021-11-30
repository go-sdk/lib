package mapx

type Map interface {
	Put(key, value interface{}, ex ...bool)
	Get(key interface{}) (interface{}, bool)
	ContainsKey(key interface{}) bool
	Remove(key interface{}) (interface{}, bool)
	Removes(keys ...interface{})
	Replace(key, value interface{}, nx ...bool) (interface{}, bool)
	Keys() []interface{}
	Values() []interface{}
	ForEach(f func(key, value interface{}) bool)

	Size() int
	IsEmpty() bool
	Clear()
	Copy() Map
}
