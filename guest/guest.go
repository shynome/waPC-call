package main

import (
	"fmt"

	"github.com/shynome/waPC-call/guest/model"

	"github.com/wapc/wapc-guest-tinygo"
)

//go:generate tinygo build -target wasi -no-debug -gc=leaking -scheduler=none

func main() {
	wapc.RegisterFunctions(wapc.Functions{
		"hello": hello,
	})
}

func hello(payload []byte) ([]byte, error) {
	var live model.Input
	if err := live.UnmarshalJSON(payload); err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("hello %s", live.Name)
	return []byte(msg), nil
}
