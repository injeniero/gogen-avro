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

type UnionRecord struct {
	Id string `json:"id"`

	Age int32 `json:"age"`
}

const UnionRecordAvroCRC64Fingerprint = "XXv\f\x93)\xf7\xe2"

func NewUnionRecord() UnionRecord {
	r := UnionRecord{}
	r.Id = "test_id"
	r.Age = 100
	return r
}

func DeserializeUnionRecord(r io.Reader) (UnionRecord, error) {
	t := NewUnionRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeUnionRecordFromSchema(r io.Reader, schema string) (UnionRecord, error) {
	t := NewUnionRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeUnionRecord(r UnionRecord, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Age, w)
	if err != nil {
		return err
	}
	return err
}

func (r UnionRecord) Serialize(w io.Writer) error {
	return writeUnionRecord(r, w)
}

func (r UnionRecord) Schema() string {
	return "{\"fields\":[{\"default\":\"test_id\",\"name\":\"id\",\"type\":\"string\"},{\"default\":100,\"name\":\"age\",\"type\":\"int\"}],\"name\":\"UnionRecord\",\"type\":\"record\"}"
}

func (r UnionRecord) SchemaName() string {
	return "UnionRecord"
}

func (_ UnionRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ UnionRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ UnionRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ UnionRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ UnionRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ UnionRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ UnionRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ UnionRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *UnionRecord) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		w := types.Int{Target: &r.Age}

		return w

	}
	panic("Unknown field index")
}

func (r *UnionRecord) SetDefault(i int) {
	switch i {
	case 0:
		r.Id = "test_id"
		return
	case 1:
		r.Age = 100
		return
	}
	panic("Unknown field index")
}

func (r *UnionRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ UnionRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ UnionRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ UnionRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ UnionRecord) Finalize()                        {}

func (_ UnionRecord) AvroCRC64Fingerprint() []byte {
	return []byte(UnionRecordAvroCRC64Fingerprint)
}

func (r UnionRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["age"], err = json.Marshal(r.Age)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *UnionRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		r.Id = "test_id"
	}
	val = func() json.RawMessage {
		if v, ok := fields["age"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Age); err != nil {
			return err
		}
	} else {
		r.Age = 100
	}
	return nil
}
