package sigmon

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, signals ...os.Signal) error {
	sig := make(chan os.Signal, 1)
	if len(signals) == 0 {
		signals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}
	signal.Notify(sig, signals...)
	select {
	case <-sig:
		return errors.New("monitored signal")
	case <-ctx.Done():
		return ctx.Err()
	}
}
