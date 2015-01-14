package collections

type GenericIterable struct {
	items  []Any
	cursor int
}

func NewGenericIterable(items ...Any) Iterable {
	return &GenericIterable{items, -1}
}

func (this *GenericIterable) HasNext() bool {
	return this.cursor < len(this.items)-1
}

func (this *GenericIterable) Next() Any {
	if this.HasNext() {
		this.cursor++
		return this.items[this.cursor]
	}
	return nil
}

func (this *GenericIterable) Length() int {
	return len(this.items)
}

func (this *GenericIterable) Reset() {
	this.cursor = -1
}

func (this *GenericIterable) AsArray() []Any {
	var array []Any
	for this.HasNext() {
		array = append(array, this.Next())
	}
	return array
}
