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

func writeMapArrayMapTestRecord(r map[string][]MapTestRecord, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = writeArrayMapTestRecord(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapArrayMapTestRecordWrapper struct {
	Target *map[string][]MapTestRecord
	keys   []string
	values [][]MapTestRecord
}

func (_ *MapArrayMapTestRecordWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapArrayMapTestRecordWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapArrayMapTestRecordWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapArrayMapTestRecordWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapArrayMapTestRecordWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapArrayMapTestRecordWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapArrayMapTestRecordWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapArrayMapTestRecordWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapArrayMapTestRecordWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapArrayMapTestRecordWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapArrayMapTestRecordWrapper) HintSize(s int) {
	if r.keys == nil {
		r.keys = make([]string, 0, s)
		r.values = make([][]MapTestRecord, 0, s)
	}
}

func (r *MapArrayMapTestRecordWrapper) NullField(_ int) {
	panic("Unsupported operation")
}

func (r *MapArrayMapTestRecordWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapArrayMapTestRecordWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v []MapTestRecord
	v = make([]MapTestRecord, 0)

	r.values = append(r.values, v)
	return &ArrayMapTestRecordWrapper{Target: &r.values[len(r.values)-1]}
}

func (_ *MapArrayMapTestRecordWrapper) AppendArray() types.Field { panic("Unsupported operation") }
