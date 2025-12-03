package runner

import (
	"sync"
	"time"

	"github.com/ihsanlearn/redirx/internal/input"
	"github.com/ihsanlearn/redirx/internal/options"
	"github.com/ihsanlearn/redirx/pkg/httputils"
	"github.com/ihsanlearn/redirx/pkg/logger"
	"github.com/ihsanlearn/redirx/pkg/scanner"
)

func Run(opts *options.Options) {
	if !opts.Silent {
		logger.PrintBanner()
		logger.Info("Runner initialized. Ready to hunt bugs!")
	}

	inputProvider := input.NewInputProvider(opts)
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

	// can be modified later
	targetPayload := "https://iihn.fun"
	if opts.Payload != "" {
		targetPayload = opts.Payload
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

					result := scanner.ScanUrl(client, url, targetPayload)
					
					if result != nil {
						if opts.Silent {
							logger.Green(result.VulnerableUrl)
						} else {
							logger.Vulnerable("%s", result.VulnerableUrl)
						}
					} else {
						if opts.Verbose {
							logger.NotVulnerable("%s", url)
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