package stack

import (
	"unsafe"
)

type unsafeStack struct {
	s internalSliceStack
}

func NewUnsafeStack() Stack {
	s := new(unsafeStack)
	s.s = internalSliceStack{}
	return s
}

func (s *unsafeStack) Push(v interface{}) {
	ptr := unsafe.Pointer(&v) // #nosec
	s.s = s.s.push(ptr)
}

func (s *unsafeStack) Pop() (elem interface{}) {
	s.s, elem, _ = s.s.pop()
	if elem == nil {
		return
	}

	elem = *(*interface{})(elem.(unsafe.Pointer))
	return
}

func (s *unsafeStack) Peek() (elem interface{}) {
	elem, _ = s.s.peek()
	if elem == nil {
		return
	}

	elem = *(*interface{})(elem.(unsafe.Pointer))
	return
}
