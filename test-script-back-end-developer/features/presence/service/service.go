package service

import (
	"errors"
	"log"
	"testAPI/features/presence"
	"testAPI/helper"
)

type presenceService struct {
	qry presence.PresenceData
}

func New(pd presence.PresenceData) presence.PresenceService {
	return &presenceService{
		qry: pd,
	}
}

func (ps *presenceService) Insert(token interface{}, newData presence.Core) (presence.Core, error) {
	id, test := helper.ExtractToken(token)
	log.Println(test)
	newData.IdUser = id
	if id <= 0 {
		log.Println("error extraxt token")
		return presence.Core{}, errors.New("user id not valid")
	}

	res, err := ps.qry.Insert(newData)
	if err != nil {
		return presence.Core{}, err
	}

	return res, nil
}
func (ps *presenceService) Approve(token interface{}, idData int, approve bool) error {
	_, idSpv := helper.ExtractToken(token)
	if idSpv <= "" {
		log.Println("error extraxt token")
		return errors.New("user id not valid")
	}

	err := ps.qry.Approve(idSpv, idData, approve)
	if err != nil {
		return err
	}

	return nil
}
func (ps *presenceService) GetData() (presence.GetPresenceResp, error) {
	res, err := ps.qry.GetData()
	if err != nil {
		return presence.GetPresenceResp{}, err
	}

	return res, nil
}
func (ps *presenceService) GetDataByIdUser(idUser int) (presence.GetPresenceResp, error) {
	res, err := ps.qry.GetDataByIdUser(idUser)
	if err != nil {
		return presence.GetPresenceResp{}, err
	}

	return res, nil
}
