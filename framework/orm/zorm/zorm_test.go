package zorm_test

import (
	"testing"

	"github.com/galaxyzeta/framework/orm/zorm"
)

type user struct {
	Id   int
	Name string
}

func TestSchemaParse(t *testing.T) {
	table := zorm.NewSchema()
	table.Parse(&user{}, zorm.DialectMap["mysql"])
	table.PrintSchema()
}
