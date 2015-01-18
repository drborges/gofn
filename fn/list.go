package fn

type List struct {
	Iterable
}

func NewList(args ...Any) Seq {
	return &List{NewIterator(args...)}
}

func (this *List) ForEach(f func(Any)) {
	for this.HasNext() {
		f(this.Next())
	}

	this.Reset()
}

func (this *List) Map(m Mapper) Seq {
	var seq []Any
	this.ForEach(func(item Any) {
		seq = append(seq, m(item))
	})

	return NewList(seq...)
}

func (this *List) Find(p Predicate) Any {
	for this.HasNext() {
		item := this.Next()
		if p(item) {
			return item
		}
	}

	this.Reset()
	return nil
}

func (this *List) Filter(p Predicate) Seq {
	var matched []Any

	this.ForEach(func(item Any) {
		if p(item) {
			matched = append(matched, item)
		}
	})

	return NewList(matched...)
}

func (this *List) Reduce(accumulator Any) func(Reducer) Any {
	return func(r Reducer) Any {
		this.ForEach(func(item Any) {
			accumulator = r(accumulator, item)
		})

		return accumulator
	}
}

func (this *List) Flatten() Seq {
	flattened := this.Reduce([]Any{})(func(acc Any, next Any) Any {
		switch next.(type) {
		case []Any:
			acc = append(acc.([]Any), next.([]Any)...)
		default:
			acc = append(acc.([]Any), next)
		}
		return acc
	})

	return NewList(flattened.([]Any)...)
}
