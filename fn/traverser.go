package fn

type Traverser struct {
	Iterable
}

func NewTraverser(iterator Iterable) Traverser {
	return Traverser{iterator}
}

func (this Traverser) ForEach(f func(Any)) {
	for this.HasNext() {
		f(this.Next())
	}
}

func (this Traverser) Map(m Mapper) Traverser {
	var list []Any
	this.ForEach(func(item Any) {
		list = append(list, m(item))
	})

	return NewTraverser(NewIterator(list))
}

func (this Traverser) Find(p Predicate) Any {
	for this.HasNext() {
		item := this.Next()
		if p(item) {
			return item
		}
	}

	return None(nil)
}

func (this Traverser) Filter(p Predicate) Traverser {
	var matched []Any

	this.ForEach(func(item Any) {
		if p(item) {
			matched = append(matched, item)
		}
	})

	return NewTraverser(NewIterator(matched))
}

func (this Traverser) Reduce(accumulator Any) func(Reducer) Any {
	return func(r Reducer) Any {
		this.ForEach(func(item Any) {
			accumulator = r(accumulator, item)
		})

		return accumulator
	}
}

func (this Traverser) Flatten() Traverser {
	flattened := this.Reduce([]Any{})(func(acc Any, next Any) Any {
		switch next.(type) {
		case []Any:
			acc = append(acc.([]Any), next.([]Any)...)
		default:
			acc = append(acc.([]Any), next)
		}
		return acc
	}).([]Any)

	return Traverser{NewIterator(flattened)}
}
