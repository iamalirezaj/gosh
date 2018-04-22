package gosh

import (
	"os"
	"reflect"
	"gosh/words"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/lib/sqlbuilder"
	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	Model ModelInterface
	Table string
	Hidden []string
	ModelInterface
	*Collection
}

type ModelInterface interface {
	New() *Model
}

func connection() (db sqlbuilder.Database) {

	db, _ = sqlbuilder.Open("mysql", mysql.ConnectionURL{
		Database: os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	})
	return
}

func (m Model) All() *Collection {
	selector := connection().Select("*").From("users")
	collections := m.Collection.SetSelector(selector)
	return collections.Run()
}

func NewModel(model ModelInterface) *Model {
	return &Model{
		Model: model,
		Table: GetTableName(model),
		Hidden: GetHidden(model),
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