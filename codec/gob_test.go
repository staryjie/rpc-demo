package codec_test

import (
	"fmt"
	"testing"

	"github.com/staryjie/rpc-demo/codec"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	F1 string
	F2 int
}

func TestGob(t *testing.T) {
	should := assert.New(t)

	// 编码测试
	goBytes, err := codec.GobEncode(&TestStruct{F1: "test_f1", F2: 12})
	if should.NoError(err) {
		fmt.Println(goBytes)
	}

	// 接码测试
	obj := TestStruct{}
	err = codec.GobDecode(goBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj)
	}
}
