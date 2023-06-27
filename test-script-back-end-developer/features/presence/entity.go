package presence

import "github.com/labstack/echo/v4"

type Core struct {
	Id        int    `json:"id"`
	IdUser    int    `json:"id_user"`
	Type      string `json:"type"`
	IsApprove string `json:"is_approve"`
	Waktu     string `json:"waktu"`
}

type PresenceHandler interface {
	Insert() echo.HandlerFunc
	Approve() echo.HandlerFunc
	GetData() echo.HandlerFunc
	GetDataByIdUser() echo.HandlerFunc
}

type PresenceService interface {
	Insert(token interface{}, newData Core) (Core, error)
	Approve(token interface{}, idData int, approve bool) error
	GetData() (GetPresenceResp, error)
	GetDataByIdUser(idUser int) (GetPresenceResp, error)
}

type PresenceData interface {
	Insert(newData Core) (Core, error)
	Approve(idSpv string, idData int, approve bool) error
	GetData() (GetPresenceResp, error)
	GetDataByIdUser(idUser int) (GetPresenceResp, error)
}
