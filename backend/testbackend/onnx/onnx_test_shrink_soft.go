package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Shrink", "TestShrinkSoft", NewTestShrinkSoft)
}

// NewTestShrinkSoft version: 4.
func NewTestShrinkSoft() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Shrink",
		Title:  "TestShrinkSoft",
		ModelB: []byte{0x8, 0x4, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x65, 0xa, 0x2f, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x6, 0x53, 0x68, 0x72, 0x69, 0x6e, 0x6b, 0x2a, 0xe, 0xa, 0x4, 0x62, 0x69, 0x61, 0x73, 0x15, 0x0, 0x0, 0xc0, 0x3f, 0xa0, 0x1, 0x1, 0x2a, 0xf, 0xa, 0x5, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x15, 0x0, 0x0, 0xc0, 0x3f, 0xa0, 0x1, 0x1, 0x12, 0x10, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x68, 0x72, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x6f, 0x66, 0x74, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0xa},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Shrink",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc0000c69a0)(name:"bias" type:FLOAT f:1.5 ),
		    (*ir.AttributeProto)(0xc0000c6a80)(name:"lambd" type:FLOAT f:1.5 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(5),
				tensor.WithBacking([]float32{-2, -1, 0, 1, 2}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(5),
				tensor.WithBacking([]float32{-0.5, 0, 0, 0, 0.5}),
			),
		},
	}
}
