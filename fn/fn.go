package fn

type Any interface{}
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
	Append(Any) Seq
	//	AppendAll(Seq) Seq
	Map(Mapper) Seq
	ForEach(func(Any))
	Find(Predicate) Any
	Filter(Predicate) Seq
	Reduce(Any) func(Reducer) Any
	Flatten() Seq
}
