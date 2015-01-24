package fn

type Any interface{}
type None interface{}
type Mapper func(Any) Any
type Predicate func(Any) bool
type Reducer func(Any, Any) Any

type Iterable interface {
	Next() Any
	HasNext() bool
}

type Traversable interface {
	Iterable
	Map(Mapper) Traversable
	ForEach(func(Any))
	Find(Predicate) Any
	Filter(Predicate) Traversable
	Reduce(Any) func(Reducer) Any
	Flatten() Traversable
}
