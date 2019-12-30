This repository provides a way to share any minor handlers for [iris](https://github.com/kataras/iris) web framework. You can view the built'n supported handlers by pressing [here](https://github.com/kataras/iris/tree/v12/middleware).

<!-- [![Build status](https://api.travis-ci.org/iris-contrib/middleware.svg?branch=v12&style=flat-square)](https://travis-ci.org/iris-contrib/middleware) -->

## Installation

Install a middleware, take for example the [cors](cors) one.

```sh
$ go get github.com/iris-contrib/middleware/cors
```

**import as**

```go
import "github.com/iris-contrib/middleware/cors"

// [...]
```

> go build

Middleware is just a chain handlers which can be executed before or after the main handler, can transfer data between handlers and communicate with third-party libraries, they are just functions.

<!-- | [permissionbolt](permissionbolt) | Middleware for keeping track of users, login states and permissions. | [permissionbolt/_example/main.go]( permissionbolt/_example/main.go) | -->

| Middleware | Description | Example |
| -----------|--------|-------------|
| [jwt](jwt) | Middleware checks for a JWT on the `Authorization` header on incoming requests and decodes it. | [jwt/_example](jwt/_example) |
| [cors](cors) | HTTP Access Control. | [cors/_example](cors/_example) |
| [secure](secure) | Middleware that implements a few quick security wins. | [secure/_example](secure/_example/main.go) |
| [tollbooth](tollboothic) | Generic middleware to rate-limit HTTP requests. | [tollbooth/_examples/limit-handler](tollbooth/_examples/limit-handler) |
| [cloudwatch](cloudwatch) |  AWS cloudwatch metrics middleware. |[cloudwatch/_example](cloudwatch/_example) |
| [new relic](newrelic) | Official [New Relic Go Agent](https://github.com/newrelic/go-agent). | [newrelic/_example](newrelic/_example) |
| [prometheus](prometheus)| Easily create metrics endpoint for the [prometheus](http://prometheus.io) instrumentation tool | [prometheus/_example](prometheus/_example) |
| [casbin](casbin)| An authorization library that supports access control models like ACL, RBAC, ABAC | [casbin/_examples](casbin/_examples) |
| [raven](raven)| Sentry client in Go | [raven/_example](https://github.com/iris-contrib/middleware/blob/v12/raven/_example/main.go) |
| [csrf](csrf)| Cross-Site Request Forgery Protection | [csrf/_example](https://github.com/iris-contrib/middleware/blob/v12/csrf/_example/main.go) |
### How can I register middleware?

**To a single route**

```go
app := iris.New()
app.Get("/mypath", myMiddleware1, myMiddleware2, func(ctx iris.Context){}, func(ctx iris.Context){}, myMiddleware5,myMainHandlerLast)
```

**To a party of routes or subdomain**

```go

p := app.Party("/sellers", authMiddleware, logMiddleware)

```

OR

```go
p := app.Party("/customers")
p.Use(logMiddleware)
```

**To all routes**

```go
app.Use(func(ctx iris.Context){}, myMiddleware2)
```

**To global, all routes, parties and subdomains**

```go
app.UseGlobal(func(ctx iris.Context){}, myMiddleware2)
```

## Can I use standard net/http handler with iris?

**Yes** you can, just pass the Handler inside the `iris.FromStd` in order to be converted into iris.Handler and register it as you saw before.

### Convert handler which has the form of `http.Handler/HandlerFunc`

```go
package main

import (
    "github.com/kataras/iris/v12"
)

func main() {
    app := iris.New()

    sillyHTTPHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
            println(r.RequestURI)
    })

    sillyConvertedToIon := iris.FromStd(sillyHTTPHandler)
    // FromStd can take (http.ResponseWriter, *http.Request, next http.Handler) too!
    app.Use(sillyConvertedToIon)

    app.Run(iris.Addr(":8080"))
}

```

## Contributing

If you are interested in contributing to this project, please push a PR.

## People

[List of all contributors](https://github.com/iris-contrib/middleware/graphs/contributors)
