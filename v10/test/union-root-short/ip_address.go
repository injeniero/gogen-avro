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

func writeIp_address(r Ip_address, w io.Writer) error {
	_, err := w.Write(r[:])
	return err
}

type Ip_addressWrapper struct {
	Target *Ip_address
}

type Ip_address [16]byte

func (b *Ip_address) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	codepoints := util.DecodeByteString(s)
	copy((*b)[:], codepoints)
	return nil
}

func (b Ip_address) MarshalJSON() ([]byte, error) {
	return []byte(util.EncodeByteString(b[:])), nil
}

func (_ Ip_addressWrapper) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ Ip_addressWrapper) SetInt(v int32)      { panic("Unsupported operation") }
func (_ Ip_addressWrapper) SetLong(v int64)     { panic("Unsupported operation") }
func (_ Ip_addressWrapper) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ Ip_addressWrapper) SetDouble(v float64) { panic("Unsupported operation") }
func (r Ip_addressWrapper) SetBytes(v []byte) {
	copy((*r.Target)[:], v)
}
func (_ Ip_addressWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ Ip_addressWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ Ip_addressWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ Ip_addressWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Ip_addressWrapper) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Ip_addressWrapper) NullField(int)                    { panic("Unsupported operation") }
func (_ Ip_addressWrapper) HintSize(int)                     { panic("Unsupported operation") }
func (_ Ip_addressWrapper) Finalize()                        {}
func (_ Ip_addressWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
