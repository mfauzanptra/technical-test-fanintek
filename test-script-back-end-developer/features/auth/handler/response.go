package handler

import "testAPI/features/auth"

type LoginResponse struct {
	Id            int    `json:"id"`
	Nama          string `json:"nama"`
	Email         string `json:"email"`
	NPP           string `json:"npp"`
	NPPSupervisor string `json:"npp_supervisor"`
}

func CoreToLoginResponse(core auth.Core) LoginResponse {
	return LoginResponse{
		Id:            core.Id,
		Nama:          core.Nama,
		Email:         core.Email,
		NPP:           core.NPP,
		NPPSupervisor: core.NPPSupervisor,
	}
}
