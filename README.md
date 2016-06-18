# Pulse
Cutting down the boilerplate to create impactful Go web apps.

## Getting started
Run `make install`

*This will*
##### Add project dependencies to glide
##### Copy `example.env` and modify for your environment
  
  `cp example.env .env`

##### Fetch database migration tool

### Build
Simply running `go build` will build the project, and running the output binary
will bring up the server at the configured AppPort (By default this is 8080).

Visit http://localhost:8080 to view the running application.

### Create

### Dependency Injection
You can inject any shared application dependencies using https://github.com/facebookgo/inject.

### Controller/Routes
Uses Gin's router based on the ever popular httprouter

Remove the sample index page from the `routes.go` file and begin creating controllers
under the controllers package directory, and registering them as handlers for routes
in `routes.go`.

### Middleware
Uses Gin's middleware chaining

### DB Models
Uses Gorm for Database modelling and connections. Only supports sqlite, postgres and mysql. See http://jinzhu.me/gorm/
