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

package middleware

import (
	"github.com/senatron/senatron/senatronserver/context"
	"github.com/senatron/senatron/senatronserver/handlers"
	"net/http"
)

// ErrorCatcher retrieves any error-codes that a controller may panic
// with and delegates to the corresponding error controller.
func ErrorCatcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			localContext := context.Get(r)

			err := recover()
			if err == nil {
				return
			}

			switch err {
			case handlers.Err404:
				handlers.FourOhFour(w, r)
			default:
				localContext.Logger.Printf("PANIC: %v", err)
				handlers.FiveHundred(w, r)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
