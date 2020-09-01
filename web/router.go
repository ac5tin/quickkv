package web

import "github.com/gofiber/fiber"

// Routes - define /api/web routes
func Routes(router *fiber.Router) {
	(*router).Post("/set/:key", set)
	(*router).Get("/get/:key", get)
	(*router).Get("/", getAll)
	(*router).Post("/mget", mget)
	(*router).Delete("/del/:key", del)
	(*router).Delete("/reset", reset)
	(*router).Post("/push/:key", push)
}
