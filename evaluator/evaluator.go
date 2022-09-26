package evaluator

import (
	"fmt"
	"vabna/ast"
	"vabna/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Env) object.Obj {
	switch node := node.(type) {
	case *ast.Program:
		return evalProg(node, env)
	case *ast.ExprStmt:
		//fmt.Println("Eval Expr => ", node)
		return Eval(node.Expr, env)
	case *ast.IntegerLit:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return getBoolObj(node.Value)
	case *ast.PrefixExpr:
		r := Eval(node.Right, env)
		if isErr(r) {
			return r
		}
		return evalPrefixExpr(node.Op, r)
	case *ast.InfixExpr:
		l := Eval(node.Left, env)
		if isErr(l) {
			return l
		}
		r := Eval(node.Right, env)
		if isErr(r) {
			return r
		}
		return evalInfixExpr(node.Op, l, r)
	case *ast.IfExpr:
		return evalIfExpr(node, env)
	case *ast.ReturnStmt:
		val := Eval(node.ReturnVal, env)
		if isErr(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.BlockStmt:
		return evalBlockStmt(node, env)
	case *ast.LetStmt:
		val := Eval(node.Value, env)
		if isErr(val) {
			return val
		}

		env.Set(node.Name.Value, val)
	case *ast.Identifier:
		return evalId(node, env)
	case *ast.FunctionLit:
		pms := node.Params
		body := node.Body
		return &object.Function{Params: pms, Body: body, Env: env}
	case *ast.CallExpr:
		fnc := Eval(node.Func, env)
		if isErr(fnc) {
			return fnc
		}

		args := evalExprs(node.Args, env)
		if len(args) == 1 && isErr(args[0]) {
			return args[0]
		}

		return applyFunc(fnc, args)

	case *ast.StringLit:
		return &object.String{Value: node.Value}
	case *ast.ArrLit:
		elms := evalExprs(node.Elms, env)
		if len(elms) == 1 && isErr(elms[0]) {
			return elms[0]
		}

		return &object.Array{Elms: elms}

	case *ast.IndexExpr:
		left := Eval(node.Left, env)
		if isErr(left) {
			return nil
		}

		index := Eval(node.Index, env)
		if isErr(index) {
			return index
		}

		return evalIndexExpr(left, index)
	}

	return nil
}



func evalIndexExpr(left, index object.Obj) object.Obj {

	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INT_OBJ:
		return evalArrIndexExpr(left, index)

	default:
		return newErr("Unsupported Index Operator %s ", left.Type())
	}

}

func evalArrIndexExpr(arr, index object.Obj) object.Obj {
	arrObj := arr.(*object.Array)
	idx := index.(*object.Integer).Value

	max := int64(len(arrObj.Elms) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arrObj.Elms[idx]
}

func applyFunc(fn object.Obj, args []object.Obj) object.Obj {

	switch fn := fn.(type) {
	case *object.Function:
		eEnv := extendFuncEnv(fn, args)
		evd := Eval(fn.Body, eEnv)
		return unwrapRValue(evd)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newErr("%s is not a function", fn.Type())

	}
}

func extendFuncEnv(fn *object.Function, args []object.Obj) *object.Env {
	env := object.NewEnclosedEnv(fn.Env)

	for pId, param := range fn.Params {
		env.Set(param.Value, args[pId])
	}

	return env
}

func unwrapRValue(o object.Obj) object.Obj {
	if rv, ok := o.(*object.ReturnValue); ok {
		return rv.Value
	}

	return o

}

func evalExprs(es []ast.Expr, env *object.Env) []object.Obj {
	var res []object.Obj

	for _, e := range es {
		ev := Eval(e, env)

		if isErr(ev) {
			return []object.Obj{ev}
		}

		res = append(res, ev)
	}

	return res
}

func evalId(node *ast.Identifier, env *object.Env) object.Obj {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newErr("id not found : " + node.Value)
	//	return val
}

func newErr(format string, a ...interface{}) *object.Error {
	return &object.Error{Msg: fmt.Sprintf(format, a...)}
}

func isErr(obj object.Obj) bool {
	if obj != nil {
		return obj.Type() == object.ERR_OBJ
	}

	return false
}

func evalBlockStmt(block *ast.BlockStmt, env *object.Env) object.Obj {

	var res object.Obj

	for _, stmt := range block.Stmts {
		res = Eval(stmt, env)

		//fmt.Println("E_BS=> " , res)

		if res != nil {
			rtype := res.Type()
			if rtype == object.RETURN_VAL_OBJ || rtype == object.ERR_OBJ {
				//fmt.Println("RET => " ,  res)
				return res
			}
		}
	}
	//fmt.Println("EBS 2=>" ,res)
	return res
}

func evalProg(prog *ast.Program, env *object.Env) object.Obj {
	var res object.Obj

	for _, stmt := range prog.Stmts {
		res = Eval(stmt, env)

		switch res := res.(type) {
		case *object.ReturnValue:
			return res.Value
		case *object.Error:
			return res
		}
	}

	return res
}

func evalIfExpr(iex *ast.IfExpr, env *object.Env) object.Obj {
	cond := Eval(iex.Cond, env)

	if isErr(cond) {
		return cond
	}

	if isTruthy(cond) {
		return Eval(iex.TrueBlock, env)
	} else if iex.ElseBlock != nil {
		return Eval(iex.ElseBlock, env)
	} else {
		return NULL
	}

}

func isTruthy(obj object.Obj) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func evalInfixExpr(op string, l, r object.Obj) object.Obj {

	switch {
	case l.Type() == object.INT_OBJ && r.Type() == object.INT_OBJ:
		return evalIntInfixExpr(op, l, r)
	case l.Type() == object.STRING_OBJ && r.Type() == object.STRING_OBJ:
		return evalStringInfixExpr(op, l, r)
	case op == "==":
		return getBoolObj(l == r)
	case op == "!=":
		return getBoolObj(l != r)
	case l.Type() != r.Type():
		return newErr("Type mismatch:  %s %s %s ", l.Type(), op, r.Type())
	default:
		return newErr("unknown Operator : %s %s %s", l.Type(), op, r.Type())
	}
}

func evalStringInfixExpr(op string, l, r object.Obj) object.Obj {
	if op != "+" {
		return newErr("Unknown Operator %s %s %s", l.Type(), op, r.Type())
	}

	lval := l.(*object.String).Value
	rval := r.(*object.String).Value
	return &object.String{Value: lval + rval}
}

func evalIntInfixExpr(op string, l, r object.Obj) object.Obj {
	lval := l.(*object.Integer).Value
	rval := r.(*object.Integer).Value

	switch op {
	case "+":
		return &object.Integer{Value: lval + rval}
	case "-":
		return &object.Integer{Value: lval + rval}
	case "*":
		return &object.Integer{Value: lval * rval}
	case "/":
		return &object.Integer{Value: lval / rval}
	case "<":
		return getBoolObj(lval < rval)
	case ">":
		return getBoolObj(lval > rval)
	case "==":
		return getBoolObj(lval == rval)
	case "!=":
		return getBoolObj(lval != rval)

	default:
		return newErr("unknown Operator : %s %s %s", l.Type(), op, r.Type())
	}
}

func evalPrefixExpr(op string, right object.Obj) object.Obj {
	switch op {
	case "!":
		return evalBangOp(right)
	case "-":
		return evalMinusPrefOp(right)
	default:
		return newErr("Unknown Operator : %s%s", op, right.Type())

	}
}

func evalMinusPrefOp(right object.Obj) object.Obj {
	if right.Type() != object.INT_OBJ {
		return newErr("unknown Operator : -%s", right.Type())
	}
	val := right.(*object.Integer).Value
	return &object.Integer{Value: -val}
}

func evalBangOp(r object.Obj) object.Obj {
	switch r {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func getBoolObj(inp bool) *object.Boolean {
	if inp {
		return TRUE
	} else {
		return FALSE
	}
}

func evalStmts(stmts []ast.Stmt, env *object.Env) object.Obj {
	var res object.Obj

	for _, stmt := range stmts {
		res = Eval(stmt, env)

		if rvalue, ok := res.(*object.ReturnValue); ok {
			return rvalue.Value
		}
	}

	return res
}
