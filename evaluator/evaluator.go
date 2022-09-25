package evaluator

import (
	"vabna/ast"
	"vabna/object"

)

var (
    NULL = &object.Null{}
    TRUE = &object.Boolean{Value: true}
    FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Obj{
    switch node := node.(type){
        case *ast.Program:
            return evalStmts(node.Stmts)
        case *ast.ExprStmt:
            return Eval(node.Expr)
        case *ast.IntegerLit:
            return &object.Integer{Value: node.Value}
        case *ast.Boolean:
            return getBoolObj(node.Value)
        case *ast.PrefixExpr:
            r := Eval(node.Right)
            return evalPrefixExpr(node.Op , r)
        case *ast.InfixExpr:
            l := Eval(node.Left)
            r := Eval(node.Right)
            return evalInfixExpr(node.Op , l , r)
        case *ast.IfExpr:
            return evalIfExpr(node)
        case *ast.ReturnStmt:
            val := Eval(node.ReturnVal)
            return &object.ReturnValue{Value: val}
        case *ast.BlockStmt:
            return evalBlockStmt(node)        
    }

    return nil
}

func evalBlockStmt(block *ast.BlockStmt) object.Obj{

    var res object.Obj

    for _, stmt := range block.Stmts{
        res := Eval(stmt)

        if res != nil && res.Type() == object.RETURN_VAL_OBJ{
            return res
        }
    }

    return res
}

func evalProg(prog *ast.Program) object.Obj{
    var res object.Obj

    for _, stmt := range prog.Stmts{
        res = Eval(stmt)

        if rvalue,ok := res.(*object.ReturnValue); ok{
            return rvalue.Value
        }
    }

    return res
}

func evalIfExpr(iex *ast.IfExpr) object.Obj{
    cond := Eval(iex.Cond)
    
    if isTruthy(cond){
        return Eval(iex.TrueBlock)
    }else if iex.ElseBlock != nil{
        return Eval(iex.ElseBlock)
    }else {
        return NULL
    }

}

func isTruthy(obj object.Obj) bool{
    switch obj{
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

func evalInfixExpr(op string , l, r object.Obj) object.Obj{

    switch{
        case l.Type() == object.INT_OBJ && r.Type() == object.INT_OBJ:
            return evalIntInfixExpr(op , l , r)
        case op == "==":
            return getBoolObj(l == r)
        case op == "!=":
            return getBoolObj(l != r)
        default:
            return NULL
    }
}

func evalIntInfixExpr(op string  , l , r object.Obj) object.Obj{
    lval := l.(*object.Integer).Value
    rval := r.(*object.Integer).Value

    switch op{
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
            return NULL
    }
}

func evalPrefixExpr(op string , right object.Obj) object.Obj{
    switch op{
        case "!":
            return evalBangOp(right)
        case "-":
            return evalMinusPrefOp(right)
        default:
            return NULL

    }
}

func evalMinusPrefOp(right object.Obj) object.Obj{
    if right.Type() != object.INT_OBJ{
        return NULL
    }
    val := right.(*object.Integer).Value
    return &object.Integer{Value: -val}
}

func evalBangOp(r object.Obj) object.Obj{
    switch r{
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

func getBoolObj(inp bool) *object.Boolean{
    if inp{
        return TRUE
    }else{
        return FALSE
    }
}

func evalStmts(stmts []ast.Stmt) object.Obj{
    var res object.Obj

    for _,stmt := range stmts{
        res = Eval(stmt)
        
        if rvalue , ok := res.(*object.ReturnValue); ok{
            return rvalue.Value
        }
    }

    return res
}
