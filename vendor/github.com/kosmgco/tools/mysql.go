package tools

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	User     string `json:"user"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     string `json:"port"`
	db       *sql.DB
}

func (m *MySQL) Get() *sql.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.User, m.Password, m.Host, m.Port, m.Database)
	if m.db == nil {
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			logrus.Fatal(err)
		}
		m.db = db
	}
	return m.db
}

func (m *MySQL) Close() {
	if m != nil && m.db != nil {
		m.db = nil
	}
}
