package reconciler

import (
	"log/slog"
	"time"
)

func Manager() {
	for {
		slog.Info("This is reconciler manager ")
		time.Sleep(60 * time.Second)
	}
}
