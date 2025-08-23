package http

import (
	"io"
	"net/http"

	"github.com/DigitalCatServices-io/claimav-wrapper/src/configuration"
	"github.com/gofiber/fiber/v2"
)

func (h *ClamAVHandler) ScanFile(c *fiber.Ctx) error {

	// get file from body
	formFile, formFileErr := c.FormFile(configuration.REST.SCANFILE_FORM_KEY)
	if formFileErr != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	file, fileErr := formFile.Open()
	if fileErr != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	defer file.Close()

	fileContent, fileContentErr := io.ReadAll(file)
	if fileContentErr != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	// request service layer
	err := h.service.ScanFile(fileContent)
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	c.SendStatus(http.StatusOK)

	return nil
}
