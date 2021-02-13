package zorm

import (
	"fmt"
	"reflect"
	"time"
)

// Dialect represents type convertion regulations between database and golang.
type Dialect interface {
	GolangTypeToDBType(reflect.Value) string
}

// DialectMap is a collection of type conversion rules.
var DialectMap map[string]Dialect

type mysql struct{}

func init() {
	DialectMap = make(map[string]Dialect)
	DialectMap["mysql"] = &mysql{}
}

func (d *mysql) GolangTypeToDBType(golangType reflect.Value) string {
	switch golangType.Kind() {
	case reflect.Int, reflect.Uint:
		return "int"
	case reflect.Int8, reflect.Uint8:
		return "tinyint"
	case reflect.Int16, reflect.Uint16:
		return "smallint"
	case reflect.Int32, reflect.Uint32:
		return "int"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32:
		return "float"
	case reflect.Float64:
		return "double"
	case reflect.String:
		return "varchar"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := golangType.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("Invalid SQL type: %v", golangType))
}
