package stack

type Stack []interface{}

func (s Stack) Push(v interface{}) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, interface{}, bool) {
	if s.IsEmpty() {
		return s, nil, false
	}

	l := len(s)
	return s[:l-1], s[l-1], true
}

func (s Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	return s[len(s)-1], true
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
