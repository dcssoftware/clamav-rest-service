package api

import "github.com/gofiber/fiber/v2"

type ClamAVHandler interface {
	ScanFile(c *fiber.Ctx) error
}
