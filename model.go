package gosh

import (
	"os"
	"reflect"
	"gosh/words"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/lib/sqlbuilder"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kr/pretty"
)

type Model struct {
	Model ModelInterface
	Table string
	Hidden []string
	*Collection
}

type ModelInterface interface {}

func connection() (db sqlbuilder.Database) {

	db, _ = sqlbuilder.Open("mysql", mysql.ConnectionURL{
		Database: os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	})
	return
}

func (m Model) relationships(data ModelInterface) []int {

	relationships := []int {}

	d := reflect.ValueOf(data)
	for i := 0; i < d.NumField(); i++ {
		typ := d.Field(i).Type()
		mod := reflect.TypeOf(new(Model))
		if typ.Kind() == reflect.Struct {
			for n := 0; n < typ.NumField(); n++ {
				if typ.Field(n).Type == mod {
					relationships = append(relationships, i)
				}
			}
		}
	}

	return relationships
}

func (m Model) All() interface{} {
	selector := connection().SelectFrom(m.Table)
	datas := reflect.New(reflect.SliceOf(reflect.TypeOf(m.Model))).Interface()
	selector.All(datas)
	return datas
}

func NewModel(model ModelInterface) *Model {
	return &Model{
		Model: model,
		Table: GetTableName(model),
		Collection: &Collection{},
	}
}

func GetBaseTableName(model ModelInterface) string {
	var BaseTableName string
	Func := reflect.ValueOf(model).MethodByName("TableName")
	if Func.IsValid() { BaseTableName = Func.Call([]reflect.Value {})[0].String() }
	return BaseTableName
}

func GetTableName(model ModelInterface) string {
	if TableName := GetBaseTableName(model); TableName != "" {
		return TableName
	}
	Value := words.String{ Value: reflect.TypeOf(model).Name(), }
	return Value.ToPlural().ToLowercase().ToString()
}

func GetHidden(model ModelInterface) []string {
	var Hidden []string
	handle := reflect.ValueOf(model).MethodByName("Hidden")
	hide := handle.Call([]reflect.Value{})[0]
	for i := 0; i < hide.Len(); i++ {
		Hidden = append(Hidden, hide.Index(i).String())
	}
	return Hidden
}