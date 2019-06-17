package stack

type sliceStack struct {
	s internalSliceStack
}

func NewSliceStack() Stack {
	s := new(sliceStack)
	s.s = internalSliceStack{}
	return s
}

func (s *sliceStack) Push(v interface{}) {
	s.s = s.s.push(v)
}

func (s *sliceStack) Pop() (elem interface{}) {
	s.s, elem, _ = s.s.pop()
	return
}

func (s *sliceStack) Peek() (elem interface{}) {
	elem, _ = s.s.peek()
	return
}

type internalSliceStack []interface{}

func (s internalSliceStack) push(v interface{}) internalSliceStack {
	return append(s, v)
}

func (s internalSliceStack) pop() (internalSliceStack, interface{}, bool) {
	if s.isEmpty() {
		return s, nil, false
	}

	l := len(s)
	return s[:l-1], s[l-1], true
}

func (s internalSliceStack) peek() (interface{}, bool) {
	if s.isEmpty() {
		return nil, false
	}

	return s[len(s)-1], true
}

func (s internalSliceStack) isEmpty() bool {
	return len(s) == 0
}
