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
	"errors"
	"io/ioutil"
	"net/http"
)

// Vote describes the outcome of a vote, both in terms of senate and
// popular vote, as well as the senators and their votes and basic
// info like the bill ID, date and so on.
type Vote struct {
	RawOutput string
}

// GetVote returns information about the given rollID, or returns an
// error if anything goes wrong.
func GetVote(apiKey string, rollID string) (vote Vote, err error) {
	uri, err := buildURI(
		"votes",
		map[string]interface{}{
			"roll_id": rollID,
		},
	)
	if err != nil {
		return
	}

	request, err := getRequest(apiKey, uri)
	client := getHTTPClient()

	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		err = errors.New(response.Status)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	vote.RawOutput = string(body)

	return
}
