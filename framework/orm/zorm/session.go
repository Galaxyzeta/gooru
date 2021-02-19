package zorm

import (
	"database/sql"

	"github.com/galaxyzeta/util/common"
)

// Session represents a transcational query to database.
type Session struct {
	db         *sql.DB
	sql        common.StringBuilder
	s          *Schema
	params     []interface{}
	selectCols []string
}

// Select generates a SELECT clause.
func (s *Session) Select(table string, cols ...string) *Session {
	s.selectCols = append(s.selectCols, cols...)
	return s
}

// Find executes a SQL query and reflects result set into struct pointed by parameter elem.
func (s *Session) Find(elem []interface{}) error {
	res, err := s.db.Query(s.sql.String(), s.params...)
	if err != nil {
		return err
	}
	defer res.Close()
	target := make([]interface{}, len(s.selectCols))
	for res.Next() != false {
		err = res.Scan(target)
		if err != nil {
			return err
		}

	}
	return nil
}
