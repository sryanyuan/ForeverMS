package gosync

import (
	"os"
	"os/signal"
)

func WaitSignals(sig ...os.Signal) os.Signal {
	// Wait for signals to quit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, sig...)
	return <-sigCh
}
