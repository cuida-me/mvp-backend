package database

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // need to load mysql driver on api
	"github.com/jinzhu/gorm"
)

const (
	minutesConnMaxLifetime = 2
	maxIdleConnections     = 50
	maxOpenConnections     = 100
)

func GetConnection(data *ConnectionData) (*gorm.DB, error) {
	if data == nil {
		return nil, fmt.Errorf("connection data is nil")
	}

	client, err := gorm.Open(data.Dialect, data.toString())
	if err != nil {
		return nil, err
	}

	client.DB().SetConnMaxLifetime(time.Minute * minutesConnMaxLifetime)
	client.DB().SetMaxIdleConns(maxIdleConnections)
	client.DB().SetMaxOpenConns(maxOpenConnections)

	return client, nil
}
