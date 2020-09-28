package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"quickkv/grpcserver"
	"quickkv/store"
	"quickkv/web"
	"syscall"

	uf "github.com/ac5tin/usefulgo"
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

	// grpc + replication start
	go grpcserver.StartServer(uint16(*gport))

	go func() {
		if os.Getenv("MASTER_SERVER_ENDPOINT") == "" || os.Getenv("MY_ADDR") == "" {
			return
		}
		log.Println("-- STARTING IN REPLICA MODE --")
		masterServer := uf.NewMisc().NewURLBuilder(os.Getenv("MASTER_SERVER_SCHEME"), os.Getenv("MASTER_SERVER_ENDPOINT"), os.Getenv("MASTER_SERVER_PORT")).Build()
		log.Printf("-- CONNECTING TO MASTER : %s --\n", masterServer)
		log.Printf("-- self.server.address = %s --\n", os.Getenv("MY_ADDR"))
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/web/replica/add", masterServer), nil)
		if err != nil {
			log.Fatal(err.Error())
		}
		q := req.URL.Query()
		q.Add("address", fmt.Sprintf("%s:%d", os.Getenv("MY_ADDR"), *gport))
		req.URL.RawQuery = q.Encode()
		req.Header.Add("content-type", "application/json")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer res.Body.Close()
	}()

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
