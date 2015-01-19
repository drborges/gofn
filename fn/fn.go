package fn

type Any interface{}
type None interface{}
type Mapper func(Any) Any
type Predicate func(Any) bool
type Reducer func(Any, Any) Any

type Iterable interface {
	Next() Any
	HasNext() bool
	Length() int
	Reset() // FIXME An Iterable should not expose this method
	AsArray() []Any
}

type Seq interface {
	Head() Any
	Tail() Seq
	Append(Any) Seq
	Map(Mapper) Seq
	ForEach(func(Any))
	Find(Predicate) Any
	Filter(Predicate) Seq
	Reduce(Any) func(Reducer) Any
	Flatten() Seq
}

//type Tree interface {}
//type Map map[Any]Any
//type Set []Any
