package avro

import (
	"io"
	"testing"

	"github.com/injeniero/gogen-avro/v10/container"
	"github.com/injeniero/gogen-avro/v10/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTrip(t,
		func() container.AvroRecord { return &ComplexUnionTestRecord{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeComplexUnionTestRecord(r)
			return &record, err
		})
}
