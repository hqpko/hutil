package hutils

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitExitSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
