package gorgonnx

import (
        "github.com/owulveryck/onnx-go"
        "gorgonia.org/gorgonia"
//        "gorgonia.org/tensor"
        "fmt"
)

type SliceO struct{
        //inShape tensor.Shape
        //toShape tensor.Shape
        //starts tensor.Tensor
        //slId []int
}

func init() {
        register("Slice", newSlice)
}

func newSlice() operator {
        return &SliceO{}
}

func (a *SliceO) apply(g *Graph, ns ...*Node) error {
        n := ns[0]
        children := getOrderedChildren(g.g, n)
        err := checkCondition(children, 5)
        if err != nil {
                return err
        }
        chomp := children[2].gorgoniaNode.Value().Data().([]int64)[0]
        size := children[0].gorgoniaNode.Shape()[3]
        n.gorgoniaNode, err = gorgonia.Slice(children[0].gorgoniaNode, nil, nil, nil, gorgonia.S(0, int(int64(size)+chomp)))
        return err
}

func (a *SliceO) init(o onnx.Operation) error {
        return nil
}

