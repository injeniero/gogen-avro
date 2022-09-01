// Code generated by github.com/injeniero/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/injeniero/gogen-avro/v10/compiler"
	"github.com/injeniero/gogen-avro/v10/container"
	"github.com/injeniero/gogen-avro/v10/vm"
)

func NewComCompanyTeamTypeOneWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewComCompanyTeamTypeOne()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type ComCompanyTeamTypeOneReader struct {
	r io.Reader
	p *vm.Program
}

func NewComCompanyTeamTypeOneReader(r io.Reader) (*ComCompanyTeamTypeOneReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewComCompanyTeamTypeOne()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &ComCompanyTeamTypeOneReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r ComCompanyTeamTypeOneReader) Read() (ComCompanyTeamTypeOne, error) {
	t := NewComCompanyTeamTypeOne()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
