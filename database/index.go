package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

type Index struct {
	Name string
	Platform []uint8
	Language []uint8
	Targets []uint8
}

type SearchByParams struct {
	Name string
	Platform string
	Language string
}

func (i *Index) SearchBy(db *sql.DB, params SearchByParams) ([]Index, error) {
	args := []interface{}{}
	_sql := "select `name`, `platform`, `language`, `targets` from tldr_index where `name` like ? "
	args = append(args, params.Name +"%")
	contains := "{"
	if params.Platform != "" {
		contains += `"os": "` + params.Platform + `"`
	}
	if params.Language != "" {
		if len(contains) > 1 {
			contains += ","
		}
		contains += `"language":"` + params.Language + `"`
	}

	if len(contains) > 1 {
		_sql += " and json_contains(targets, '" + contains + "}')"
	}

	rows, err := db.Query(_sql, args...)
	if err != nil {
		logrus.Errorf("query err: %s", err)
		return nil, err
	}

	data := []Index{}
	for rows.Next() {
		var (
			name                        string
			platform, language, targets []uint8
		)
		if err := rows.Scan(&name, &platform, &language, &targets); err != nil {
			logrus.Errorf("scan err: %s", err)
			return nil, err
		}
		data = append(data, Index{
			Name:     name,
			Platform: platform,
			Language: language,
			Targets:  targets,
		})
	}
	return data, nil
}
