package database

import "database/sql"

type Content struct {
	Name string
	Platform string
	Language string
	Content string
}

func (c *Content) GetDistinctPlatformBy(db *sql.DB, language string) ([]string, error) {
	_sql := "select distinct `platform` from tldr_content "
	args := []interface{}{}
	if language != "" {
		_sql += "where `language` = ?"
		args = append(args, language)
	}

	rows, err := db.Query(_sql, args...)
	if err != nil {
		return nil, err
	}

	platforms := []string{}
	for rows.Next() {
		var platform string
		if err := rows.Scan(&platform); err != nil {
			return nil, err
		}
		platforms = append(platforms, platform)
	}
	return platforms, nil
}

func (c *Content) GetDistinctLanguageBy(db *sql.DB, platform string) ([]string, error) {
	_sql := "select distinct `language` from tldr_content "
	args := []interface{}{}
	if platform != "" {
		_sql += "where platform = ?"
		args = append(args, platform)
	}

	rows, err := db.Query(_sql, args...)
	if err != nil {
		return nil, err
	}

	languages := []string{}
	for rows.Next() {
		var language string
		if err := rows.Scan(&language); err != nil {
			return nil, err
		}
		languages = append(languages, language)
	}
	return languages, nil
}

type GetContentParams struct {
	Name string
	Platform string
	Language string
}


func (c *Content) GetContent(db *sql.DB, params GetContentParams)  {

}
