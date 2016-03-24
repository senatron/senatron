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

package handlers

import (
	"errors"
	"net/http"
)

// Err404 triggers a 404 response when thrown in a panic.
var Err404 = errors.New("Not found")

// Err500 triggers a 500 response when thrown in a panic.
var Err500 = errors.New("Internal error")

// FourOhFour writes out a standard page not found error message.
func FourOhFour(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Not Found"))
}

// FiveHundred writes out a standard internal server error message.
func FiveHundred(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Internal Server Error"))
}
