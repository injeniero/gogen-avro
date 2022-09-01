package avro

import (
	"io"
	"testing"

	"github.com/injeniero/gogen-avro/v10/container"
	"github.com/injeniero/gogen-avro/v10/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTripExactBytes(t,
		func() container.AvroRecord { return &ConflictingNamespaceFixedTestRecord{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeConflictingNamespaceFixedTestRecord(r)
			return &record, err
		})
}
