package agent

import "gitlab.com/gojis/vm/internal/runtime/agent/stack"

type ExecutionContextStack struct {
	stack stack.Stack
}

func NewExecutionContextStack() *ExecutionContextStack {
	return new(ExecutionContextStack)
}

func (s ExecutionContextStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s ExecutionContextStack) Push(ctx *ExecutionContext) {
	s.stack = s.stack.Push(ctx)
}

func (s ExecutionContextStack) Pop() (*ExecutionContext, bool) {
	stack, elem, ok := s.stack.Pop()
	s.stack = stack
	return elem.(*ExecutionContext), ok
}

func (s ExecutionContextStack) Peek() (*ExecutionContext, bool) {
	val, ok := s.stack.Peek()
	return val.(*ExecutionContext), ok
}

func (s ExecutionContextStack) FindTopDown(predicate func(*ExecutionContext) bool) *ExecutionContext {
	for i := len(s.stack) - 1; i >= 0; i-- {
		if predicate(s.stack[i].(*ExecutionContext)) {
			return s.stack[i].(*ExecutionContext)
		}
	}
	return nil
}
