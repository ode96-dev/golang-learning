package main

import "github.com/gofiber/fiber/v2"

/*
Fiber
- inspired by express.js(node.js). its fast, elegant, minimal memory usage
- super fast (based on fasthttp)
- express-style syntax
- built-in middleware (compression, logger, CORS, etc)
- websocket support
- easy error handling and routing
*/

func main() {
	app := fiber.New()

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello fiber")
	})

	app.Listen(":3000")

}
