package web

import (
	"fmt"
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

	s := store.STORE
	if err := s.Set(fmt.Sprintf("%s", key), value); err != nil {
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
	s := store.STORE
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
	s := store.STORE
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
	s := store.STORE
	v := s.MGet(keys)
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   v,
	})
}

func del(c *fiber.Ctx) {
	key := c.Params("key")
	s := store.STORE
	if err := s.Del(key); err != nil {
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

func reset(c *fiber.Ctx) {
	s := store.STORE
	if err := s.Reset(); err != nil {
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

func push(c *fiber.Ctx) {
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

	s := store.STORE
	if err := s.Push(fmt.Sprintf("%s", key), value); err != nil {
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

func unshift(c *fiber.Ctx) {
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

	s := store.STORE
	if err := s.Unshift(fmt.Sprintf("%s", key), value); err != nil {
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

func pop(c *fiber.Ctx) {
	key := c.Params("key")

	s := store.STORE
	v, err := s.Pop(fmt.Sprintf("%s", key))
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

func shift(c *fiber.Ctx) {
	key := c.Params("key")

	query := struct {
		Shift int `query:"shift"`
	}{}

	if err := c.QueryParser(&query); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return
	}

	s := store.STORE
	if err := s.Shift(fmt.Sprintf("%s", key), query.Shift); err != nil {
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

func arrdel(c *fiber.Ctx) {
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

	s := store.STORE
	if err := s.ArrRm(fmt.Sprintf("%s", key), value); err != nil {
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

func prefix(c *fiber.Ctx) {
	key := c.Query("q")
	s := store.STORE
	data := s.Prefix(key)
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   data,
	})
}

func search(c *fiber.Ctx) {
	srch := c.Query("q")
	s := store.STORE
	data := s.KeySearch(srch)
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   data,
	})
}

func addReplicaServer(c *fiber.Ctx) {
	server := c.Query("address", "")
	if server == "" {
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Address not supplied",
		})
		return
	}
	s := store.STORE
	go s.AddReplicaServer(server)
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
}
