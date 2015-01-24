package fn

type Converter func(Any) Any

func Iter2Array(i Iterable) []Any {
	var arr []Any

	for i.HasNext() {
		arr = append(arr, i.Next())
	}

	return arr
}
