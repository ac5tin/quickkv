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

func get(c *fiber.Ctx) {
	key := c.Params("key")
	s := c.Locals("store").(*store.Store)
	v, err := s.Get(key)
	if err != nil {
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
		"data":   v,
	})
}

func getAll(c *fiber.Ctx) {
	s := c.Locals("store").(*store.Store)
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   s.GetAll(),
	})
}

func mget(c *fiber.Ctx) {
	var keys []string
	if err := c.BodyParser(&keys); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return
	}
	s := c.Locals("store").(*store.Store)
	v := s.MGet(keys)
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   v,
	})
}
