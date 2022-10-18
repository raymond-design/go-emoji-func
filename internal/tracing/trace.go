package tracing

import (
	"context"
	"fmt"
	"os"
	"runtime/trace"

	"github.com/raymond-design/go-emoji-func/internal/debug"
)

func Trace(b context.Context) (context.Context, func(), error) {
	if !debug.Traced {
		return b, func() {}, nil
	}

	fi, err := os.Create(debug.Debug.Trace)

	if err != nil {
		return b, func() {}, fmt.Errorf("couldn't output trace: %v\n", err)
	}

	if err := trace.Start(fi); err != nil {
		return b, func() {}, fmt.Errorf("couldn't start trace: %v\n", err)
	}

	ctx, t := trace.NewTask(b, "gef")

	return ctx, func() {
		defer fi.Close()
		defer trace.Stop()
		defer t.End()
	}, nil
}