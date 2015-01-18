package fn

type GenericSeq struct {
	Iterable
}

func NewGenericSeq(args ...Any) Seq {
	return &GenericSeq{NewIterator(args...)}
}

func (this *GenericSeq) ForEach(f func(Any)) {
	for this.HasNext() {
		f(this.Next())
	}

	this.Reset()
}

func (this *GenericSeq) Map(m Mapper) Seq {
	var seq []Any
	this.ForEach(func(item Any) {
		seq = append(seq, m(item))
	})

	return NewGenericSeq(seq...)
}

func (this *GenericSeq) Find(p Predicate) Any {
	for this.HasNext() {
		item := this.Next()
		if p(item) {
			return item
		}
	}

	this.Reset()
	return nil
}

func (this *GenericSeq) Filter(p Predicate) Seq {
	var matched []Any

	this.ForEach(func(item Any) {
		if p(item) {
			matched = append(matched, item)
		}
	})

	return NewGenericSeq(matched...)
}

func (this *GenericSeq) Reduce(accumulator Any) func(Reducer) Any {
	return func(r Reducer) Any {
		this.ForEach(func(item Any) {
			accumulator = r(accumulator, item)
		})

		return accumulator
	}
}
