package options

import (
	"os"

	"github.com/projectdiscovery/goflags"
	"github.com/ihsanlearn/redirx/pkg/logger"
)

type Options struct {
	URLs    string
	URLList string

	Threads     int
	Timeout     int
	Payload     string
	PayloadList string
	HPP         bool
	VerifySSL bool
	JSCheck     bool
    Version     bool
	RateLimit   int
    Delay       int
    KeepAlive   bool

	Output string

	Silent  bool
	Verbose bool
}

func ParseOptions() *Options {
	opts := &Options{}

	flagSet := goflags.NewFlagSet()
	
	flagSet.SetDescription("RedirX is a high-performance Open Redirect scanner written in Go.")

	flagSet.CreateGroup("input", "Input",
		flagSet.StringVarP(&opts.URLs, "url", "u", "", "Target URL for scanning (comma separated)"),
		flagSet.StringVarP(&opts.URLList, "list", "l", "", "File containing list of target URLs"),
	)

	flagSet.CreateGroup("config", "Configuration",
		flagSet.IntVarP(&opts.Threads, "threads", "t", 25, "Number of concurrent threads"),
		flagSet.IntVarP(&opts.Timeout, "timeout", "T", 10, "Timeout request (detik)"),
		flagSet.StringVarP(&opts.Payload, "payload", "p", "", "Custom payload for scanning"),
		flagSet.StringVarP(&opts.PayloadList, "payload-list", "pl", "", "File containing list of custom payloads"),
		flagSet.BoolVarP(&opts.HPP, "hpp", "H", false, "Enable HTTP Parameter Pollution"),
		flagSet.BoolVarP(&opts.VerifySSL, "verify-ssl", "", false, "Disable SSL verification"),
		flagSet.BoolVarP(&opts.JSCheck, "js-check", "", false, "Enable DOM based scan (Experimental)"),
		flagSet.IntVarP(&opts.RateLimit, "rate-limit", "rl", 10, "Maximum requests per second"),
		flagSet.IntVarP(&opts.Delay, "delay", "d", 0, "Delay between requests (milliseconds)"),
		flagSet.BoolVarP(&opts.KeepAlive, "keep-alive", "k", true, "Enable keep-alive connections"),
	)

	flagSet.CreateGroup("output", "Output",
		flagSet.StringVarP(&opts.Output, "output", "o", "", "File for saving scan results"),
	)

	flagSet.CreateGroup("misc", "Optimization",
		flagSet.BoolVarP(&opts.Silent, "silent", "s", false, "Silent mode (hanya print vuln)"),
		flagSet.BoolVarP(&opts.Verbose, "verbose", "v", false, "Verbose mode (print error & debug)"),
        flagSet.BoolVarP(&opts.Version, "version", "V", false, "Display application version"),
	)

	if err := flagSet.Parse(); err != nil {
		logger.Error("Failed parsing flags: %s", err)
		os.Exit(1)
	}
    
    if opts.Version {
        logger.Info("RedirX Version 1.0.0")
        os.Exit(0)
    }

	if opts.URLs == "" && opts.URLList == "" && !hasStdin() {
		logger.Error("No target URL provided! Use -u, -l or pipe stdin.")
        logger.Info("Run 'redirx -h' for help.")
		os.Exit(1)
	}

	return opts
}

func hasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return stat.Mode()&os.ModeCharDevice == 0
}