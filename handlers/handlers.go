package handlers

import (
	"boilerplate/common"
	"boilerplate/database"
	"boilerplate/models"
	"boilerplate/service"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// UserGet returns a user
func UserList(c *fiber.Ctx) error {
	users := database.Get()
	return c.JSON(fiber.Map{
		"success": true,
		"user":    users,
	})
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) error {
	user := &models.User{
		Name: c.FormValue("user"),
	}
	database.Insert(user)
	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func PanicCheck(c *fiber.Ctx) error {
	common.CheckError(errors.New("ok"))
	return nil
}

func PushFile(s service.MinioService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":  err,
				"status": "failure",
			})
		}

		data, err := file.Open()
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":  err,
				"status": "failure",
			})
		}

		if !s.BucketExists("nttransfer") {
			err = s.CreateBucket("nttransfer")
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"error":  err,
					"status": "failure",
				})
			}
		}
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, data); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":  err,
				"status": "failure",
			})
		}

		path := fmt.Sprintf("%d_%s", time.Now().UnixNano(), strings.Replace(file.Filename, "/", "-", -1))
		err = s.PutFile("nttransfer", path, buf.Bytes())
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":  err,
				"status": "failure",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"path":   "http://nttransfer-api.herokuapp.com/download/" + path,
			"status": "success",
		})
	}
}

func GetFile(s service.MinioService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		path := c.Params("path")
		data, err := s.GetFile("nttransfer", path)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":  err,
				"status": "failure",
			})
		}
		err = s.DeleteFile("nttransfer", path)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":  err,
				"status": "failure",
			})
		}
		return c.Status(200).Send(data)
	}
}
