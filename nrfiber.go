package nrfiber

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func transformRequestHeaders(r *fasthttp.Request) http.Header {
	header := make(http.Header)
	r.Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		header.Set(sk, sv)
	})

	return header
}

func transformResponseHeaders(r *fasthttp.Response) http.Header {
	header := make(http.Header)
	r.Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		header.Set(sk, sv)
	})

	return header
}

func toHTTPRequest(ctx *fasthttp.RequestCtx) *http.Request {
	uri := ctx.Request.URI()
	url := &url.URL{
		Scheme:   string(uri.Scheme()),
		Path:     string(uri.Path()),
		Host:     string(uri.Host()),
		RawQuery: string(uri.QueryString()),
	}

	return &http.Request{
		Method: string(ctx.Request.Header.Method()),
		URL:    url,
		Proto:  "HTTP/1.1",
		Header: transformRequestHeaders(&ctx.Request),
		Host:   string(uri.Host()),
		TLS:    ctx.TLSConnectionState(),
	}
}

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := ConfigDefault
	// Override config if provided
	if len(config) > 0 {
		cfg = config[0]
	}

	app := cfg.NewRelicApp

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		txn := app.StartTransaction(c.Path())
		txn.SetWebRequestHTTP(toHTTPRequest(c.Context()))

		c.Locals("newrelic_transaction", txn)

		defer func() {
			rw := txn.SetWebResponse(&ResponseWriter{
				header: transformResponseHeaders(&c.Context().Response),
			})
			rw.Write(c.Context().Response.Body())
			rw.WriteHeader(c.Context().Response.StatusCode())
			txn.End()
		}()
		return c.Next()
	}
}
