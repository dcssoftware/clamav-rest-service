package api

import (
	"github.com/gofiber/fiber/v2"

	claimavHttp "github.com/DigitalCatServices-io/claimav-wrapper/src/resources/clamav/http"
	claimavService "github.com/DigitalCatServices-io/claimav-wrapper/src/resources/clamav/service"
)

type ApiModel struct {
	ClamAVHandler ClamAVHandler
}

func CreateAPI() (*ApiModel, error) {

	clamavSvc := claimavService.NewClamAVService()
	clamavHandler := claimavHttp.NewClamAVHandler(clamavSvc)

	return &ApiModel{
		ClamAVHandler: clamavHandler,
	}, nil
}

func (app *ApiModel) RunAPI() {
	api := fiber.New()

	api.All("/scan-document", app.ClamAVHandler.ScanFile)

	api.Listen(":33779")
}
