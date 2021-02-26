# NewRelic Integration Middleware

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

First import the middleware from Fiber,

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/nobuyo/nrfiber"
)
```

Then create a Fiber app with `app := fiber.New()`.

### Default Config

```go
app.Use(nrfiber.New())
```

### Custom Config

```go
app.Use(nrfiber.New(nrfiber.Config{
  AppName:    "Application Name",
  LicenseKey: "CUSTOM_LICENSE_KEY_HERE",
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

	// AppName defines name of this fiber app in NewRelic
	//
	// Optional. Default value "Fiber App".
	AppName string

	// AppName defines name of this fiber app in NewRelic
	//
	// Optional. Default value `os.Getenv("NEW_RELIC_LICENSE_KEY")`.
	LicenseKey string
}
```

## Default Config

```go
var ConfigDefault = Config{
	Next:       nil,
	AppName:    "Fiber App",
	LicenseKey: os.Getenv("NEW_RELIC_LICENSE_KEY"),
}
```
