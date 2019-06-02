package parser

import (
	"sync"
)

type Ast struct {
	rootsLock sync.Mutex
	roots     map[string]IProgramContext
}

func NewEmptyAst() *Ast {
	a := new(Ast)
	a.roots = make(map[string]IProgramContext)
	return a
}

func (a *Ast) AddRoot(source string, root IProgramContext) {
	a.rootsLock.Lock()
	defer a.rootsLock.Unlock()

	a.roots[source] = root
}

func (a *Ast) VisitAsync(visitor ECMAScriptVisitor) {
	a.rootsLock.Lock()
	defer a.rootsLock.Unlock()

	for _, root := range a.roots {
		go visitor.Visit(root)
	}
}

func (a *Ast) Visit(visitor ECMAScriptVisitor) {
	a.rootsLock.Lock()
	defer a.rootsLock.Unlock()

	for _, root := range a.roots {
		visitor.Visit(root)
	}
}
