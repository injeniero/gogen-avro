// Code generated by github.com/injeniero/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     parent.avsc
 *     child.avsc
 */
package avro

import (
	"io"

	"github.com/injeniero/gogen-avro/v10/compiler"
	"github.com/injeniero/gogen-avro/v10/container"
	"github.com/injeniero/gogen-avro/v10/vm"
)

func NewChildWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewChild()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type ChildReader struct {
	r io.Reader
	p *vm.Program
}

func NewChildReader(r io.Reader) (*ChildReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewChild()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &ChildReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r ChildReader) Read() (Child, error) {
	t := NewChild()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
