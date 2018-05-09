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

func TestGenUID(t *testing.T) {
	uid, err := GenUID(1)
	if err != nil {
		t.Error("failed to gen uid")
	}
	if uid <= 0 {
		t.Error("uid must > 0")
	}

}
