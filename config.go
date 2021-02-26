package nrfiber

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

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

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:       nil,
	AppName:    "Fiber App",
	LicenseKey: os.Getenv("NEW_RELIC_LICENSE_KEY"),
}
