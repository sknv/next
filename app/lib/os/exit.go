package os

import (
	"os"
	"os/signal"
)

func WaitForExit() {
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)
	<-exit
}
