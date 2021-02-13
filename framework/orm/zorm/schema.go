package zorm

import (
	"fmt"
	"go/ast"
	"reflect"
)

type field struct {
	name  string
	ftype string
	tag   string
}

// Schema represents a database table.
type Schema struct {
	model      interface{}
	fieldNames []string
	fieldMap   map[string]*field
}

// Parse converts an Golang struct into a Schema object.
// Model should be a pointer.
func (s *Schema) Parse(model interface{}, d Dialect) {
	modelType := reflect.Indirect(reflect.ValueOf(model)).Type()
	fieldNum := modelType.NumField()
	s.model = model
	for i := 0; i < fieldNum; i++ {
		f := modelType.Field(i)
		injectFieldObject := &field{
			name:  f.Name,
			ftype: d.GolangTypeToDBType(reflect.Indirect(reflect.New(f.Type))),
		}
		if v, ok := f.Tag.Lookup("zorm"); ok {
			injectFieldObject.tag = v
		}
		if ast.IsExported(f.Name) && !f.Anonymous {
			s.fieldNames = append(s.fieldNames, f.Name)
			s.fieldMap[f.Name] = injectFieldObject
		}
	}
}

// NewSchema returns a new table.
func NewSchema() *Schema {
	return &Schema{fieldNames: make([]string, 0), fieldMap: make(map[string]*field)}
}

// PrintSchema shows the db table for debug use.
func (s *Schema) PrintSchema() {
	for _, v := range s.fieldMap {
		fmt.Printf("field[Name = %s, Type = %s, Tag = %s]\n", v.name, v.ftype, v.tag)
	}
}
