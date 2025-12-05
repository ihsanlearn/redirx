package payloads

import (
	_ "embed"
	"strings"
)

//go:embed payloads.txt
var defaultPayloads string

func GetDefaultPayloads() []string {
	var list []string

	lines := strings.Split(strings.ReplaceAll(defaultPayloads, "\r\n", "\n"), "\n")

	for _, line := range lines {
		clean := strings.TrimSpace(line)
		if clean != "" {
			list = append(list, clean)
		}
	}

	return list
}