package number

import "math/big"

func (n *Number) Add(x Number) Number{
    switch n.GetType(){
        case "FLOAT":
            a := n.Value.(*FloatNumber).Value
            if x.IsInt{
                b := x.Value.(*IntNumber).Value

                fb := new(big.Float).SetInt(&b)

                val := a.Add(&a,  fb)

                return Number{ Value: &FloatNumber{ Value: *val } , IsInt: false }
            }

            b := x.Value.(*FloatNumber).Value
            val := a.Add(&a , &b)

            return Number{ Value: &FloatNumber{ Value: *val } , IsInt: false }
        case "INT":
            a := n.Value.(*IntNumber).Value

            if x.IsInt{
                b := x.Value.(*IntNumber).Value
                val := a.Add(&a,&b)
                return Number{ Value: &IntNumber{ Value: *val } , IsInt: true }

            }

            b := x.Value.(*FloatNumber).Value
            fa := new(big.Float).SetInt(&a)
            val := fa.Add(fa , &b)
            return Number{ Value: &FloatNumber{ Value: *val } , IsInt: false }
        default:
            return Number{}
    }
}
