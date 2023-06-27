package handler

import (
	"log"
	"net/http"
	"strings"
	"testAPI/features/auth"
	"testAPI/helper"

	"github.com/labstack/echo/v4"
)

type authControl struct {
	srv auth.AuthService
}

func New(srv auth.AuthService) auth.AuthHandler {
	return &authControl{
		srv: srv,
	}
}

func (ac *authControl) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}

		if err := c.Bind(&input); err != nil {
			log.Println("error bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		err := ac.srv.Register(*ToCore(input))
		if err != nil {
			if err != nil {
				if strings.Contains(err.Error(), "validation") {
					return c.JSON(http.StatusBadRequest, helper.ErrorResponse("name, email, password, npp are required"))
				} else {
					return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal error"))
				}
			}
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success register",
		})
	}
}

func (ac *authControl) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}

		if err := c.Bind(&input); err != nil {
			log.Println("error bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		token, res, err := ac.srv.Login(input.Email, input.Password)
		if err != nil {
			if strings.Contains(err.Error(), "no rows") {
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("user not found"))
			} else if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("wrong password"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal error"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success login",
			"data":    CoreToLoginResponse(res),
			"token":   token,
		})
	}
}
