package model

import (
	"reflect"
	"database/sql"
	"github.com/iancoleman/strcase"
	"upper.io/db.v3/lib/sqlbuilder"
	"github.com/kr/pretty"
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
	collection := reflect.StructOf(fields)
	return collection
}

func (c *Collection) CreateColumns(columns []*sql.ColumnType, err error) *Collection {

	if err != nil { panic(err) }

	for _, column := range columns {
		typ := column.ScanType()

		pretty.Println(typ.String())

		if typ.String() == "sql.RawBytes" { typ = reflect.TypeOf([]byte{}) }
		c.Fields = append(c.Fields, reflect.StructField{
			Name: strcase.ToCamel(column.Name()),
			Type: typ,
			Tag: reflect.StructTag(`db:"` + column.Name() + `"`),
		})
	}
	return c
}

func (c *Collection) GetSliceCollection() interface{} {
	query , err := c.Selector.Query()
	Struct := c.CreateStruct(query, err)
	return reflect.New(reflect.SliceOf(Struct)).Interface()
}

func (c *Collection) GetSingleCollection() interface{} {
	query , err := c.Selector.Query()
	Struct := c.CreateStruct(query, err)
	return reflect.New(Struct).Interface()

	//arr := reflect.ValueOf(collection)
	//slices := arr.Elem().Slice(0, arr.Elem().Len())
	//columns, _ := query.ColumnTypes()
	//for i := 0; i < slices.Len(); i++ {
	//	slice := map[string] interface{} {}
	//	for si := 0; si < slices.Index(i).NumField(); si++ {
	//		field := slices.Index(i)
	//		slice[columns[si].Name()] = field.Field(si).Interface()
	//	}
	//	c.MapStruct = append(c.MapStruct, slice)
	//}
}

func (c *Collection) GetFields() []reflect.StructField {
	return c.Fields
}

func (c *Collection) GetSelector() sqlbuilder.Selector {
	return c.Selector
}