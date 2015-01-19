package fn

type List []Any

func NewList(items ...Any) List {
	return List(items)
}

func (this List) ForEach(f func(Any)) {
	for _, item := range this {
		f(item)
	}
}

func (this List) Map(m Mapper) List {
	var list []Any
	this.ForEach(func(item Any) {
		list = append(list, m(item))
	})

	return list
}

func (this List) Find(p Predicate) Any {
	for _, item := range this {
		if p(item) {
			return item
		}
	}

	return None(nil)
}

func (this List) Filter(p Predicate) List {
	var matched []Any

	this.ForEach(func(item Any) {
		if p(item) {
			matched = append(matched, item)
		}
	})

	return matched
}

func (this List) Reduce(accumulator Any) func(Reducer) Any {
	return func(r Reducer) Any {
		this.ForEach(func(item Any) {
			accumulator = r(accumulator, item)
		})

		return accumulator
	}
}

func (this List) Flatten() List {
	return this.Reduce(List{})(func(acc Any, next Any) Any {
		switch next.(type) {
		case []Any:
			acc = append(acc.(List), next.([]Any)...)
		case List:
			acc = append(acc.(List), next.(List)...)
		default:
			acc = append(acc.(List), next)
		}
		return acc
	}).(List)
}

func (this List) Append(item Any) List {
	this = append(this, item)
	return this
}
