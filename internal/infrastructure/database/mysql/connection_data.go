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

func (cd *ConnectionData) SetupBetaConnectionData(username, password, host, schema string) *ConnectionData {
	cd.Dialect = mySQL
	cd.Username = "root"
	cd.Password = "CGE-2Ebch6FGcdaG--CeecaCe2Bb2bCg"
	cd.Host = "viaduct.proxy.rlwy.net"
	cd.Schema = "railway"
	return cd
}

func (cd *ConnectionData) SetupLocalConnectionData(username, password, host, schema string) *ConnectionData {
	cd.Dialect = mySQL
	cd.Host = host
	cd.Schema = schema
	cd.Username = username
	cd.Password = password

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
