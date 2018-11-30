// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/lapsang-boys/galvin/typeless"
)

type Selector func (nt typeless.NodeType) bool

var (
	Any = func(t typeless.NodeType) bool { return true }
	Body = func(t typeless.NodeType) bool { return t == typeless.Body }
	File = func(t typeless.NodeType) bool { return t == typeless.File }
	FunctionAbstraction = func(t typeless.NodeType) bool { return t == typeless.FunctionAbstraction }
	FunctionApplication = func(t typeless.NodeType) bool { return t == typeless.FunctionApplication }
	Identifier = func(t typeless.NodeType) bool { return t == typeless.Identifier }
	Literal = func(t typeless.NodeType) bool { return t == typeless.Literal }
	Expression = OneOf(typeless.Expression...)
)

func OneOf(types ...typeless.NodeType) Selector {
	if len(types) == 0 {
		return func(typeless.NodeType) bool { return false }
	}
	const bits = 32
	max := 1
	for _, t := range types {
		if int(t) > max {
			max = int(t)
		}
	}
	size := (max + bits) / bits
	bitarr := make([]uint32, size)
	for _, t := range types {
		bitarr[uint(t)/bits] |= 1 << (uint(t) % bits)
	}
	return func(t typeless.NodeType) bool {
		i := uint(t)/bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}