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
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// Vote describes the outcome of a vote, both in terms of senate and
// popular vote, as well as the senators and their votes and basic
// info like the bill ID, date and so on.
type Vote struct {
	RollID   string `json:"roll_id"`
	VotedAt  string `json:"voted_at"`
	RollType string `json:"roll_type"`
	Question string `json:"question"`
	Required string `json:"required"`
	Result   string `json:"result"`

	// At most one of BillID or NominationID will be non-empty.
	BillID       string `json:"bill_id"`
	NominationID string `json:"nomination_id"`

	Voters map[string]struct {
		Vote string `json:"vote"`
		Info struct {
			BioguideID string `json:"bioguide_id"`
			State      string `json:"state"`
			Party      string `json:"party"`
		} `json:"voter"`
	} `json:"voters"`
}

// ErrVoteNotFound signals a failure in looking up a given vote,
// probably because no vote by the given roll ID exists.
var ErrVoteNotFound = errors.New("No results for that roll ID")

// GetVote returns information about the given rollID, or returns an
// error if anything goes wrong.
func GetVote(apiKey string, rollID string) (vote Vote, err error) {
	fields := strings.Join(
		[]string{
			"roll_id",
			"bill_id",
			"nomination_id",
			"roll_type",
			"question",
			"required",
			"result",
			"voted_at",
			"voters",
		},
		",",
	)
	uri, err := buildURI(
		"votes",
		map[string]interface{}{
			"roll_id": rollID,
			"fields":  fields,
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

	resultContainer := struct {
		Results []Vote `json:"results"`
		Count   int    `json:"count"`
	}{}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&resultContainer)
	if err != nil {
		return
	}

	if resultContainer.Count == 0 {
		err = ErrVoteNotFound
		return
	} else if resultContainer.Count != 1 {
		// This should never happen
		err = errors.New("More than one vote found for a single roll ID")
		return
	}

	vote = resultContainer.Results[0]
	return
}
