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

package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/sebest/xff"
	"github.com/senatron/senatron/senatronserver/context"
	"github.com/senatron/senatron/senatronserver/handlers"
	"github.com/senatron/senatron/senatronserver/middleware"
	"net/http"
	"path"
	"path/filepath"
)

func initRoutes(
	globalContext *context.GlobalContext,
	staticResourcesPath string,
) {
	r := mux.NewRouter().StrictSlash(true)
	globalContext.Router = r

	basicStack := alice.New(
		// This bottom instance of ErrorCatcher will catch any
		// failures in the logging or cleanup code, as a last resort.
		middleware.ErrorCatcher,
		xff.Handler,
		middleware.ContextCleaner,
		middleware.Logger(globalContext),
		middleware.ErrorCatcher,
	)

	r.NotFoundHandler = basicStack.ThenFunc(handlers.FourOhFour)

	r.Handle("/", basicStack.Then(handlers.Index(globalContext)))

	staticHandler := func(subpath string) http.Handler {
		return basicStack.Then(
			http.StripPrefix(
				path.Join("/static/", subpath),
				http.FileServer(
					http.Dir(
						filepath.Join(staticResourcesPath, subpath),
					),
				),
			),
		)
	}
	s := r.PathPrefix("/static").Subrouter()

	s.Handle("/js/{rest:.*}", staticHandler("/js"))
	s.Handle("/css/{rest:.*}", staticHandler("/css"))
}
