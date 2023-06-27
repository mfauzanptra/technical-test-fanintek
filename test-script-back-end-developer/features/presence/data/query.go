package data

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"testAPI/features/presence"
)

type presenceQry struct {
	db *sql.DB
}

func New(db *sql.DB) presence.PresenceData {
	return &presenceQry{
		db: db,
	}
}

func (pq *presenceQry) Insert(newData presence.Core) (presence.Core, error) {
	query := "INSERT INTO epresence (id_users, type, waktu) VALUES ($1, $2, $3) RETURNING id"
	stmt, err := pq.db.Prepare(query)
	if err != nil {
		log.Println("error prepare insert query: ", err)
		return presence.Core{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(newData.IdUser, newData.Type, newData.Waktu).Scan(&newData.Id)
	if err != nil {
		log.Println("error exec insert query: ", err)
		return presence.Core{}, err
	}

	return newData, nil
}
func (pq *presenceQry) Approve(idSpv string, idData int, approve bool) error {
	idSpvData := ""
	checkQuery := "SELECT npp_supervisor FROM epresence e JOIN users u ON u.id = e.id_users WHERE e.id = $1"
	err := pq.db.QueryRow(checkQuery, idData).Scan(&idSpvData)
	if err != nil {
		log.Println("error check approve query: ", err)
		return err
	}

	if idSpv != idSpvData {
		return errors.New("authority fail")
	}

	query := "UPDATE epresence SET is_approve = $1 WHERE id = $2"
	_, err = pq.db.Exec(query, approve, idData)
	if err != nil {
		log.Println("error exec query update approval: ", err)
		return err
	}

	return nil
}
func (pq *presenceQry) GetData() (presence.GetPresenceResp, error) {
	listPresence := []presence.GetPresence{}

	query := "SELECT e.id_users, u.nama, DATE(e.waktu) tanggal, TO_CHAR(MIN(CASE WHEN e.type = 'IN' THEN e.waktu END), 'HH24:MI:SS'), TO_CHAR(MAX(CASE WHEN e.type = 'OUT' THEN e.waktu END), 'HH24:MI:SS'), CASE WHEN MIN(CASE WHEN e.type = 'IN' THEN e.is_approve::int END) = 1 THEN 'APPROVE' ELSE 'REJECT' END, CASE WHEN MAX(CASE WHEN e.type = 'OUT' THEN e.is_approve::int END) = 1 THEN 'APPROVE' ELSE 'REJECT' END FROM epresence e JOIN users u ON e.id_users = u.id GROUP BY e.id_users, u.nama, DATE(e.waktu) ORDER BY tanggal"
	rows, err := pq.db.Query(query)
	if err != nil {
		log.Println("error GetData query: ", err)
		return presence.GetPresenceResp{}, err
	}

	for rows.Next() {
		pres := presence.GetPresence{}
		rows.Scan(&pres.IdUser, &pres.NamaUser, &pres.Tanggal, &pres.WaktuMasuk, &pres.WaktuPulang, &pres.StatusMasuk, &pres.StatusPulang)
		tgl := strings.Split(pres.Tanggal, "T")
		pres.Tanggal = tgl[0]
		listPresence = append(listPresence, pres)
	}

	res := presence.GetPresenceResp{
		Message: "success get data",
		Data:    listPresence,
	}
	return res, nil
}
func (pq *presenceQry) GetDataByIdUser(idUser int) (presence.GetPresenceResp, error) {
	listPresence := []presence.GetPresence{}

	query := "SELECT e.id_users, u.nama, DATE(e.waktu) tanggal, TO_CHAR(MIN(CASE WHEN e.type = 'IN' THEN e.waktu END), 'HH24:MI:SS'), TO_CHAR(MAX(CASE WHEN e.type = 'OUT' THEN e.waktu END), 'HH24:MI:SS'), CASE WHEN MIN(CASE WHEN e.type = 'IN' THEN e.is_approve::int END) = 1 THEN 'APPROVE' ELSE 'REJECT' END, CASE WHEN MAX(CASE WHEN e.type = 'OUT' THEN e.is_approve::int END) = 1 THEN 'APPROVE' ELSE 'REJECT' END FROM epresence e JOIN users u ON e.id_users = u.id WHERE u.id = $1 GROUP BY e.id_users, u.nama, DATE(e.waktu) ORDER BY tanggal"
	rows, err := pq.db.Query(query, idUser)
	if err != nil {
		log.Println("error GetData query: ", err)
		return presence.GetPresenceResp{}, err
	}

	for rows.Next() {
		pres := presence.GetPresence{}
		rows.Scan(&pres.IdUser, &pres.NamaUser, &pres.Tanggal, &pres.WaktuMasuk, &pres.WaktuPulang, &pres.StatusMasuk, &pres.StatusPulang)
		tgl := strings.Split(pres.Tanggal, "T")
		pres.Tanggal = tgl[0]
		listPresence = append(listPresence, pres)
	}

	res := presence.GetPresenceResp{
		Message: "success get data",
		Data:    listPresence,
	}
	return res, nil
}
