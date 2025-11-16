package evaluator

import (
	"fmt"
	"go-monkey-compiler/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return NewError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return NewError("argument to `first` must be ARRAY, got %s",
					args[0].Type())
			}
			array := args[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[0]
			}
			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return NewError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}
			array := args[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[len(array.Elements)-1]
			}
			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return NewError("argument to `rest` must be ARRAY, got %s",
					args[0].Type())
			}
			array := args[0].(*object.Array)
			if len(array.Elements) > 0 {
				newElements := make([]object.Object, len(array.Elements)-1)
				copy(newElements, array.Elements[1:])
				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return NewError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return NewError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
			array := args[0].(*object.Array)
			newElements := make([]object.Object, len(array.Elements)+1)
			copy(newElements, array.Elements)
			newElements[len(array.Elements)] = args[1]
			return &object.Array{Elements: newElements}
		},
	},
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}
