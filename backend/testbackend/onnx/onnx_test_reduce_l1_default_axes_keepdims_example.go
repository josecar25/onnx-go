package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ReduceL1", "TestReduceL1DefaultAxesKeepdimsExample", NewTestReduceL1DefaultAxesKeepdimsExample)
}

// NewTestReduceL1DefaultAxesKeepdimsExample version: 3.
func NewTestReduceL1DefaultAxesKeepdimsExample() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ReduceL1",
		Title:  "TestReduceL1DefaultAxesKeepdimsExample",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x95, 0x1, 0xa, 0x2a, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x7, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x22, 0x8, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x4c, 0x31, 0x2a, 0xf, 0xa, 0x8, 0x6b, 0x65, 0x65, 0x70, 0x64, 0x69, 0x6d, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x5f, 0x6c, 0x31, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x78, 0x65, 0x73, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x64, 0x69, 0x6d, 0x73, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5a, 0x1a, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x62, 0x1d, 0xa, 0x7, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"data"},
		     Output:    []string{"reduced"},
		     Name:      "",
		     OpType:    "ReduceL1",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc0000c6b60)(name:"keepdims" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 2, 2),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 1),
				tensor.WithBacking([]float32{78}),
			),
		},
	}
}
