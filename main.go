package main

import (
	"flag"
	"fmt"
	"log"
	"quickkv/store"
	"quickkv/web"
	"syscall"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	port            = flag.Int("p", 8310, "Port to listen to")
	filepath        = flag.String("f", "qkv", "File path")
	enc             = flag.Bool("enc", false, "Enable encryption")
	pw       string = ""
)

func main() {
	flag.Parse()

	// encryption
	if *enc {
		fmt.Println("- Enabled Encryption mode -\nPlease Enter Password: (input is hidden)")
		bytepw, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Panic("Failed to read password")
		}
		pw = string(bytepw)
	}
	// web server
	app := fiber.New()
	// settings
	app.Settings.BodyLimit = 1024 * 1024 * 1024
	// middleware
	app.Use(middleware.Compress())
	// store
	store.Init(*filepath, pw)
	app.Use(func(c *fiber.Ctx) {
		c.Next()
	})
	// ==== API ROUTES =====
	app.Get("/ping", func(c *fiber.Ctx) { c.Status(200).Send("pong") })

	webapi := app.Group("/api/web")
	web.Routes(&webapi)
	// ===== ERROR RECOVER =====
	app.Use(middleware.Recover())
	// ==== LOGGER =====
	app.Use(middleware.Logger())
	// start server
	log.Println(fmt.Sprintf("Listening on PORT %d", *port))
	if err := app.Listen(*port); err != nil {
		log.Fatal(err.Error())
	}
}
