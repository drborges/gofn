package fn

type Cursor int

type Iterator struct {
	items  []Any
	cursor Cursor
}

func NewIterator(items ...Any) Iterable {
	return &Iterator{items, -1}
}

func (this *Iterator) HasNext() bool {
	maxCursorPosition := Cursor(len(this.items) - 1)
	return this.cursor < maxCursorPosition
}

func (this *Iterator) Next() Any {
	if this.HasNext() {
		this.cursor++
		return this.items[this.cursor]
	}
	return None(nil)
}

func (this *Iterator) Length() int {
	return len(this.items)
}

func (this *Iterator) Reset() {
	this.cursor = -1
}

func (this *Iterator) AsArray() []Any {
	return this.items
}
