// Code generated by github.com/injeniero/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/injeniero/gogen-avro/v10/compiler"
	"github.com/injeniero/gogen-avro/v10/vm"
	"github.com/injeniero/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type RecursiveUnionTestRecord struct {
	RecursiveField *UnionNullRecursiveUnionTestRecord `json:"RecursiveField"`
}

const RecursiveUnionTestRecordAvroCRC64Fingerprint = "\xc6U)C\v\x8a\xa6\x89"

func NewRecursiveUnionTestRecord() RecursiveUnionTestRecord {
	r := RecursiveUnionTestRecord{}
	return r
}

func DeserializeRecursiveUnionTestRecord(r io.Reader) (RecursiveUnionTestRecord, error) {
	t := NewRecursiveUnionTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRecursiveUnionTestRecordFromSchema(r io.Reader, schema string) (RecursiveUnionTestRecord, error) {
	t := NewRecursiveUnionTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRecursiveUnionTestRecord(r RecursiveUnionTestRecord, w io.Writer) error {
	var err error
	err = writeUnionNullRecursiveUnionTestRecord(r.RecursiveField, w)
	if err != nil {
		return err
	}
	return err
}

func (r RecursiveUnionTestRecord) Serialize(w io.Writer) error {
	return writeRecursiveUnionTestRecord(r, w)
}

func (r RecursiveUnionTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"RecursiveField\",\"type\":[\"null\",\"RecursiveUnionTestRecord\"]}],\"name\":\"RecursiveUnionTestRecord\",\"type\":\"record\"}"
}

func (r RecursiveUnionTestRecord) SchemaName() string {
	return "RecursiveUnionTestRecord"
}

func (_ RecursiveUnionTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RecursiveUnionTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RecursiveUnionTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RecursiveUnionTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RecursiveUnionTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RecursiveUnionTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RecursiveUnionTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ RecursiveUnionTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RecursiveUnionTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.RecursiveField = NewUnionNullRecursiveUnionTestRecord()

		return r.RecursiveField
	}
	panic("Unknown field index")
}

func (r *RecursiveUnionTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RecursiveUnionTestRecord) NullField(i int) {
	switch i {
	case 0:
		r.RecursiveField = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ RecursiveUnionTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RecursiveUnionTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RecursiveUnionTestRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ RecursiveUnionTestRecord) Finalize()                        {}

func (_ RecursiveUnionTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(RecursiveUnionTestRecordAvroCRC64Fingerprint)
}

func (r RecursiveUnionTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["RecursiveField"], err = json.Marshal(r.RecursiveField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RecursiveUnionTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["RecursiveField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RecursiveField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for RecursiveField")
	}
	return nil
}
