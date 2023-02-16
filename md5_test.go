package cryptor_test

import (
	"testing"

	"github.com/goexl/cryptor"
)

type md5test struct {
	in       string
	expected string
}

func TestMd5(t *testing.T) {
	tests := []md5test{
		{in: "storezhang", expected: "7081bfba86d133616da630946a617cf4"},
		{in: "store", expected: "8cd892b7b97ef9489ae4479d3f4ef0fc"},
		{in: "-=)(*&^%$#@!~", expected: "28452befda9ffac2f03a6d6babc93c1a"},
	}
	for index, test := range tests {
		got := cryptor.New(test.in).Md5().Hex()
		if test.expected != got {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
