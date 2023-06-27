package data

import (
	"database/sql"
	"log"
	"testAPI/features/auth"
)

type authQuery struct {
	db *sql.DB
}

func New(db *sql.DB) auth.AuthData {
	return &authQuery{
		db: db,
	}
}

func (aq *authQuery) Register(newUser auth.Core) error {
	query := "INSERT INTO Users(nama, email, npp, npp_supervisor, password) VALUES ($1,$2,$3,$4,$5)"
	stmt, err := aq.db.Prepare(query)
	if err != nil {
		log.Println("error prepare query register: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newUser.Nama, newUser.Email, newUser.NPP, newUser.NPPSupervisor, newUser.Password)
	if err != nil {
		log.Println("error execute query register: ", err)
		return err
	}
	return nil
}

func (aq *authQuery) Login(email string) (auth.Core, error) {
	res := auth.Core{}

	query := "SELECT * FROM Users WHERE email = $1"
	stmt, err := aq.db.Prepare(query)
	if err != nil {
		log.Println("error prepare login query: ", err)
		return auth.Core{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(email).Scan(&res.Id, &res.Nama, &res.Email, &res.NPP, &res.NPPSupervisor, &res.Password)
	if err != nil {
		log.Println("error exec login query: ", err)
		return auth.Core{}, err
	}

	return res, nil
}
