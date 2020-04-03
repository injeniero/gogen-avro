// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)

type UnionNullArrayIntMapIntNestedUnionRecordTypeEnum int

const (
	UnionNullArrayIntMapIntNestedUnionRecordTypeEnumNull UnionNullArrayIntMapIntNestedUnionRecordTypeEnum = 0

	UnionNullArrayIntMapIntNestedUnionRecordTypeEnumArrayInt UnionNullArrayIntMapIntNestedUnionRecordTypeEnum = 1

	UnionNullArrayIntMapIntNestedUnionRecordTypeEnumMapInt UnionNullArrayIntMapIntNestedUnionRecordTypeEnum = 2

	UnionNullArrayIntMapIntNestedUnionRecordTypeEnumNestedUnionRecord UnionNullArrayIntMapIntNestedUnionRecordTypeEnum = 3
)

type UnionNullArrayIntMapIntNestedUnionRecord struct {
	Null *types.NullVal

	ArrayInt []int32

	MapInt *MapInt

	NestedUnionRecord *NestedUnionRecord

	UnionType UnionNullArrayIntMapIntNestedUnionRecordTypeEnum
}

func writeUnionNullArrayIntMapIntNestedUnionRecord(r *UnionNullArrayIntMapIntNestedUnionRecord, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {

	case UnionNullArrayIntMapIntNestedUnionRecordTypeEnumNull:
		return vm.WriteNull(r.Null, w)

	case UnionNullArrayIntMapIntNestedUnionRecordTypeEnumArrayInt:
		return writeArrayInt(r.ArrayInt, w)

	case UnionNullArrayIntMapIntNestedUnionRecordTypeEnumMapInt:
		return writeMapInt(r.MapInt, w)

	case UnionNullArrayIntMapIntNestedUnionRecordTypeEnumNestedUnionRecord:
		return writeNestedUnionRecord(r.NestedUnionRecord, w)

	}
	return fmt.Errorf("invalid value for *UnionNullArrayIntMapIntNestedUnionRecord")
}

func NewUnionNullArrayIntMapIntNestedUnionRecord() *UnionNullArrayIntMapIntNestedUnionRecord {
	return &UnionNullArrayIntMapIntNestedUnionRecord{}
}

func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetInt(v int32)    { panic("Unsupported operation") }
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetString(v string) {
	panic("Unsupported operation")
}
func (r *UnionNullArrayIntMapIntNestedUnionRecord) SetLong(v int64) {
	r.UnionType = (UnionNullArrayIntMapIntNestedUnionRecordTypeEnum)(v)
}
func (r *UnionNullArrayIntMapIntNestedUnionRecord) Get(i int) types.Field {
	switch i {

	case 0:

		return r.Null

	case 1:

		r.ArrayInt = make([]int32, 0)

		return (*ArrayIntWrapper)(&r.ArrayInt)

	case 2:

		r.MapInt = NewMapInt()

		return r.MapInt

	case 3:

		r.NestedUnionRecord = NewNestedUnionRecord()

		return r.NestedUnionRecord

	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayIntMapIntNestedUnionRecord) Finalize() {}
