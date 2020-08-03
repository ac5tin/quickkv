package web

import (
	"log"
	"quickkv/store"

	"github.com/gofiber/fiber"
)

func set(c *fiber.Ctx) {
	key := c.Params("key")
	var value interface{}
	if err := c.BodyParser(&value); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return
	}

	s := c.Locals("store").(*store.Store)
	if err := s.Set(key, value); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
}
