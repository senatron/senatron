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

## Making It All Work (for OSX users)

#### Step 1: Install Go on your system (skip if you already have Go set up)

- Go [here](https://golang.org/dl/) and download the golang installer package for Apple OS X and follow the prompts to install Go on your system.

- Create a directory in your home directory (usually home is **` ~ `**, unless you've done something different with your system) and name it whatever you want. I called mine **`gocode/`**.

- Open up your **`.bash_profile`** or **`.bashrc`** or whatever your shell configuration file is, and set the path for Go. I use the bash shell and a **`.bash_profile`** file, and setting the path looks like this:
```js
export GOPATH=$HOME/gocode
```
*You add that line in your configuration file. If you use something different for your shell sessions, find out how to "set a path" for your specific set up (usually it starts with `export...`). Make sure the name of the directory you created in the second bullet point is in the path, in case you named it something else besides `gocode/`. After this is done, make sure to either start a new shell session (new terminal/CLI window or tab), or close and reopen the currently open one.*

#### Step 2: Back end project set up

- From **`gocode/`** (or your previously set up Go directory), run **`go get github.com/senatron/senatron/senatronserver`**. This command actually pulls the repo and put things in place for the back end server to be able to run.

- Now that you have the repo on your local machine, **`cd`** into it. It should be located inside the gocode directory you made, and should look something like this: **`~/gocode/src/github.com/senatron/senatron`**. The second "senatron" directory is what you will want to open up in your favorite text editor. This directory contains the project.

#### Step 3: Front end project set up

- Change into the **`static/`** directory. This is where you run **`npm install`**. When npm is finished, run **`gulp build`**. *This assumes you have `node/npm` and `gulp` already installed, go do that first if you don't have it.*

- **`cd`** back up one level to the project root, and create a new directory called **`config/`**.

- Inside **`config/`**, create a new file called **`test.conf`**. You will want to add something like the following to this file just to make things a little easier when running the app:
```html
[http]
port = 8080

static_resources_path = /Users/<your-user>/<the-go-directory-you-created>/src/github.com/senatron/senatron/static/build/

[sunlight]
api_key=<YOUR API KEY>
```
#### Step 4: Running the project

For the last few steps, you'll want to have two terminal/CLI windows or tabs open.

- **`cd`** one of the tabs to the **`static/`** directory and run just **`gulp`**. Then **`cd`** with your other tab all the way back up to the Go directory you created earlier in Step 1. Mine is called **`gocode`**. Then **`cd`** into **`bin/`**. Finally, from **`bin`** run this command to start the Go server:
```js
./senatronserver --config ../src/github.com/senatron/senatron/config/test.conf
```
*You may get a dialog that asks you if you want to accept or deny incoming connections, just click accept, and you should be able to see the app running on port 8080 in your browser :)*

- You're all set up now! One tab will run the Go server, which you can kill and restart by doing **`^C`** and rerunning the **`./senatronserver...`** command above. And then the other tab is where all the front end and **`gulp`** stuff happens!
