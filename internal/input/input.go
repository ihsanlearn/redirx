package input

import (
	"bufio"
	"os"
	"strings"

	"github.com/ihsanlearn/redirx/internal/options"
	"github.com/ihsanlearn/redirx/pkg/logger"
)

type InputProvider struct {
	Options *options.Options
}

func NewInputProvider(options *options.Options) *InputProvider {
	return &InputProvider{Options: options}
}

func (i *InputProvider) StreamURLs() chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		seen := make(map[string]bool)

		send := func(raw string) {
			clean := strings.TrimSpace(raw)
			if clean == "" {
				return
			}

			if !strings.HasPrefix(clean, "http://") && !strings.HasPrefix(clean, "https://") {
				clean = "http://" + clean
			}

			if seen[clean] {
				return
			}

			seen[clean] = true

			out <- clean
		}

		if hasStdin() {
			logger.Info("Reading URLs from stdin")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				url := strings.TrimSpace(scanner.Text())
				if url != "" {
					send(url)
				}
			}
		}

		if i.Options.URLList != "" {
			logger.Info("Reading URLs from file %s", i.Options.URLList)
			file, err := os.Open(i.Options.URLList)
			if err != nil {
				logger.Error("Failed to open file %s", i.Options.URLList)
				os.Exit(1)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			buf := make([]byte, 0, 1024*1024)
			scanner.Buffer(buf, 50*1024*1024)

			for scanner.Scan() {
				url := strings.TrimSpace(scanner.Text())
				if url != "" {
					send(url)
				}
			}
		}

		if i.Options.URLs != "" {
			logger.Info("Reading URL from flag %s", i.Options.URLs)
			urls := strings.Split(i.Options.URLs, ",")
			for _, url := range urls {
				if url != "" {
					send(url)
				}
			}
		}
	}()

	return out
}

func hasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return stat.Mode()&os.ModeCharDevice == 0
}