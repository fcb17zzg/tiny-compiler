package vm

import (
	"go-monkey-compiler/code"
	"go-monkey-compiler/object"
)

type Frame struct {
	fn          *object.Closure
	ip          int // instruction pointer
	basePointer int
}

func NewFrame(fn *object.Closure, basePointer int) *Frame {
	return &Frame{
		fn:          fn,
		ip:          -1,
		basePointer: basePointer,
	}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Fn.Instructions
}
