package model

import (
	"os"
	"reflect"
	"gosh/words"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/lib/sqlbuilder"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/kr/pretty"
)

type Model struct {
	Model reflect.Value
	Table string
	Hidden []string
	Selector sqlbuilder.Selector
	Attributes interface{}
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
	return m.Attributes
}

func (m Model) First() interface{} {
	return m.Get()
}

func (m Model) Load(relationships ...string) Model {

	for _, relationName := range relationships {
		method := m.Model.MethodByName(relationName)
		relation := method.Call([]reflect.Value {})[0].Interface()

		switch reflect.TypeOf(relation) {
		case reflect.TypeOf(HasOneRelation{}):
			RelationModel := reflect.ValueOf(relation).FieldByName("Model").Interface()
			RelationData := reflect.ValueOf(RelationModel).MethodByName("Find").Call([]reflect.Value {
				reflect.ValueOf(1),
			})[0].MethodByName("First").Call([]reflect.Value {})[0]
			m.Model.Elem().FieldByName("Rolel").Set(RelationData)
		}
	}

	return m
}

func (m Model) Find(id int) Model {
	selector := connection().SelectFrom(m.Table + " AS t1 ").Where("t1.id = ?", id)
	m.Selector = selector
	datas := m.Model.Interface()
	m.Selector.One(datas)
	m.Attributes = datas
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
		Model: reflect.New(reflect.TypeOf(model)),
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