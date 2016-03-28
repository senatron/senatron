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
	"fmt"
	"github.com/bieber/conflag"
	"github.com/senatron/senatron/senatronserver/context"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Config defines configuration options for the server.
type Config struct {
	Help bool
	HTTP struct {
		Port                int
		StaticResourcesPath string
	}
	Log struct {
		FilePath string
	}
}

func main() {
	config, parser := getConfig()
	_, err := parser.Read()
	if err != nil || config.Help {
		exitCode := 0

		if err != nil {
			log.Println(err)
			exitCode = 1
		}

		if width, _, err := terminal.GetSize(0); err == nil {
			fmt.Println(parser.Usage(uint(width)))
		}
		os.Exit(exitCode)
	}

	var logOut io.Writer = os.Stderr
	if config.Log.FilePath != "" {
		fout, err := os.Create(config.Log.FilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer fout.Close()
		log.SetOutput(fout)
		logOut = fout
	}

	rand.Seed(time.Now().Unix())

	globalContext := &context.GlobalContext{
		LogOut: logOut,
	}

	initRoutes(globalContext, config.HTTP.StaticResourcesPath)

	err = initTemplates(globalContext, config.HTTP.StaticResourcesPath)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", globalContext.Router)

	log.Printf("Starting server on port %d...", config.HTTP.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.HTTP.Port), nil))
}

func getConfig() (*Config, *conflag.Config) {
	config := &Config{}
	config.HTTP.Port = 8080

	parser, err := conflag.New(config)
	if err != nil {
		log.Fatal(err)
	}

	parser.ProgramName("senatronserver")
	parser.ProgramDescription("HTTP server for senatron")
	parser.ConfigFileLongFlag("config")

	parser.Field("Help").
		ShortFlag('h').
		Description("Print usage text and exit.")

	parser.Field("HTTP.Port").
		ShortFlag('p').
		Description("Port to serve HTTP traffic on.")

	parser.Field("HTTP.StaticResourcesPath").
		ShortFlag('s').
		LongFlag("static-resources").
		Required().
		Description("Root directory to load static resources from.")

	parser.Field("Log.FilePath").
		ShortFlag('l').
		LongFlag("log-file").
		Description("Optional log output file (logs go to stderr by default)")

	return config, parser
}
