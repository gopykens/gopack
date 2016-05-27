package qlang

import (
	"qlang.io/exec.v2"
	"qlang.io/qlang.spec.v1"
)

// -----------------------------------------------------------------------------

func (p *Compiler) vMap() {

	arity := p.popArity()
	p.code.Block(exec.Call(qlang.MapFrom, arity*2))
}

// -----------------------------------------------------------------------------

func (p *Compiler) tSlice() {

	p.code.Block(exec.Slice)
}

// -----------------------------------------------------------------------------

func (p *Compiler) vSlice() {

	hasSlice := p.popArity()
	hasInit := 0
	arityInit := 0
	if hasSlice > 0 {
		hasInit = p.popArity()
		if hasInit > 0 {
			arityInit = p.popArity()
		}
	}
	arity := p.popArity()

	if hasSlice > 0 {
		if arity > 0 {
			panic("must be []type")
		}
		if hasInit > 0 { // []T{a1, a2, ...}
			p.code.Block(exec.SliceFromTy(arityInit))
		} else { // []T
			p.code.Block(exec.Slice)
		}
	} else { // [a1, a2, ...]
		p.code.Block(exec.SliceFrom(arity))
	}
}

// -----------------------------------------------------------------------------
