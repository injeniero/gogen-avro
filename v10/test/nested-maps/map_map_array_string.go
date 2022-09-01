// Code generated by github.com/injeniero/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"github.com/injeniero/gogen-avro/v10/vm"
	"github.com/injeniero/gogen-avro/v10/vm/types"
	"io"
)

func writeMapMapArrayString(r map[string]map[string][]string, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = writeMapArrayString(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapMapArrayStringWrapper struct {
	Target *map[string]map[string][]string
	keys   []string
	values []map[string][]string
}

func (_ *MapMapArrayStringWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapMapArrayStringWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapMapArrayStringWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapMapArrayStringWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapMapArrayStringWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapMapArrayStringWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapMapArrayStringWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapMapArrayStringWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapMapArrayStringWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapMapArrayStringWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapMapArrayStringWrapper) HintSize(s int) {
	if r.keys == nil {
		r.keys = make([]string, 0, s)
		r.values = make([]map[string][]string, 0, s)
	}
}

func (r *MapMapArrayStringWrapper) NullField(_ int) {
	panic("Unsupported operation")
}

func (r *MapMapArrayStringWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapMapArrayStringWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v map[string][]string
	v = make(map[string][]string)

	r.values = append(r.values, v)
	return &MapArrayStringWrapper{Target: &r.values[len(r.values)-1]}
}

func (_ *MapMapArrayStringWrapper) AppendArray() types.Field { panic("Unsupported operation") }
