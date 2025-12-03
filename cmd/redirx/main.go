package main

import (
	"github.com/ihsanlearn/redirx/internal/options"
	"github.com/ihsanlearn/redirx/internal/runner"
)

func main() {
	opts := options.ParseOptions()

	runner.Run(opts)
}