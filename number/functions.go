package number

import (
	"math/big"
	"vabna/token"
)


func NumberCompare(op string, a Number , b Number)(){

}

func NumberOperation(op string, n Number , x Number) Number{
    
    var fb  *big.Float
    var fa  *big.Float
    var val *big.Float
    switch n.GetType(){
        case "FLOAT":
            fa = &n.Value.(*FloatNumber).Value
            //var fb  *big.Float

            //var val *big.Float
            if x.IsInt{
                b := x.Value.(*IntNumber).Value

                fb  = new(big.Float).SetInt(&b)
            }else{
                fb  =  &x.Value.(*FloatNumber).Value
            }

             switch op{
                    case token.PLUS:
                        val = new(big.Float).Add( fa , fb) 
                    case token.MINUS:
                        val = new(big.Float).Sub( fa , fb)
                    case token.MUL:
                        val = new(big.Float).Mul( fa, fb)
                    case token.DIV:
                        val = new(big.Float).Quo( fa , fb)
                    
                }
            return Number{ Value: &FloatNumber{ Value: *val } , IsInt: false }
        case "INT":
            a := n.Value.(*IntNumber).Value

            if x.IsInt{
                ib := x.Value.(*IntNumber).Value
                var val *big.Int
                switch op{
                    case token.PLUS:
                        val = new(big.Int).Add(&a , &ib) 
                    case token.MINUS:
                        val = new(big.Int).Sub(&a , &ib)
                    case token.MUL:
                        val = new(big.Int).Mul(&a, &ib)
                    case token.DIV:
                        val = new(big.Int).Div(&a , &ib)
                    
                }                


                return Number{ Value: &IntNumber{ Value: *val } , IsInt: true }

            }

            fb = &x.Value.(*FloatNumber).Value
            fa = new(big.Float).SetInt(&a)
            switch op{
                    case token.PLUS:
                        val = new(big.Float).Add(fa , fb) 
                    case token.MINUS:
                        val = new(big.Float).Sub(fa , fb)
                    case token.MUL:
                        val = new(big.Float).Mul(fa, fb)
                    case token.DIV:
                        val = new(big.Float).Quo(fa , fb)
                    
                } 
            return Number{ Value: &FloatNumber{ Value: *val } , IsInt: false }
        default:
            return Number{}
    }
}
