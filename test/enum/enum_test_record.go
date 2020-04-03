// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     enum.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

type EnumTestRecord struct {
	EnumField TestEnumType
}

const EnumTestRecordAvroCRC64Fingerprint = "\x8e\x96\x00̛x3\xfa"

func NewEnumTestRecord() *EnumTestRecord {
	return &EnumTestRecord{}
}

func DeserializeEnumTestRecord(r io.Reader) (*EnumTestRecord, error) {
	t := NewEnumTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeEnumTestRecordFromSchema(r io.Reader, schema string) (*EnumTestRecord, error) {
	t := NewEnumTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeEnumTestRecord(r *EnumTestRecord, w io.Writer) error {
	var err error

	err = writeTestEnumType(r.EnumField, w)
	if err != nil {
		return err
	}

	return err
}

func (r *EnumTestRecord) Serialize(w io.Writer) error {
	return writeEnumTestRecord(r, w)
}

func (r *EnumTestRecord) Schema() string {
	return "{\"fields\":[{\"default\":\"testSymbol3\",\"name\":\"EnumField\",\"type\":{\"doc\":\"Test enum\",\"name\":\"TestEnumType\",\"symbols\":[\"TestSymbol1\",\"testSymbol2\",\"testSymbol3\"],\"type\":\"enum\"}}],\"name\":\"EnumTestRecord\",\"type\":\"record\"}"
}

func (r *EnumTestRecord) SchemaName() string {
	return "EnumTestRecord"
}

func (_ *EnumTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *EnumTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *EnumTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *EnumTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *EnumTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *EnumTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *EnumTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *EnumTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnumTestRecord) Get(i int) types.Field {
	switch i {

	case 0:

		return (*types.Int)(&r.EnumField)

	}
	panic("Unknown field index")
}

func (r *EnumTestRecord) SetDefault(i int) {
	switch i {

	case 0:
		r.EnumField = TestEnumTypeTestSymbol3
		return

	}
	panic("Unknown field index")
}

func (_ *EnumTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *EnumTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *EnumTestRecord) Finalize()                        {}

func (_ *EnumTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(EnumTestRecordAvroCRC64Fingerprint)
}
