package controllers

import (
	"github.com/goadesign/goa"
	"github.com/negronjl/csapi/app"
	"fmt"
	"encoding/json"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
	zapi_list ZAPIStructure
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service, zapi_list ZAPIStructure) *HealthController {
	return &HealthController{
		Controller: service.NewController("HealthController"),
		zapi_list:  zapi_list,
	}
}

// Health runs the health action.
func (c *HealthController) Health(ctx *app.HealthHealthContext) error {
	// HealthController_Health: start_implement

	fmt.Printf("DC: [%s]\n", ctx.Dc)
	fmt.Printf("Host Group: [%s]\n", ctx.Hgroup)
	fmt.Printf("Host Name: [%s]\n", ctx.Hostname)

	for index, element := range c.zapi_list {
		fmt.Printf("zapi_alias: [%s]\n", index)
		fmt.Printf("zapi_url: [%s]\n", element.zapi_url)
		fmt.Printf("zapi_username: [%s]\n", element.zapi_username)
		fmt.Printf("zapi_password: [%s]\n", element.zapi_password)
		version, err := element.zapi_object.Version()
		if err != nil {
			goa.ErrBadRequest("Unable to get version information: %s", err)
		}
		fmt.Printf("zapi_version: [%s]\n\n", version)
	}

	result, err := GetDCStatus(c.zapi_list[ctx.Dc])
	if err != nil {
		return fmt.Errorf("Erro communicating with ZAPI: %v\n", err)
	}
	retval, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("Unable to UnMarshall JSON object\n", err)
	}

	// HealthController_Health: end_implement
	return ctx.OK(retval)
}
