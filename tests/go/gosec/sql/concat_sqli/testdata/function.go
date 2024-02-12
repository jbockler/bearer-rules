package sqli

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/foo/utils/database"
)

var DB *sql.DB
var err error

type Profile struct {
	Uid         int
	Name        string
	City        string
	PhoneNumber string
}

func NewProfile() *Profile {
	return &Profile{}
}

func (p *Profile) UnsafeQueryGetData(uid string) error {

	/* this funciton use to get data Profile from database with vulnerable query */
	DB, err = database.Connect()

	getProfileSql := fmt.Sprintf(`SELECT p.user_id, p.full_name, p.city, p.phone_number
								FROM Profile as p,Users as u
								where p.user_id = u.id
								and u.id=%s`, uid) //here is the vulnerable query

	// bearer:expected go_gosec_sql_concat_sqli
	rows, err := DB.Query(getProfileSql)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&p.Uid, &p.Name, &p.City, &p.PhoneNumber)
		if err != nil {
			log.Printf("Row scan error: %s", err.Error())
			return err
		}
	}
	return nil
}

func (p *Profile) SafeQueryGetData(uid string) error {
	DB, err = database.Connect()

	const (
		getProfileSql = `SELECT p.user_id, p.full_name, p.city, p.phone_number
		FROM Profile as p,Users as u
		where p.user_id = u.id
		and u.id=?`
	)

	stmt, err := DB.Prepare(getProfileSql) //prepare statement
	if err != nil {
		return err
	}

	defer stmt.Close()
	err = stmt.QueryRow(uid).Scan(&p.Uid, &p.Name, &p.City, &p.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}
