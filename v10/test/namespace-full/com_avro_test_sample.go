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

// GoGen test
type ComAvroTestSample struct {
	// Core data information required for any event
	Header *UnionNullHeaderworksData `json:"header"`
	// Core data information required for any event
	Body *UnionNullBodyworksData `json:"body"`
}

const ComAvroTestSampleAvroCRC64Fingerprint = "\xdf}\x93 \x19f\x18\n"

func NewComAvroTestSample() ComAvroTestSample {
	r := ComAvroTestSample{}
	r.Header = nil
	r.Body = nil
	return r
}

func DeserializeComAvroTestSample(r io.Reader) (ComAvroTestSample, error) {
	t := NewComAvroTestSample()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeComAvroTestSampleFromSchema(r io.Reader, schema string) (ComAvroTestSample, error) {
	t := NewComAvroTestSample()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeComAvroTestSample(r ComAvroTestSample, w io.Writer) error {
	var err error
	err = writeUnionNullHeaderworksData(r.Header, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBodyworksData(r.Body, w)
	if err != nil {
		return err
	}
	return err
}

func (r ComAvroTestSample) Serialize(w io.Writer) error {
	return writeComAvroTestSample(r, w)
}

func (r ComAvroTestSample) Schema() string {
	return "{\"doc\":\"GoGen test\",\"fields\":[{\"default\":null,\"doc\":\"Core data information required for any event\",\"name\":\"header\",\"type\":[\"null\",{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",\"headerworks.datatype.UUID\"]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"Data\",\"namespace\":\"headerworks\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Core data information required for any event\",\"name\":\"body\",\"type\":[\"null\",{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"bodyworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",\"headerworks.datatype.UUID\"]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"Data\",\"namespace\":\"bodyworks\",\"type\":\"record\"}]}],\"name\":\"com.avro.test.sample\",\"type\":\"record\"}"
}

func (r ComAvroTestSample) SchemaName() string {
	return "com.avro.test.sample"
}

func (_ ComAvroTestSample) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetString(v string)   { panic("Unsupported operation") }
func (_ ComAvroTestSample) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ComAvroTestSample) Get(i int) types.Field {
	switch i {
	case 0:
		r.Header = NewUnionNullHeaderworksData()

		return r.Header
	case 1:
		r.Body = NewUnionNullBodyworksData()

		return r.Body
	}
	panic("Unknown field index")
}

func (r *ComAvroTestSample) SetDefault(i int) {
	switch i {
	case 0:
		r.Header = nil
		return
	case 1:
		r.Body = nil
		return
	}
	panic("Unknown field index")
}

func (r *ComAvroTestSample) NullField(i int) {
	switch i {
	case 0:
		r.Header = nil
		return
	case 1:
		r.Body = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ComAvroTestSample) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ComAvroTestSample) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ComAvroTestSample) HintSize(int)                     { panic("Unsupported operation") }
func (_ ComAvroTestSample) Finalize()                        {}

func (_ ComAvroTestSample) AvroCRC64Fingerprint() []byte {
	return []byte(ComAvroTestSampleAvroCRC64Fingerprint)
}

func (r ComAvroTestSample) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["header"], err = json.Marshal(r.Header)
	if err != nil {
		return nil, err
	}
	output["body"], err = json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ComAvroTestSample) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["header"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Header); err != nil {
			return err
		}
	} else {
		r.Header = NewUnionNullHeaderworksData()

		r.Header = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["body"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Body); err != nil {
			return err
		}
	} else {
		r.Body = NewUnionNullBodyworksData()

		r.Body = nil
	}
	return nil
}
