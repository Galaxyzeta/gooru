package zorm

import (
	"strings"

	"github.com/galaxyzeta/util/common"
)

// SQLGenerator generates sql clauses.
type SQLGenerator struct {
	sql common.StringBuilder
	res string
}

// Select generates a SELECT clause.
func (g *SQLGenerator) Select(table string, cols ...string) *SQLGenerator {
	g.sql.WriteString("SELECT ").WriteString(strings.Join(cols, ",")).WriteString(" FROM ").WriteString(table)
	return g
}
