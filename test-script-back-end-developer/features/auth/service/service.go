package service

import (
	"errors"
	"log"
	"strings"
	"testAPI/features/auth"
	"testAPI/helper"

	"github.com/go-playground/validator/v10"
)

type AuthService struct {
	qry auth.AuthData
	vld *validator.Validate
}

func New(ad auth.AuthData) auth.AuthService {
	return &AuthService{
		qry: ad,
		vld: validator.New(),
	}
}

func (as *AuthService) Register(newUser auth.Core) error {
	err := as.vld.Struct(newUser)
	if err != nil {
		log.Println("error validation register: ", err)
		return err
	}

	hashed := helper.GeneratePassword(newUser.Password)
	newUser.Password = hashed

	err = as.qry.Register(newUser)
	if err != nil {
		return errors.New("server problem")
	}

	return nil
}

func (as *AuthService) Login(email, password string) (string, auth.Core, error) {
	res, err := as.qry.Login(email)
	if err != nil {
		errmsg := ""
		if strings.Contains(err.Error(), "no rows") {
			errmsg = err.Error()
		} else {
			errmsg = "server problem"
		}
		return "", auth.Core{}, errors.New(errmsg)
	}

	if err := helper.ComparePassword(res.Password, password); err != nil {
		return "", auth.Core{}, errors.New("wrong password")
	}

	token, _ := helper.GenerateJWT(int(res.Id), res.NPP)

	return token, res, nil
}
