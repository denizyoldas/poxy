package main

import "github.com/gofiber/fiber/v2"

// net/http --> gin, echo , gorilla/mux
// valyala/fasthttp --> fiber

// https://docs.gofiber.io/

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// https://docs.gofiber.io/api/middleware/proxy
	// Make request within handler
	app.Get("/:key/*", ProxyHandler)

	// TODO: implement handler!
	app.Delete("cache/:key/*", EvictCacheHandler)

	// TODO: implement handler!
	app.Delete("limit/:key/*", ResetLimitHandler)

	app.Listen(":1010")
}
