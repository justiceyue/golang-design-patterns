package chain

import "testing"

func TestChain(t *testing.T) {
	r := &RouterGroup{}
	r.Use(middleware1, middleware2)
	r.Next()
}
