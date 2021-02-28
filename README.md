# NewRelic Integration Middleware
[![Go Reference](https://pkg.go.dev/badge/github.com/nobuyo/nrfiber.svg)](https://pkg.go.dev/github.com/nobuyo/nrfiber)

NewRelic Integration middleware for [Fiber](https://github.com/gofiber/fiber).

## Table of Contents

- [NewRelic Integration Middleware](#newrelic-integration-middleware)
	- [Table of Contents](#table-of-contents)
	- [Signatures](#signatures)
	- [Examples](#examples)
		- [Default Config](#default-config)
		- [Custom Config](#custom-config)
	- [Config](#config)
	- [Default Config](#default-config-1)

## Signatures

```go
func New(config ...Config) fiber.Handler
```

## Examples

First import the middleware,

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/nobuyo/nrfiber"
)
```

Then create a Fiber app with `app := fiber.New()`.

### Default Config

```go
nrapp, err := newrelic.NewApplication(
	newrelic.ConfigAppName("Application Name"),
	newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
	newrelic.ConfigDebugLogger(os.Stdout),
)

app.Use(nrfiber.New(nrfiber.Config{
	NewRelicApp: nrapp
}))
```

## Config

```go

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// NewRelicApp is newrelic.Application
	//
	// Required.
	NewRelicApp *newrelic.Application
}
```

## Default Config

```go
var ConfigDefault = Config{
	Next:        nil,
	NewRelicApp: &newrelic.Application{},
}
```
