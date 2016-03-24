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
	"github.com/senatron/senatron/senatronserver/context"
	"net/http"
)

// Index renders the homepage.
func Index(globalContext *context.GlobalContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := globalContext.Templates.Index.Execute(
			w,
			map[string]interface{}{},
		)
		if err != nil {
			panic(err)
		}
	}
}
