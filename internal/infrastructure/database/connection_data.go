package database

import "fmt"

const mySQL = "mysql"

type ConnectionData struct {
	Host     string
	Schema   string
	Username string
	Password string
	Dialect  string
}

func (cd *ConnectionData) SetupProdConnectionData() *ConnectionData {
	cd.Dialect = mySQL
	cd.Username = ""
	cd.Password = ""
	cd.Host = ""
	cd.Schema = ""

	return cd
}

func (cd *ConnectionData) SetupBetaConnectionData() *ConnectionData {
	cd.Dialect = mySQL
	cd.Username = ""
	cd.Password = ""
	cd.Host = ""
	cd.Schema = ""
	return cd
}

func (cd *ConnectionData) SetupLocalConnectionData() *ConnectionData {
	cd.Dialect = mySQL
	cd.Host = "localhost:3306"
	cd.Schema = "cuidamelocal"
	cd.Username = "root"
	cd.Password = "1234"

	return cd
}

func (cd *ConnectionData) SetupTestConnectionData() *ConnectionData {
	cd.Dialect = "sqlite3"
	cd.Host = "file::memory:?cache=shared"
	cd.Schema = "mvp-backend"

	return cd
}

func (cd *ConnectionData) toString() string {
	if cd.Dialect == "sqlite3" {
		return cd.Host
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=%s", cd.Username, cd.Password, cd.Host,
		cd.Schema, "UTC")
}
