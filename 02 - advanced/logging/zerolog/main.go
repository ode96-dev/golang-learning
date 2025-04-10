package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

/*
zerolog
- built for high performance structured logging, optimized for speed
	- very fast
	- no memory allocation
	- simple chaining API
	- json-first
*/

func main() {
	// example
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	log := zerolog.New(output).With().Timestamp().Logger()

	log.Info().Str("foo", "bar").Msg("Hello World")

}
