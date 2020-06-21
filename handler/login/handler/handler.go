package handler

import "github.com/treeforest/gos/transport"

// 服务名
type Test struct {
	// something
}

// 此处的入参应该是对应的 message
func (h *Test) SayHello(req transport.Request) {
	// do something
}

// 此处的入参应该是对应的 message
func (h *Test) SayHi(req transport.Request) {
	// do something
}

// 此处的入参应该是对应的 message
func (h *Test) Jump(req transport.Request) {
	// do something
}
