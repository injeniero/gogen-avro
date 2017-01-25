/*
 * CODE GENERATED AUTOMATICALLY WITH github.com/alanctgardner/gogen-avro
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 *
 * SOURCES:
 *     block.avsc
 *     header.avsc
 */
package container

import (
	"io"
)

type AvroContainerBlock struct {
	NumRecords  int64
	RecordBytes []byte
	Sync        Sync
}

func DeserializeAvroContainerBlock(r io.Reader) (*AvroContainerBlock, error) {
	return readAvroContainerBlock(r)
}

func (r *AvroContainerBlock) Schema() string {
	return "{\"fields\":[\"long\",\"bytes\",{\"name\":\"sync\",\"size\":16,\"type\":\"fixed\"}],\"name\":\"AvroContainerBlock\",\"type\":\"record\"}"
}

func (r *AvroContainerBlock) Serialize(w io.Writer) error {
	return writeAvroContainerBlock(r, w)
}
