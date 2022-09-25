package object

import "fmt"

const (
    INT_OBJ = "INTEGER"
    BOOL_OBJ = "BOOLEAN"
    RETURN_VAL_OBJ = "RETURN_VAL"
    NULL_OBJ = "NIL"
)

type ObjType string

type Obj interface{
    Type() ObjType
    Inspect() string
}


//Integer 1,2,3,4,5.....100
type Integer struct{
    Value int64
}

func (i* Integer) Type() ObjType{
    return INT_OBJ
}

func (i *Integer) Inspect() string {
    return fmt.Sprintf("%d" , i.Value)
}


//Booleans true,false
type Boolean struct{
    Value bool
}

func (b *Boolean) Type() ObjType { return BOOL_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t" , b.Value) }


//NULL_OBJ
type Null struct {}

func (n *Null) Type() ObjType { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }

type ReturnValue struct{
    Value Obj
}

func (r *ReturnValue) Type() ObjType { return RETURN_VAL_OBJ }
func (r *ReturnValue) Inspect() string { return r.Value.Inspect() }
