package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ortizdavid/go-null-dbvalues/entities"
)

type CustomerModel struct {
}

func (cm *CustomerModel) Create(c entities.Customer) error {
	db, err := sql.Open("mysql", "root:PASSWORD@tcp(127.0.0.1:3306)/go_null_dbvalues")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO customer (name, age, gender, height) VALUES (?, ?, ?, ?)", c.Name, c.Age, c.Gender, c.Height)
	if err != nil {
		return err
	}
	return nil
}