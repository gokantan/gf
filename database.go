package gf

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"strconv"
)

type sqlDBFactory struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Pwd      string
	Database string

	IsEnable bool
}

func (this *sqlDBFactory) Check() error {
	this.IsEnable = (this.Driver != "")

	db, err := sql.Open(this.Driver,
		this.User + ":" + this.Pwd + "@tcp(" + this.Host + ":" + strconv.Itoa(this.Port) + ")/" + this.Database)
	if db != nil {
		db.Close()
	}

	return err
}

func (this *sqlDBFactory) NewConnect() (*sql.DB) {
	db, err := sql.Open(this.Driver,
		this.User + ":" + this.Pwd + "@tcp(" + this.Host + ":" + strconv.Itoa(this.Port) + ")/" + this.Database)
	if err != nil {
		return nil
	}
	return db
}