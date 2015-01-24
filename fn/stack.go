package fn

type Stack []Any

func (s Stack) Push(item Any) Stack {
	return Stack(append(s, item))
}

func (s Stack) Pop(dst *Any) Stack {
	lastIndex := len(s) - 1
	*dst = s[lastIndex]

	return s[:lastIndex]
}
