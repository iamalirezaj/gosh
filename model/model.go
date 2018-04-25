package model

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
	Selector sqlbuilder.Selector
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

func (m Model) Get() interface{} {
	datas := reflect.New(reflect.TypeOf(m.Model)).Interface()
	m.Selector.One(datas)
	return datas
}

func (m Model) First() interface{} {
	return m.Get()
}

func (m Model) Load(relationships ...string) Model {

	datas := reflect.ValueOf(m.Model)

	for _, relationName := range relationships {
		method := datas.MethodByName(relationName)
		relation := method.Call([]reflect.Value {})[0].Interface()

		switch reflect.TypeOf(relation) {
		case reflect.TypeOf(HasOneRelation{}):
			relationModel := reflect.ValueOf(relation).FieldByName("Model").Interface()
			TableName := reflect.ValueOf(relationModel).FieldByName("Table").String()
			m.Selector = HasOneRelation{}.Prepare(m.Selector, TableName, "role_id", "id")
		}
	}

	return m
}

func (m Model) Find(id int) Model {
	selector := connection().SelectFrom(m.Table + " AS t1 ").Where("t1.id = ?", id)
	m.Selector = selector
	return m
}

func (m Model) All() interface{} {
	selector := connection().SelectFrom(m.Table)
	datas := reflect.New(reflect.SliceOf(reflect.TypeOf(m.Model))).Interface()
	selector.All(datas)
	return datas
}

func HasMany(model ModelInterface) HasManyRelation {
	return HasManyRelation{ Model: NewModel(model) }
}

func HasOne(model ModelInterface) HasOneRelation {
	return HasOneRelation{ Model: NewModel(model) }
}

func NewModel(model ModelInterface) Model {
	return Model{
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