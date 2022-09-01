// Code generated by github.com/injeniero/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/injeniero/gogen-avro/v10/vm"
	"github.com/injeniero/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

// Test  enum
type TestEnumType int32

const (
	TestEnumTypeTestSymbol3 TestEnumType = 0
	TestEnumTypeTestSymbol1 TestEnumType = 1
)

func (e TestEnumType) String() string {
	switch e {
	case TestEnumTypeTestSymbol3:
		return "TestSymbol3"
	case TestEnumTypeTestSymbol1:
		return "TestSymbol1"
	}
	return "unknown"
}

func writeTestEnumType(r TestEnumType, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewTestEnumTypeValue(raw string) (r TestEnumType, err error) {
	switch raw {
	case "TestSymbol3":
		return TestEnumTypeTestSymbol3, nil
	case "TestSymbol1":
		return TestEnumTypeTestSymbol1, nil
	}

	return TestEnumTypeTestSymbol1, nil

}

func (b TestEnumType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *TestEnumType) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewTestEnumTypeValue(stringVal)
	*b = val
	return err
}

type TestEnumTypeWrapper struct {
	Target *TestEnumType
}

func (b TestEnumTypeWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b TestEnumTypeWrapper) SetInt(v int32) {
	*(b.Target) = TestEnumType(v)
}

func (b TestEnumTypeWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b TestEnumTypeWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b TestEnumTypeWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b TestEnumTypeWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b TestEnumTypeWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b TestEnumTypeWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b TestEnumTypeWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b TestEnumTypeWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b TestEnumTypeWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b TestEnumTypeWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b TestEnumTypeWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b TestEnumTypeWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b TestEnumTypeWrapper) Finalize() {}
