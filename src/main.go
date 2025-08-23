package main

import (
	"fmt"

	webapi "github.com/DigitalCatServices-io/claimav-wrapper/src/web-api"
)

func main() {
	api, err := webapi.CreateAPI()
	if err != nil {
		err = fmt.Errorf("%s", err.Error())
		panic(err)
	}

	api.RunAPI()

}
