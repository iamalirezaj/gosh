package gosh

import (
	"database/sql"
	"reflect"
	"github.com/iancoleman/strcase"
	"upper.io/db.v3/lib/sqlbuilder"
)

type Collection struct {
	Fields []reflect.StructField
	Struct interface{}
	MapStruct []map[string] interface{}
	Selector sqlbuilder.Selector
}

func (c *Collection) SetSelector(selector sqlbuilder.Selector) *Collection {
	c.Selector = selector
	return c
}

func (c *Collection) CreateStruct(query *sql.Rows, err error) interface{} {
	if err != nil { panic(err) }
	fields := c.CreateColumns(query.ColumnTypes()).Fields
	collection := reflect.StructOf(fields)
	return reflect.New(reflect.SliceOf(collection)).Interface()
}

func (c *Collection) CreateColumns(columns []*sql.ColumnType, err error) *Collection {

	if err != nil { panic(err) }

	for _, column := range columns {
		typ := column.ScanType()
		if typ.String() == "sql.RawBytes" { typ = reflect.TypeOf([]uint8{}) }
		c.Fields = append(c.Fields, reflect.StructField{
			Name: strcase.ToCamel(column.Name()),
			Type: typ,
			Tag: reflect.StructTag(`db:"` + column.Name() + `"`),
		})
	}
	return c
}

func (c *Collection) ToArray() []map[string] interface{} {
	return c.MapStruct
}

func (c *Collection) Run() *Collection {
	query , err := c.Selector.Query()
	collection := c.CreateStruct(query, err)
	c.Selector.All(collection)
	c.Struct = collection

	arr := reflect.ValueOf(collection)
	slices := arr.Elem().Slice(0, arr.Elem().Len())
	columns, _ := query.ColumnTypes()

	for i := 0; i < slices.Len(); i++ {
		slice := map[string] interface{} {}
		for si := 0; si < slices.Index(i).NumField(); si++ {
			field := slices.Index(i)
			slice[columns[si].Name()] = field.Field(si).Interface()
		}
		c.MapStruct = append(c.MapStruct, slice)
	}

	return c
}

func (c *Collection) GetFields() []reflect.StructField {
	return c.Fields
}

func (c *Collection) GetSelector() sqlbuilder.Selector {
	return c.Selector
}

func (c *Collection) ToStruct() interface{} {
	return c.Struct
}