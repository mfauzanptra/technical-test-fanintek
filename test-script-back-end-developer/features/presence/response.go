package presence

type GetPresence struct {
	IdUser       int    `json:"id_user"`
	NamaUser     string `json:"nama_user"`
	Tanggal      string `json:"tanggal"`
	WaktuMasuk   string `json:"waktu_masuk"`
	WaktuPulang  string `json:"waktu_pulang"`
	StatusMasuk  string `json:"status_masuk"`
	StatusPulang string `json:"status_pulang"`
}

type GetPresenceResp struct {
	Message string        `json:"message"`
	Data    []GetPresence `json:"data"`
}
