package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"testAPI/features/presence"
	"testAPI/helper"

	"github.com/labstack/echo/v4"
)

type presenceHandler struct {
	srv presence.PresenceService
}

func New(ps presence.PresenceService) presence.PresenceHandler {
	return &presenceHandler{
		srv: ps,
	}
}

func (ph *presenceHandler) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		input := InsertReq{}

		if err := c.Bind(&input); err != nil {
			log.Println("error insert request: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		res, err := ph.srv.Insert(token, ToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "id") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid user id"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success insert data",
			"data":    res,
		})
	}
}
func (ph *presenceHandler) Approve() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		input := ApproveReq{}
		if err := c.Bind(&input); err != nil {
			log.Println("error insert request: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Println("Read param error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong product id parameter"))
		}

		err = ph.srv.Approve(token, id, input.Approve)
		if err != nil {
			if strings.Contains(err.Error(), "authority") {
				return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized account"))
			} else if strings.Contains(err.Error(), "valid") {
				return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("not valid id spv"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
		})
	}
}
func (ph *presenceHandler) GetData() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, _ := ph.srv.GetData()
		return c.JSON(http.StatusOK, res)
	}
}
func (ph *presenceHandler) GetDataByIdUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		idString := c.Param("id")
		idUser, _ := strconv.Atoi(idString)
		res, _ := ph.srv.GetDataByIdUser(idUser)
		return c.JSON(http.StatusOK, res)
	}
}
