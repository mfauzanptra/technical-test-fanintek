package handler

import "testAPI/features/presence"

type InsertReq struct {
	Type  string `json:"type"`
	Waktu string `json:"waktu"`
}

type ApproveReq struct {
	Approve bool `json:"approve"`
}

func ToCore(req InsertReq) presence.Core {
	return presence.Core{
		Type:  req.Type,
		Waktu: req.Waktu,
	}
}
