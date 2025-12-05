package runner

import (
	"sync"
	"time"

	"github.com/ihsanlearn/redirx/internal/input"
	"github.com/ihsanlearn/redirx/internal/options"
	"github.com/ihsanlearn/redirx/internal/output"
	"github.com/ihsanlearn/redirx/payloads"
	"github.com/ihsanlearn/redirx/pkg/httputils"
	"github.com/ihsanlearn/redirx/pkg/logger"
	"github.com/ihsanlearn/redirx/pkg/scanner"
	"github.com/ihsanlearn/redirx/pkg/utils"
)

func Run(opts *options.Options) {
	if !opts.Silent {
		logger.PrintBanner()
	}

	inputProvider := input.NewInputProvider(opts)

	writer, err := output.NewWriter(opts.Output)
	if err != nil {
		logger.Error("%s", err)
		return
	}
	defer writer.Close()

	client := httputils.NewScannerClient(opts)
	urls := inputProvider.StreamURLs()

	var wg sync.WaitGroup

	var ticker *time.Ticker
	if opts.RateLimit > 0 {
		interval := time.Second / time.Duration(opts.RateLimit)
		ticker = time.NewTicker(interval)

		if !opts.Silent {
			logger.Info("Rate limit set to %d requests per second", opts.RateLimit)
		}
	}

	var targetPayloads []string
	if opts.PayloadList != "" {
		logger.Info("Using payload list %s", opts.PayloadList)
		var err error
		targetPayloads, err = utils.ReadFileLines(opts.PayloadList)
		if err != nil {
			logger.Error("%s", err)
			return
		}
	} else if opts.Payload != "" {
		targetPayloads = []string{opts.Payload}
	} else {
		logger.Info("Using default payloads")
		targetPayloads = payloads.GetDefaultPayloads()
	}

	for i := 0; i < opts.Threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urls {
				if ticker != nil {
					<-ticker.C
				}

				if opts.Delay > 0 {
					time.Sleep(time.Duration(opts.Delay) * time.Millisecond)
				}

				url, notOk := httputils.ProbeURL(url, opts.Timeout)
				if notOk != "" {
					if opts.Verbose {
						logger.Dead("%s", notOk)
					}
					continue
				}

				results := scanner.ScanUrl(client, url, targetPayloads, opts.HPP)

				for _, result := range results {
					if result != nil {
						logger.Vulnerable("%s", result.VulnerableUrl)

						writer.Write(result.VulnerableUrl)
					} else {
						if opts.Verbose {
							logger.NotVulnerable("%s", url)
						}
					}
				}
			}
		}()
	}

	wg.Wait()

	if ticker != nil {
		ticker.Stop()
	}

	if !opts.Silent {
		logger.Info("Scan completed")
	}
}
