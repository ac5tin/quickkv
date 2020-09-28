package web

import "github.com/gofiber/fiber/v2"

// Routes - define /api/web routes
func Routes(router *fiber.Router) {
	(*router).Post("/set/:key", set)
	(*router).Get("/get/:key", get)
	(*router).Get("/", getAll)
	(*router).Post("/mget", mget)
	(*router).Delete("/del/:key", del)
	(*router).Delete("/reset", reset)
	(*router).Post("/push/:key", push)
	(*router).Post("/unshift/:key", unshift)
	(*router).Post("/rm/:key", arrdel)
	(*router).Get("/pop/:key", pop)
	(*router).Get("/shift/:key", shift)
	(*router).Get("/prefix", prefix)
	(*router).Get("/search", search)
	(*router).Get("/replica/add", addReplicaServer)
	(*router).Get("/replica", listReplicas)
}
