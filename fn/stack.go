package fn

type Stack []Any

var EmptyStack = Stack{}

func (s *Stack) Push(item Any) *Stack {
	*s = append(*s, item)
	return s
}

func (s *Stack) Pop() Any {
	if len(*s) == 0 {
		return EmptyStack
	}

	current := *s
	lastIndex := len(current) - 1
	lastItem := current[lastIndex]
	*s = current[:lastIndex]
	return lastItem
}
