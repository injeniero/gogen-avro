// Code generated by github.com/injeniero/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"io"

	"github.com/injeniero/gogen-avro/v10/util"
	"github.com/injeniero/gogen-avro/v10/vm/types"
)

func writeComCompanySharedConflictingFixedType(r ComCompanySharedConflictingFixedType, w io.Writer) error {
	_, err := w.Write(r[:])
	return err
}

type ComCompanySharedConflictingFixedTypeWrapper struct {
	Target *ComCompanySharedConflictingFixedType
}

type ComCompanySharedConflictingFixedType [2]byte

func (b *ComCompanySharedConflictingFixedType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	codepoints := util.DecodeByteString(s)
	copy((*b)[:], codepoints)
	return nil
}

func (b ComCompanySharedConflictingFixedType) MarshalJSON() ([]byte, error) {
	return []byte(util.EncodeByteString(b[:])), nil
}

func (_ ComCompanySharedConflictingFixedTypeWrapper) SetBoolean(v bool) {
	panic("Unsupported operation")
}
func (_ ComCompanySharedConflictingFixedTypeWrapper) SetInt(v int32)  { panic("Unsupported operation") }
func (_ ComCompanySharedConflictingFixedTypeWrapper) SetLong(v int64) { panic("Unsupported operation") }
func (_ ComCompanySharedConflictingFixedTypeWrapper) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ ComCompanySharedConflictingFixedTypeWrapper) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (r ComCompanySharedConflictingFixedTypeWrapper) SetBytes(v []byte) {
	copy((*r.Target)[:], v)
}
func (_ ComCompanySharedConflictingFixedTypeWrapper) SetString(v string) {
	panic("Unsupported operation")
}
func (_ ComCompanySharedConflictingFixedTypeWrapper) SetUnionElem(v int64) {
	panic("Unsupported operation")
}
func (_ ComCompanySharedConflictingFixedTypeWrapper) Get(i int) types.Field {
	panic("Unsupported operation")
}
func (_ ComCompanySharedConflictingFixedTypeWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ComCompanySharedConflictingFixedTypeWrapper) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ ComCompanySharedConflictingFixedTypeWrapper) NullField(int) { panic("Unsupported operation") }
func (_ ComCompanySharedConflictingFixedTypeWrapper) HintSize(int)  { panic("Unsupported operation") }
func (_ ComCompanySharedConflictingFixedTypeWrapper) Finalize()     {}
func (_ ComCompanySharedConflictingFixedTypeWrapper) SetDefault(i int) {
	panic("Unsupported operation")
}
