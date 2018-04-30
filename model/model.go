package model

import (
	"gosh"
	"reflect"
	"gosh/words"
	"upper.io/db.v3/lib/sqlbuilder"
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

func (Model) Connection() sqlbuilder.Database {
	return gosh.Container.Connection
}

func (m Model) Get() interface{} {
	collection := m.GetCollection()
	attributes := collection.GetSingleCollection()
	m.Selector.One(attributes)
	return attributes
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
			//RelationModel := reflect.ValueOf(relation).FieldByName("Model").Interface()
			//RelationData := reflect.ValueOf(RelationModel).MethodByName("Find").Call([]reflect.Value {
			//	reflect.ValueOf(1),
			//})[0].MethodByName("First").Call([]reflect.Value {})[0]
			//m = m.SetFields([]reflect.StructField {{
			//	Name: "Role",
			//	Type: reflect.TypeOf(reflect.Interface),
			//}})
			//Struct := m

			//pretty.Println(Struct)
			//.Elem().FieldByName("Role").Set(RelationData)
		}
	}

	return m
}

func (m Model) Find(id int) Model {
	m.Selector = m.Connection().SelectFrom(m.Table + " AS t1 ").Where("t1.id = ?", id)
	return m
}

func (m Model) All() interface{} {
	m.Selector = m.Connection().SelectFrom(m.Table)
	collection := m.GetCollection()
	attributes := collection.GetSliceCollection()
	m.Selector.All(attributes)
	return attributes
}

func (m Model) SetFields(fields []reflect.StructField) Model {
	m.Collection.Fields = append(m.Collection.Fields, fields...)
	return m
}

func (m Model) GetCollection() *Collection {
	return m.Collection.SetSelector(m.Selector)
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