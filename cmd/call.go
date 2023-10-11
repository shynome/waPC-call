/*
Copyright © 2023 shynome <shynome@gmail.com>
*/
package cmd

import (
	"context"
	"io"
	"os"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
	"github.com/wapc/wapc-go"
	"github.com/wapc/wapc-go/engines/wazero"
)

func getInput(stdin io.Reader) []byte {
	if f, ok := stdin.(*os.File); ok {
		stat, _ := f.Stat()
		if stat.Mode()&os.ModeNamedPipe == 0 {
			return nil
		}
	}
	return try.To1(io.ReadAll(stdin))
}

func call(wasm, operation string, input []byte) (_ []byte, err error) {
	defer err0.Then(&err, nil, nil)

	guset := try.To1(os.ReadFile(wasm))

	engine := wazero.Engine()

	ctx := context.Background()
	module, err := engine.New(ctx, wapc.NoOpHostCallHandler, guset, &wapc.ModuleConfig{
		Logger: wapc.PrintlnLogger,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	try.To(err)
	defer module.Close(ctx)

	instance := try.To1(module.Instantiate(ctx))
	defer instance.Close(ctx)

	result := try.To1(instance.Invoke(ctx, operation, input))

	return result, nil
}
