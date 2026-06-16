package httpclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func BuildURL(userURL string,queries []string) (string, error) {

	parsedURL, err := url.Parse(userURL)

	if err != nil {

		return "", err

	}

	params := parsedURL.Query()

	for _, query := range queries {

		parts := strings.SplitN(query, "=", 2)

		if len(parts) != 2 {

			fmt.Printf("Invalid query format: %sn", query)

			continue

		}

		key := strings.TrimSpace(parts[0])

		value := strings.TrimSpace(parts[1])

		params.Set(key, value)

	}

	parsedURL.RawQuery = params.Encode()

	return parsedURL.String(), nil

}
func ApplyHeaders(req *http.Request,  headers []string) {
	// parse header
	// add header
	for _, header := range headers {
		parts := strings.SplitN(header, ":", 2)
		if len(parts) != 2 {
			fmt.Printf("Invalid header format: %s\n", header)
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		req.Header.Set(key, value)
	}
}