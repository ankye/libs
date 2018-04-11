package utils

import (
	"testing"

	log "github.com/gonethopper/libs/logs"
)

func TestGenUUID(t *testing.T) {
	u, err := GenUUID()
	if err != nil {
		t.Error("failed to gen uuid")
	}
	log.Info("get uuid %s", u)
}
