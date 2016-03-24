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
	"github.com/bieber/logger"
	"github.com/senatron/senatron/senatronserver/context"
	"net/http"
	"sync"
	"time"
)

var loggerMutex = sync.Mutex{}

// Logger wraps a handler with basic HTTP logging
func Logger(
	globalContext *context.GlobalContext,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			localContext := context.Get(r)
			localContext.Logger = logger.New()

			t0 := time.Now()
			localContext.Logger.WriteString("====\n")
			localContext.Logger.Printf(
				"[%s] %s %s",
				r.Method,
				r.RemoteAddr,
				r.URL.String(),
			)

			next.ServeHTTP(w, r)

			localContext.Logger.Printf("FINISHED IN %v", time.Now().Sub(t0))

			loggerMutex.Lock()
			_, err := localContext.Logger.WriteTo(globalContext.LogOut)
			loggerMutex.Unlock()

			if err != nil {
				panic(err)
			}
		})
	}
}
