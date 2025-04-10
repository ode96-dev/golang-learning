package main

import (
	"log/slog"
	"os"
)

/*
log/slog
- Go's new official structured logging
	- structured fields(key-value)
	- JSON or text input
	- log levels
	- log levels - debug, info, warn, error
	- context aware logging
*/

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("user login", "user", "admin", "role", "superuser")
}
