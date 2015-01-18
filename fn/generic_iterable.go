package fn

type Cursor int

type GenericIterable struct {
	items  []Any
	cursor Cursor
}

func NewGenericIterable(items ...Any) Iterable {
	return &GenericIterable{items, -1}
}

func (this *GenericIterable) HasNext() bool {
	maxCursorPosition := Cursor(len(this.items) - 1)
	return this.cursor < maxCursorPosition
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
	return this.items
}
