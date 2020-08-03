package main

import (
	"flag"
	"fmt"
	"log"
	"quickkv/store"
	"quickkv/web"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

var (
	port     = flag.Int("p", 8310, "Port to listen to")
	filepath = flag.String("f", "qkv", "File path")
)

func main() {
	flag.Parse()
	app := fiber.New()
	// middleware
	app.Use(middleware.Compress())
	// store
	s := store.Init(*filepath)
	app.Use(func(c *fiber.Ctx) {
		c.Locals("store", s)
		c.Next()
	})
	// ==== API ROUTES =====
	app.Get("/ping", func(c *fiber.Ctx) { c.Status(200).Send("pong") })

	webapi := app.Group("/api/web")
	web.Routes(&webapi)
	// ===== ERROR RECOVER =====
	app.Use(middleware.Recover())
	// start server
	log.Println(fmt.Sprintf("Listening on PORT %d", *port))
	if err := app.Listen(*port); err != nil {
		log.Fatal(err.Error())
	}
}
