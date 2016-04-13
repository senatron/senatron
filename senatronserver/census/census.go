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

package census

import (
	"errors"
)

// Based on 2014 census estimates from
// www.census.gov/popest/data/state/totals/2014/tables/NST-EST2014-01.csv
var populations = map[string]int{
	"AK": 736732,
	"AL": 4849377,
	"AR": 2966369,
	"AZ": 6731484,
	"CA": 38802500,
	"CO": 5355866,
	"CT": 3596677,
	"DC": 658893,
	"DE": 935614,
	"FL": 19893297,
	"GA": 10097343,
	"HI": 1419561,
	"IA": 3107126,
	"ID": 1634464,
	"IL": 12880580,
	"IN": 6596855,
	"KS": 2904021,
	"KY": 4413457,
	"LA": 4649676,
	"MA": 6745408,
	"MD": 5976407,
	"ME": 1330089,
	"MI": 9909877,
	"MN": 5457173,
	"MO": 6063589,
	"MS": 2994079,
	"MT": 1023579,
	"NC": 9943964,
	"ND": 739482,
	"NE": 1881503,
	"NH": 1326813,
	"NJ": 8938175,
	"NM": 2085572,
	"NV": 2839099,
	"NY": 19746227,
	"OH": 11594163,
	"OK": 3878051,
	"OR": 3970239,
	"PA": 12787209,
	"PR": 3548397,
	"RI": 1055173,
	"SC": 4832482,
	"SD": 853175,
	"TN": 6549352,
	"TX": 26956958,
	"UT": 2942902,
	"VA": 8326289,
	"VT": 626562,
	"WA": 7061530,
	"WI": 5757564,
	"WV": 1850326,
	"WY": 584153,
}

// Get returns the population of the given state (by capitalized,
// two-letter state code), or 0 and an error if the code is invalid.
func Get(state string) (int, error) {
	if pop, ok := populations[state]; ok {
		return pop, nil
	}
	return 0, errors.New("State not found")
}

// AllStates returns a full list of available state codes.
func AllStates() []string {
	out := make([]string, len(populations))

	i := 0
	for k := range populations {
		out[i] = k
		i++
	}

	return out
}
