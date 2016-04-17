/*
 * Copyright 2016, Robert Bieber
 *
 * This file is part of senatron.
 *
 * senatron is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * senatron is distributed in the hope that it will be useful,
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with senatron.  If not, see <http://www.gnu.org/licenses/>.
 */

package sunlight

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURI = "https://congress.api.sunlightfoundation.com"

// buildURI constructs a Congress API URI from the given endpoint and
// query parameters.  Every one of the params must be either a string
// or a slice of strings (for multi-valued parameters).  Passing any
// other type will return an error.
func buildURI(
	endpoint string,
	params map[string]interface{},
) (*url.URL, error) {
	uri, err := url.Parse(baseURI + "/" + endpoint)
	if err != nil {
		return nil, err
	}

	processedParams := make(map[string][]string, len(params))
	for k, v := range params {
		switch typed := v.(type) {
		case string:
			processedParams[k] = []string{typed}
		case []string:
			processedParams[k] = typed
		default:
			return nil, fmt.Errorf(
				"sunlight: Can't use type %T as a query parameter",
				typed,
			)
		}
	}
	uri.RawQuery = url.Values(processedParams).Encode()

	return uri, nil
}

func getHTTPClient() *http.Client {
	return &http.Client{Timeout: time.Second * 10}
}

func getRequest(
	apiKey string,
	uri *url.URL,
) (request *http.Request, err error) {
	request, err = http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("X-APIKEY", apiKey)
	return
}
