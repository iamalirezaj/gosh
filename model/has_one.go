package model

import (
	"upper.io/db.v3/lib/sqlbuilder"
)

type HasOneRelation struct {
	Model ModelInterface
}

func (HasOneRelation) Prepare(selector sqlbuilder.Selector, TableName string, t1column string, t2column string) sqlbuilder.Selector {

	return selector.LeftJoin(TableName + " AS t2 ").On("t1." + t1column + " = t2." + t2column)
}