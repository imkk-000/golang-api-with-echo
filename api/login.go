package api

import (
	"go-echo-api/model"
	"go-echo-api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// LoginAPI is struct api
type LoginAPI struct {
	Service service.LoginService
}

// NewLoginAPI is wrap struct loginAPI
func NewLoginAPI(service service.LoginService) *LoginAPI {
	return &LoginAPI{
		Service: service,
	}
}

func (a LoginAPI) Login(c echo.Context) error {
	var loginRequest model.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, model.LoginResponse{
			Status:       "fail",
			ErrorMessage: "validate error",
		})
	}
	if err := c.Validate(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, model.LoginResponse{
			Status:       "fail",
			ErrorMessage: "validate error",
		})
	}
	return c.JSON(http.StatusOK, model.LoginResponse{
		Status: "success",
	})
}
