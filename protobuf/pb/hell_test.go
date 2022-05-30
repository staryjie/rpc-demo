package pb_test

import (
	"fmt"
	"testing"

	"github.com/staryjie/rpc-demo/protobuf/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestMarshal(t *testing.T) {
	should := assert.New(t)
	str := &pb.String{Value: "Hello"}

	// onject --> protoful --> []byte  // 序列化
	pbBytes, err := proto.Marshal(str)
	if should.NoError(err) {
		if should.Equal([]uint8([]byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f}), pbBytes) {
			fmt.Println(string(pbBytes))
		} else {
			fmt.Print("Result is not Hello")
		}
	}

	// []byte --> protoful --> object
	obj := pb.String{}
	err = proto.Unmarshal(pbBytes, &obj)
	if should.NoError(err) {
		if should.Equal("Hello", obj.Value) {
			fmt.Printf("%v\n", obj.Value)
		} else {
			fmt.Print("Result Not Value = 'Hello'")
		}
	}
}
