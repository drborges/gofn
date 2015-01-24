package fn

type Cursor int

type Iterator struct {
	items []Any
}

var EmptyIterator = None(nil)

func NewIterator(items []Any) Iterable {
	return &Iterator{items}
}

func NewIteratorFromItems(items ...Any) Iterable {
	return &Iterator{items}
}

func (this *Iterator) HasNext() bool {
	return len(this.items) > 0
}

func (this *Iterator) Next() Any {
	if this.HasNext() {
		next := this.items[0]
		this.items = this.items[1:]
		return next
	}
	return EmptyIterator
}
