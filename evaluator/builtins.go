package evaluator

import (
	"fmt"
	"vabna/object"
)

var builtins = map[string]*object.Builtin{

	"len": {
		Fn: func(args ...object.Obj) object.Obj {
			if len(args) != 1 {
				return newErr("wrong number of arguments. got %d but wanted 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elms))}
			default:
				return newErr("argument type %s to `len` is not supported", args[0].Type())
			}
		},
	},

	"প্রথম": {
		Fn: func(args ...object.Obj) object.Obj {
			if len(args) != 1 {
				return newErr("wrong number of argument %d", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newErr("first cannot be used with %s", args[0].Type())
			}

			array := args[0].(*object.Array)
			if len(array.Elms) > 0 {
				return array.Elms[0]
			}
			return NULL
		},
	},

	"শেষ": {
		Fn: func(args ...object.Obj) object.Obj {
			if len(args) != 1 {
				return newErr("wrong number of argument %d", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newErr("last cannot be used with %s", args[0].Type())
			}

			array := args[0].(*object.Array)
			arr_len := len(array.Elms)
			if arr_len > 0 {
				return array.Elms[arr_len-1]
			}
			return NULL
		},
	},

	"বাদবাকি": {
		Fn: func(args ...object.Obj) object.Obj {
			if len(args) != 1 {
				return newErr("wrong number of argument %d", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newErr("rest cannot be used with %s", args[0].Type())
			}

			array := args[0].(*object.Array)
			arrLen := len(array.Elms)
			if arrLen > 0 {
				newElms := make([]object.Obj, arrLen-1, arrLen-1)
				copy(newElms, array.Elms[1:arrLen])
				return &object.Array{Elms: newElms}
			}
			return NULL
		},
	},

	"যোগ": {
		Fn: func(args ...object.Obj) object.Obj {

			if len(args) != 2 {
				return newErr("wrong number of argument %d", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newErr("push cannot be used with %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			arrLen := len(arr.Elms)

			newElms := make([]object.Obj, arrLen+1, arrLen+1)
			copy(newElms, arr.Elms)
			newElms[arrLen] = args[1]
			return &object.Array{Elms: newElms}

		},
	},

	"দেখাও": {
		Fn: func(args ...object.Obj) object.Obj {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}
