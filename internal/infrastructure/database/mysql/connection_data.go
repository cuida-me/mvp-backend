package mysql

import (
	"fmt"

	msql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	cd.Password = ""

	return cd
}

func (cd *ConnectionData) SetupTestConnectionData() *ConnectionData {
	cd.Dialect = "sqlite3"
	cd.Host = "file::memory:?cache=shared"
	cd.Schema = "mvp-backend"

	return cd
}

func (cd *ConnectionData) toDialect() gorm.Dialector {
	//if cd.Dialect == "sqlite3" {
	//	host := cd.Host
	//	return slite.Open(host)
	//}

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=%s", cd.Username, cd.Password, cd.Host,
		cd.Schema, "UTC")

	return msql.Open(url)
}
