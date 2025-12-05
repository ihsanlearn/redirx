package scanner

import (
	"net/http"
	"net/url"
	"strings"
)

type Result struct {
	VulnerableUrl string
	Payload       string
	RedirectTo    string
	Param         string
}

func ScanUrl(client *http.Client, targetURL string, payloads []string) []*Result {
	u, err := url.Parse(targetURL)
	if err != nil {
		return nil
	}

	queryParams := u.Query()
	if len(queryParams) == 0 {
		return nil
	}

	var findings []*Result

	for param := range queryParams {
		fuzzedParams := make(url.Values)
		for k, v := range queryParams {
			fuzzedParams[k] = v
		}

		for _, payload := range payloads {
			fuzzedParams.Set(param, payload)

			u.RawQuery = fuzzedParams.Encode()
			finalURL := u.String()

			req, err := http.NewRequest("GET", finalURL, nil)
			if err != nil {
				continue
			}
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
		
			resp, err := client.Do(req)
			if err != nil {
				continue
			}
			resp.Body.Close()
		
			if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
				location, err := resp.Location()
				if err != nil {
					continue
				}
	
				locStr := location.String()
				if strings.Contains(locStr, payload) {
					findings = append(findings, &Result{
						VulnerableUrl: finalURL,
						Payload:       payload,
						RedirectTo:    locStr,
						Param:         param,
					})

					break
				}
			}
		}
	}

	return findings
}