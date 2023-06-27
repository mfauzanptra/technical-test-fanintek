package auth

import "github.com/labstack/echo/v4"

type Core struct {
	Id            int    `json:"id"`
	Nama          string `json:"nama" validate:"required"`
	Email         string `json:"email" validate:"required"`
	NPP           string `json:"npp" validate:"required"`
	NPPSupervisor string `json:"npp_supervisor"`
	Password      string `json:"password" validate:"required"`
}

type AuthHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type AuthService interface {
	Register(newUser Core) error
	Login(email, password string) (string, Core, error)
}

type AuthData interface {
	Register(newUser Core) error
	Login(email string) (Core, error)
}
