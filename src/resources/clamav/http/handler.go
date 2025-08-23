package http

import (
	clamavservice "github.com/DigitalCatServices-io/claimav-wrapper/src/resources/clamav/service"
)

type ClamAVHandler struct {
	service *clamavservice.ClamAVService
}

func NewClamAVHandler(clamavservice *clamavservice.ClamAVService) *ClamAVHandler {
	return &ClamAVHandler{
		service: clamavservice,
	}
}
