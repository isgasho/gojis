package agent

import (
	"fmt"

	"github.com/gojisvm/gojis/internal/runtime/agent/job"
	"github.com/gojisvm/gojis/internal/runtime/binding"
	"github.com/gojisvm/gojis/internal/runtime/errors"
	"github.com/gojisvm/gojis/internal/runtime/lang"
	"github.com/gojisvm/gojis/internal/runtime/realm"
	"github.com/google/uuid"
)

type QueueKind uint8

const (
	QueueScript QueueKind = iota
	QueuePromise
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

	ScriptJobs  *job.Queue
	PromiseJobs *job.Queue
}

func New() *Agent {
	a := new(Agent)
	a.ExecutionContextStack = NewExecutionContextStack()
	a.LittleEndian = false
	panic("TODO: default value: CanBlock")
	a.Signifier = NewID()
	a.ScriptJobs = job.NewQueue()
	a.PromiseJobs = job.NewQueue()
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

func (a *Agent) RunningExecutionContext() *ExecutionContext {
	return a.ExecutionContextStack.Peek()
}

func (a *Agent) GetActiveScriptOrModule() lang.InternalValue {
	if a.ExecutionContextStack.IsEmpty() {
		return lang.Null
	}

	panic("TODO")
}

func (a *Agent) ResolveBinding(name lang.String, env binding.Environment) *binding.Reference {
	if env == nil {
		env = a.RunningExecutionContext().LexicalEnvironment
	}

	strict := false // FIXME: 8.3.2, Step 3
	panic("TODO: 8.3.2, Step 3")

	return binding.GetIdentifierReference(env, name, strict)
}

func (a *Agent) GetThisEnvironment() binding.Environment {
	lex := a.RunningExecutionContext().LexicalEnvironment

	/*
		Step out until an environment has a this binding.
		The Global Environment has a this binding, and is the
		only environment which has no outer environment, so this
		will always terminate.
	*/
	for !lex.HasThisBinding() {
		lex = lex.Outer()

		if lex == nil {
			panic("Outer environment cannot be nil, this means that the global object does not have a this binding, which must not happen, or that an environment which is not the global environment has a nil reference as an outer environment.")
		}
	}

	return lex
}

func (a *Agent) ResolveThisBinding() (lang.Value, errors.Error) {
	return a.GetThisEnvironment().GetThisBinding()
}

func (a *Agent) GetNewTarget() lang.Value {
	panic("TODO")
}

func (a *Agent) GetGlobalObject() lang.Value {
	return a.RunningExecutionContext().Realm.GlobalObj
}

func (a *Agent) EnqueueJob(q QueueKind, jobName string, arguments []lang.Value) {
	callerCtx := a.RunningExecutionContext()
	callerRealm := callerCtx.Realm
	callerScriptOrModule := callerCtx.ScriptOrModule
	pending := job.PendingJob{
		Job:            jobName,
		Arguments:      arguments,
		Realm:          callerRealm,
		ScriptOrModule: callerScriptOrModule,
		HostDefined:    lang.Undefined,
	}
	// TODO: do we need to modify pending in any way?

	switch q {
	case QueueScript:
		a.ScriptJobs.Enqueue(pending)
	case QueuePromise:
		a.PromiseJobs.Enqueue(pending)
	default:
		panic(fmt.Sprintf("Unknown queue kind: %v", q))
	}
}

func (a *Agent) InitializeHostDefinedRealm() {
	r := realm.CreateRealm()
	newCtx := &ExecutionContext{
		Function:       lang.Null,
		Realm:          r,
		ScriptOrModule: lang.Null,
	}
	a.ExecutionContextStack.Push(newCtx)

	/*
		If the host requires use of an exotic object to serve as realm's global object, let global be such an object created in
		an implementation-defined manner. Otherwise, let global be undefined, indicating that an ordinary object
		should be created as the global object.
	*/
	// Let global be undefined
	global := lang.Undefined

	/*
		If the host requires that the thisthis binding in realm's global scope return an object other than the global object,
		let thisValue be such an object created in an implementation-defined manner. Otherwise, let thisValue be
		undefined, indicating that realm's global thisthis binding should be the global object.
	*/
	// Let thisValue be undefined
	thisValue := lang.Undefined

	r.SetRealmGlobalObject(global, thisValue)

	globalObj := r.SetDefaultGlobalBindings()
	// TODO: Create any implementation-defined global object properties on globalObj.
	_ = globalObj
}

func (a *Agent) RunJobs() {
	a.InitializeHostDefinedRealm()

	panic("TODO")
}
