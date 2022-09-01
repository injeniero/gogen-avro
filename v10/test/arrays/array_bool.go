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

func writeArrayBool(r []bool, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteBool(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayBoolWrapper struct {
	Target *[]bool
}

func (_ ArrayBoolWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayBoolWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayBoolWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayBoolWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayBoolWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayBoolWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayBoolWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayBoolWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayBoolWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayBoolWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayBoolWrapper) Finalize()                        {}
func (_ ArrayBoolWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayBoolWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]bool, 0, s)
	}
}
func (r ArrayBoolWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayBoolWrapper) AppendArray() types.Field {
	var v bool

	*r.Target = append(*r.Target, v)
	return &types.Boolean{Target: &(*r.Target)[len(*r.Target)-1]}
}
