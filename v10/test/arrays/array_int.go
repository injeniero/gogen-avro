// Code generated by github.com/injeniero/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/injeniero/gogen-avro/v10/vm"
	"github.com/injeniero/gogen-avro/v10/vm/types"
)

func writeArrayInt(r []int32, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteInt(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayIntWrapper struct {
	Target *[]int32
}

func (_ ArrayIntWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayIntWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayIntWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayIntWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayIntWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayIntWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayIntWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayIntWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayIntWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayIntWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayIntWrapper) Finalize()                        {}
func (_ ArrayIntWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayIntWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]int32, 0, s)
	}
}
func (r ArrayIntWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayIntWrapper) AppendArray() types.Field {
	var v int32

	*r.Target = append(*r.Target, v)
	return &types.Int{Target: &(*r.Target)[len(*r.Target)-1]}
}
