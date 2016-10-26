package main

import (
	"github.com/Berlin-opendb-hack/opendb-mock/app"
	"github.com/goadesign/goa"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// GetAddresses runs the GetAddresses action.
func (c *UserController) GetAddresses(ctx *app.GetAddressesUserContext) error {
	// UserController_GetAddresses: start_implement

	// Put your logic here

	// UserController_GetAddresses: end_implement
	res := app.OpendbHackAddressCollection{}
	return ctx.OK(res)
}

// GetUserInfo runs the GetUserInfo action.
func (c *UserController) GetUserInfo(ctx *app.GetUserInfoUserContext) error {
	// UserController_GetUserInfo: start_implement

	// Put your logic here

	// UserController_GetUserInfo: end_implement
	res := &app.OpendbHackUserinfo{}
	return ctx.OK(res)
}
