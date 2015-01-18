package fn

type Any interface{}
type Mapper func(Any) Any
type Predicate func(Any) bool
type Reducer func(Any, Any) Any

type Iterable interface {
	Next() Any
	HasNext() bool
	Length() int
	Reset()
	AsArray() []Any
}

type Seq interface {
	Iterable
	Map(Mapper) Seq
	ForEach(func(Any))
	Find(Predicate) Any
	FindAll(Predicate) Seq
	Reduce(Any) func(Reducer) Any
	//	Flatten() Seq
}
