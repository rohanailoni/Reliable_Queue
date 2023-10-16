package serialization

import (
	"ReliableQueue/Output"
	"ReliableQueue/model"
	"bytes"
	"encoding/gob"
)

func EventEncoder(event *model.Event) *Output.ByteSlice {
	var buffer bytes.Buffer
	var err error
	enc := gob.NewEncoder(&buffer)
	if err = enc.Encode(event); err != nil {
		//encoding error should have a common one `EncodingError`
		return &Output.ByteSlice{
			Base: Output.Base{
				Error: err,
			},
			Bytes: nil,
		}
	}
	data := buffer.Bytes()
	return &Output.ByteSlice{
		Base: Output.Base{
			Error: err,
		},
		Bytes: data,
	}
}
