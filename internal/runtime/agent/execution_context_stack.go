package agent

import "github.com/gojisvm/gojis/internal/runtime/agent/stack"

type ExecutionContextStack struct {
	stack stack.Stack
}

func NewExecutionContextStack() *ExecutionContextStack {
	return new(ExecutionContextStack)
}

func (s ExecutionContextStack) IsEmpty() bool {
	return s.stack.Peek() == nil
}

func (s ExecutionContextStack) Push(ctx *ExecutionContext) {
	s.stack.Push(ctx)
}

func (s ExecutionContextStack) Pop() *ExecutionContext {
	elem := s.stack.Pop()
	if elem == nil {
		return nil
	}

	return elem.(*ExecutionContext)
}

func (s ExecutionContextStack) Peek() *ExecutionContext {
	elem := s.stack.Peek()
	if elem == nil {
		return nil
	}

	return elem.(*ExecutionContext)
}
