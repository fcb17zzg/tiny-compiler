package vm

import (
	"go-monkey-compiler/code"
	"go-monkey-compiler/object"
)

type Frame struct {
	fn *object.CompiledFunction
	ip int // instruction pointer
}

func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{
		fn: fn,
		ip: -1,
	}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
