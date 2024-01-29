//go:generate protoc -I=proto -I=$GOPATH/src -I=$GOPATH/src/github.com/dharitri/protobuf/protobuf  --gogoslick_out=. dct.proto
package dct

import (
    "math/big"

	dct  "github.com/subrahamanyam341/andes-abc-01/data/dct/proto"
)

// New returns a new batch from given buffers
func New() *dct.ESDigitalToken {
	return &dct.ESDigitalToken{
		Value: big.NewInt(0),
	}
}
