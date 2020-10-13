package web

import (
	"errors"
	"fmt"
	"log"
	"quickkv/store"

	"github.com/gofiber/fiber/v2"
)

func set(c *fiber.Ctx) error {
	key := c.Params("key")
	var value interface{}
	if err := c.BodyParser(&value); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}

	s := store.STORE
	if err := s.Set(fmt.Sprintf("%s", key), value); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}

func get(c *fiber.Ctx) error {
	key := c.Params("key")
	s := store.STORE
	v, err := s.Get(key)
	if err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   v,
	})
	return nil
}

func getAll(c *fiber.Ctx) error {
	s := store.STORE
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   s.GetAll(),
	})
	return nil
}

func mget(c *fiber.Ctx) error {
	var keys []string
	if err := c.BodyParser(&keys); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}
	s := store.STORE
	v := s.MGet(keys)
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   v,
	})
	return nil
}

func del(c *fiber.Ctx) error {
	key := c.Params("key")
	s := store.STORE
	if err := s.Del(key); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}

func reset(c *fiber.Ctx) error {
	s := store.STORE
	if err := s.Reset(); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}

func push(c *fiber.Ctx) error {
	key := c.Params("key")
	var value interface{}
	if err := c.BodyParser(&value); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}

	s := store.STORE
	if err := s.Push(fmt.Sprintf("%s", key), value); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}

func unshift(c *fiber.Ctx) error {
	key := c.Params("key")
	var value interface{}
	if err := c.BodyParser(&value); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}

	s := store.STORE
	if err := s.Unshift(fmt.Sprintf("%s", key), value); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}

func pop(c *fiber.Ctx) error {
	key := c.Params("key")

	s := store.STORE
	v, err := s.Pop(fmt.Sprintf("%s", key))
	if err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   v,
	})
	return nil
}

func shift(c *fiber.Ctx) error {
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
		return err
	}

	s := store.STORE
	if err := s.Shift(fmt.Sprintf("%s", key), query.Shift); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}

func arrdel(c *fiber.Ctx) error {
	key := c.Params("key")
	var value interface{}
	if err := c.BodyParser(&value); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}

	s := store.STORE
	if err := s.ArrRm(fmt.Sprintf("%s", key), value); err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  err.Error(),
		})
		return err
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}

func prefix(c *fiber.Ctx) error {
	key := c.Query("q")
	s := store.STORE
	data := s.Prefix(key)
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   data,
	})
	return nil
}

func search(c *fiber.Ctx) error {
	srch := c.Query("q")
	s := store.STORE
	data := s.KeySearch(srch)
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   data,
	})
	return nil
}

func addReplicaServer(c *fiber.Ctx) error {
	server := c.Query("address", "")
	if server == "" {
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Address not supplied",
		})
		return errors.New("Address not supplied")
	}
	s := store.STORE
	go s.AddReplicaServer(server)
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}

func listReplicas(c *fiber.Ctx) error {
	s := store.STORE
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   s.ListReplicas(),
	})
	return nil
}

func resetReplicasList(c *fiber.Ctx) error {
	s := store.STORE
	s.ResetReplicasList()
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}

func increment(c *fiber.Ctx) error {
	key := c.Params("key")
	s := store.STORE
	if err := s.Increment(key); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// all done
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}

func decrement(c *fiber.Ctx) error {
	key := c.Params("key")
	s := store.STORE
	if err := s.Decrement(key); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// all done
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "success",
	})
	return nil
}
