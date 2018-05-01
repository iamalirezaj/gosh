package model

import (
	"reflect"
	"database/sql"
	"github.com/kr/pretty"
	"github.com/iancoleman/strcase"
	"upper.io/db.v3/lib/sqlbuilder"
)

type Collection struct {
	Fields []reflect.StructField
	Selector sqlbuilder.Selector
}

func (c *Collection) SetSelector(selector sqlbuilder.Selector) *Collection {
	c.Selector = selector
	return c
}

func (c *Collection) CreateStruct(query *sql.Rows, err error) reflect.Type {
	if err != nil { panic(err) }
	fields := c.CreateColumns(query.ColumnTypes()).Fields
	return reflect.StructOf(fields)
}

func (c *Collection) CreateColumns(columns []*sql.ColumnType, err error) *Collection {

	if err != nil { panic(err) }

	for _, column := range columns {
		typ := column.ScanType()

		if typ.Name() == "RawBytes" {

			pretty.Println(typ)

			typ = reflect.TypeOf([]byte{})
		}

		c.Fields = append(c.Fields, reflect.StructField{
			Name: strcase.ToCamel(column.Name()),
			Type: typ,
			Tag: reflect.StructTag(`db:"` + column.Name() + `" json:"` + column.Name() + `"`),
		})
	}

	return c
}

func (c *Collection) GetStruct() reflect.Type {
	return c.CreateStruct(c.Selector.Query())
}

func (c *Collection) GetSliceCollection() interface{} {
	return reflect.New(reflect.SliceOf(c.GetStruct())).Interface()
}

func (c *Collection) GetSingleCollection() interface{} {
	return reflect.New(c.GetStruct()).Interface()
}

func (c *Collection) GetFields() []reflect.StructField {
	return c.Fields
}

func (c *Collection) GetSelector() sqlbuilder.Selector {
	return c.Selector
}