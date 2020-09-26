package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"quickkv/grpcserver"
	"quickkv/store"
	"quickkv/web"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	port            = flag.Int("p", 8310, "Port to listen to")
	gport           = flag.Int("gp", 27444, "GRPC Server Port")
	filepath        = flag.String("f", "qkv", "File path")
	enc             = flag.Bool("enc", false, "Enable encryption")
	pw       string = ""
)

func main() {
	gob.Register([]interface{}{})          // register for []interface{}
	gob.Register(map[string]interface{}{}) // register for []interface{}

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
	// store
	store.Init(*filepath, pw)

	// grpc
	go grpcserver.StartServer(uint16(*gport))

	// web server
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024,
	})
	// middleware
	app.Use(compress.New())
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Next()
		return nil
	})
	// ==== API ROUTES =====
	app.Get("/ping", func(c *fiber.Ctx) error { c.Status(200).Send([]byte("pong")); return nil })

	webapi := app.Group("/api/web")
	web.Routes(&webapi)
	// ===== ERROR RECOVER =====
	app.Use(recover.New())
	// ==== LOGGER =====
	app.Use(logger.New())
	// start server
	log.Println(fmt.Sprintf("Listening on PORT %d", *port))
	if err := app.Listen(fmt.Sprintf(":%d", *port)); err != nil {
		log.Fatal(err.Error())
	}
}
