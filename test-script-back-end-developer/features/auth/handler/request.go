package handler

import "testAPI/features/auth"

type RegisterRequest struct {
	Nama          string `json:"nama"`
	Email         string `json:"email"`
	NPP           string `json:"npp"`
	NPPSupervisor string `json:"npp_supervisor"`
	Password      string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(request RegisterRequest) *auth.Core {
	return &auth.Core{
		Nama:          request.Nama,
		Email:         request.Email,
		NPP:           request.NPP,
		NPPSupervisor: request.NPPSupervisor,
		Password:      request.Password,
	}
}
