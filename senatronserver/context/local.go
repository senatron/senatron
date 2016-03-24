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

package context

import (
	"github.com/bieber/logger"
	"net/http"
	"sync"
)

// LocalContext stores context relevant to a single request.  It
// should be both written to and read from by middleware, and read
// from by controllers.
type LocalContext struct {
	Logger     *logger.Logger
}

var localMutex = sync.Mutex{}
var localContexts = make(map[*http.Request]*LocalContext)

// Get returns the LocalContext for a given request, or creates one if
// it doesn't already exist.
func Get(request *http.Request) *LocalContext {
	localMutex.Lock()
	defer localMutex.Unlock()

	if c, ok := localContexts[request]; ok {
		return c
	}
	c := &LocalContext{}
	localContexts[request] = c
	return c
}

// Clear removes the LocalContext entry for a request after it's
// finished.
func Clear(request *http.Request) {
	localMutex.Lock()
	defer localMutex.Unlock()
	delete(localContexts, request)
}
