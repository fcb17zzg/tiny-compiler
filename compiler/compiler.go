package compiler

import (
	"go-monkey-compiler/ast"
	"go-monkey-compiler/code"
	"go-monkey-compiler/object"
)

// 编译器
type Compiler struct {
	instructions code.Instructions // 字节码
	constants    []object.Object   // 常量池
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

// 字节码
type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}
