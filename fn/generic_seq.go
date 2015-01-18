package fn

type GenericSeq struct {
	Iterable
}

func NewGenericSeq(args ...Any) Seq {
	return &GenericSeq{NewGenericIterable(args...)}
}

func (this *GenericSeq) ForEach(f func(Any)) {
	for this.HasNext() {
		f(this.Next())
	}

	this.Reset()
}

func (this *GenericSeq) Map(m Mapper) Seq {
	var seq []Any
	for this.HasNext() {
		seq = append(seq, m(this.Next()))
	}

	this.Reset()
	return NewGenericSeq(seq...)
}

func (this *GenericSeq) FindAll(p Predicate) Seq {
	var matched []Any

	for this.HasNext() {
		item := this.Next()
		if p(item) {
			matched = append(matched, item)
		}
	}

	this.Reset()

	return NewGenericSeq(matched...)
}
