# senatron

`senatron` is going to be a web app to tell you the difference between
senate votes and the popular votes they represent.  It's also going to
have a much cooler public facing name as soon as we pick one.  This is
a brief guide to what's where in the repo.

## Source Code

### senatronserver/

This is the Go package that compiles to the server executable.  It
contains three files, `main.go`, `routes.go`, and `templates.go`.
`main.go` is the entry point for the executable, and it's responsible
for setting up global state and kicking off the web server.  This is
basically what happens on startup:

1. First we read configuration.  This comes from a combination of
   default values, config file (optional), and command-line flags.
   The function `getConfig` at the bottom of the file sets up all the
   configuration options.

2. If the user specified a log output file, open it for writing and
   set it as the file-handle to log data to (the default is STDERR).

3. Initialize routes.  This is handled by the `initRoutes` function
   defined in `routes.go`, which uses the `gorilla/mux` package to
   route URIs to handlers and set up the static resources directory.

4. Initialize templates.  This is handled by the `initTemplates`
   function defined in `templates.go`, which compiles all the template
   files for HTML output and stores their compiled versions in the
   global context.

5. Tell Go's http library to connect the `gorilla/mux` router to the
   root URI, then start up the HTTP server.  That's it.

#### senatronserver/handlers

This sub-package contains all the handler functions, which are
responsible for fielding HTTP requests and writing the output (which
will usually boil down to fetching some data and then rendering one of
the templates defined earlier).  In MVC terms, these are the
"controllers."  For more on handler functions, see [the `net/http`
docs](https://golang.org/pkg/net/http/);

#### senatronserver/middleware

This sub-package contains middleware functions.  These are functions
that wrap an HTTP handler and return a new one while doing something
extra.  We wrap multiple of them together to provide standard
functionality to all controllers.  For instance, the `Logger`
middleware wraps the handler by executing it and then logging the
request and the amount of time it took.  The `ErrorCatcher` middleware
makes sure that if a handler panics, we at least return some kind of a
response to the browser, and so on.  These are chained together using
the `alice` package (which simply composes middleware functions for
you into an easy-to-reuse chain) in the `initRoutes` function in the
`main` package.

#### senatronserver/context

There are two types of context in the server, and they're both defined
in this context.  Global context is potentially relevant to all
handlers, and there's only one instance of it for the entire app.  It
stores things like the compiled templates, the app configuration, the
`gorilla/mux` router information, and so on.

Local context stores data that's useful when handling an individual
request.  It's maintained in a global map (because the standard
`net/http` library unfortunately makes no allowance for state that can
be shared between middleware and handlers), and you can get your hands
on it when necessary by calling `context.Get(r)` from your middleware
or handler function, where `r` is the pointer to the `http.Request`
instance for the current request.  This is useful for things like
doing user authentication in middleware, or keeping a Logger instance
that bundles together all logged data for the current request.

### static/

This directory contains, unsurprisingly, static resources.  Currently
that means js, css, and template files.  Javascript gets compiled with
browserify and babel: every `source.js` file in the top-level of the
`static/js/` directory gets compiled with all its dependencies into a
corresponding `source.js` file in `static/build/js/`.  Meanwhile css
and template files just get copied verbatim to `static/build/css/` and
`static/build/template/`, so we end up with all our static resource
files piled together in a single build directory.

## Making It All Work

You'll need to have Go and npm both set up and working on your system.
First pull the repo by git cloning `github.com/senatron/senatron` into
your Go source directory.  You can download and install the server and
its dependencies by running

```
go get github.com/senatron/senatron/senatronserver
```

Then `cd` into the `static/` directory and build the static resources
by running

```
npm install
gulp build
```

To monitor the directory and rebuild static files when changed, just
run `gulp` without any argument.  Once you've got everything built,
you can start the server up: it will just need to know what port you
want it to run on and where you've got the built static resources.
For convenience, I just keep a test configuration file that looks like
this:

```
[http]
port = 8080
static_resources_path = /home/bieber/senatron/static/build/

[sunlight]
api_key=<YOUR API KEY>
```

Then, from the repo's root directory, you can fire up the server by
running

```
senatron --config config/test.conf
```

Substituting in the path to your config file, of course.  At that
point you can just hit `localhost:8080` (or whatever port you want to
use) in your browser to test.  Assuming you have your Go binary
directory added to your path, which I highly recommend.

*For more OSX-specific project set up instructions, see here: https://gist.github.com/heyheyjose/9db9e4bb438cfc98883506005355f1d8*
