package codec

import (
	"bytes"
	"encoding/gob"
)

// object --> gob --> []byte
func GobEncode(obj interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(obj); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

// []byte --> gob --> object
func GobDecode(data []byte, obj interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(data))

	return decoder.Decode(obj)
}
