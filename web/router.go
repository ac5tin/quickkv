package web

import "github.com/gofiber/fiber"

// Routes - define /api/web routes
func Routes(router *fiber.Router) {
	(*router).Post("/set/:key", set)
}
