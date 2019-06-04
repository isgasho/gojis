package agent

import (
	"github.com/google/uuid"
	"gitlab.com/gojis/vm/internal/runtime/lang"
)

type AgentID uuid.UUID

func NewID() AgentID {
	return AgentID(uuid.New())
}

type Agent struct {
	ExecutionContextStack *ExecutionContextStack

	LittleEndian       bool
	CanBlock           bool
	Signifier          AgentID
	IsLockFree1        bool
	IsLockFree2        bool
	CandidateExecution interface{} // TODO: Table 26, CandidateExecutionRecord
}

func New() *Agent {
	a := new(Agent)
	a.ExecutionContextStack = NewExecutionContextStack()
	a.LittleEndian = false
	panic("TODO: default value: CanBlock")
	a.Signifier = NewID()
	return a
}

// TODO: remove once spec is implemented (use agent.Signifier instead)
func (a *Agent) AgentSignifier() AgentID {
	return a.Signifier
}

// TODO: remove once spec is implemented (use agent.CanBlock instead)
func (a *Agent) AgentCanSuspend() bool {
	return a.CanBlock
}

func (a *Agent) GetActiveScriptOrModule() lang.InternalValue {
	if a.ExecutionContextStack.IsEmpty() {
		return lang.Null
	}

	ctx := a.ExecutionContextStack.FindTopDown(func(ctx *ExecutionContext) bool {
		if ctx.ScriptOrModule != lang.Null {
			return true
		}
		return false
	})
	if ctx == nil {
		return lang.Null
	}

	return ctx.ScriptOrModule
}
