# Pulse
Cutting down the boilerplate to create impactful Go web apps.

## Getting started
Add project dependencies to glide

  glide init
  glide up

Open main.go and add the import paths for app/config/lib

  "github.com/$yourusername/$thisproject/app"
  "github.com/$yourusername/$thisproject/lib"
  "github.com/$yourusername/$thisproject/config"

This can be simply accomplished by just saving the main.go file and allowing the
goimports to run and grab those import paths automatically.

Copy `example.env` and modify for your environment

  cp example.env .env

### Build
Simply running `go build` will build the project, and running the output binary
will bring up the server at the configured AppPort (By default this is 8080).

Visit http://localhost:8080 to view the running application.

### Create

### Dependency Injection
You can inject any shared application dependencies using https://github.com/facebookgo/inject. See the default index controller in `routes.go` to
see how you can inject a dependency, or refer to the inject docs.

### Controller/Routes
Uses Bone as the default Mux. See http://go-zoo.github.io/bone/

Remove the sample index page from the `routes.go` file and begin creating controllers
under the controllers package directory, and registering them as handlers for routes
in `routes.go`.

### Middleware
Uses Negroni for application middleware chaining. See https://github.com/codegangsta/negroni

### DB Models
Uses Gorm for Database modelling and connections. Only supports sqlite, postgres and mysql. See http://jinzhu.me/gorm/
