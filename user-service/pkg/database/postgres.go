package database

import (
	"database/sql"
	"fmt"
)

func PostgreConnection(host string, port string, user string, password string, dbname string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
