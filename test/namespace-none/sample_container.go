// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
)

func NewSampleWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewSample()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type SampleReader struct {
	r io.Reader
	p *vm.Program
}

func NewSampleReader(r io.Reader) (*SampleReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewSample()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &SampleReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r SampleReader) Read() (*Sample, error) {
	t := NewSample()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
