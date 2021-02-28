package nrfiber

import (
	"github.com/gofiber/fiber/v2"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// NewRelicApp is newrelic.Application
	//
	// Required. Default: empty Application
	NewRelicApp *newrelic.Application
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:        nil,
	NewRelicApp: &newrelic.Application{},
}
