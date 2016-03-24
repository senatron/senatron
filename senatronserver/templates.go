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
	"github.com/senatron/senatron/senatronserver/context"
	"html/template"
	"path/filepath"
)

func initTemplates(
	globalContext *context.GlobalContext,
	staticResourcesPath string,
) error {
	var err error
	staticPath := func(path string) string {
		return filepath.Join(staticResourcesPath, "template", path)
	}

	globalContext.Templates.Index, err = template.ParseFiles(
		staticPath("index.got"),
	)
	if err != nil {
		return err
	}

	return nil
}
