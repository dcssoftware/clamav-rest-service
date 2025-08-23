package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/DigitalCatServices-io/claimav-wrapper/src/configuration"
	"github.com/gofiber/fiber/v2"
)

func (h *ClamAVHandler) ScanFile(c *fiber.Ctx) error {

	// get file from body
	formFile, formFileErr := c.FormFile(configuration.REST.SCANFILE_FORM_KEY)
	if formFileErr != nil {
		return c.Status(http.StatusBadRequest).SendString("")
	}

	file, fileErr := formFile.Open()
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return c.Status(http.StatusBadRequest).SendString("")
	}

	defer file.Close()

	fileContent, fileContentErr := io.ReadAll(file)
	if fileContentErr != nil {
		fmt.Println(fileContentErr.Error())
		return c.Status(http.StatusBadRequest).SendString("")
	}

	// request service layer
	err := h.service.ScanFile(fileContent)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusUnprocessableEntity).SendString("")
	}

	return c.Status(http.StatusOK).SendString("")
}
