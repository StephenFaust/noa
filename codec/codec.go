package codec

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Codec interface {
	Encode(w io.Writer, data *bytes.Buffer) (err error)
	Decode(r io.Reader) (data *bytes.Buffer, err error)
}

var DefaultCodec = LengthSplitCodec{}

type LengthSplitCodec struct {
}

// Encode 编码
func (codec LengthSplitCodec) Encode(w io.Writer, data *bytes.Buffer) (err error) {
	l := data.Len()
	var size [binary.MaxVarintLen64]byte
	if l == 0 {
		n := binary.PutUvarint(size[:], uint64(0))
		_, err := w.Write(size[:n])
		if err != nil {
			return err
		}
		return nil
	}
	n := binary.PutUvarint(size[:], uint64(l))
	_, err = w.Write(size[:n])
	if err != nil {
		return err
	}
	_, err = io.Copy(w, data)
	if err != nil {
		return err
	}
	return err
}

func (codec LengthSplitCodec) Decode(r io.Reader) (data *bytes.Buffer, err error) {
	size, err := binary.ReadUvarint(r.(io.ByteReader))
	if err != nil {
		return nil, err
	}
	if size != 0 {
		data = bytes.NewBuffer(make([]byte, size))
		for index := 0; index < data.Len(); {
			n, err := r.Read(data.Bytes()[index:])
			if err != nil {
				return nil, err
			}
			index += n
		}
		return data, nil
	}
	return nil, err
}
