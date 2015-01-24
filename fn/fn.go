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

type Seq interface {
	Head() Any
	Tail() Seq
	Iter() Iterable
	Append(Any) Seq
	Map(Mapper) Seq
	ForEach(func(Any))
	Find(Predicate) Any
	Filter(Predicate) Seq
	Reduce(Any) func(Reducer) Any
	Flatten() Seq
}

type Node interface {
	Value() Any
	Parent() Node
	LeftChild() Node
	RightChild() Node
}

type Tree interface {
	Iterable
	Root() Node
}

//type Map map[Any]Any
//type Set []Any
