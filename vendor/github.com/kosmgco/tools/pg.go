package tools

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PosgreSQL struct {
	User     string `json:"user"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     string `json:"port"`
	db       *sql.DB
}

func (p *PosgreSQL) Get() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Database)
	if p.db == nil {
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			logrus.Fatal(err)
		}
		p.db = db
		return db
	}
	return p.db
}

func (p *PosgreSQL) Close() {
	if p != nil && p.db != nil {
		p.db.Close()
		p.db = nil
	}
}
