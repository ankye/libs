package signal

import (
	"os"
	"os/signal"

	log "github.com/gonethopper/libs/logs"
)

// InitSignal register signals handler.
func InitSignal() {
	chanSig := make(chan os.Signal, 1)
	defer close(chanSig)
	signal.Notify(chanSig, os.Interrupt, os.Kill)
	sig := <-chanSig
	log.Emergency("recv kill %s", sig.String())
}
