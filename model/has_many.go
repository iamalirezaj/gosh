package model

import (
	"upper.io/db.v3/lib/sqlbuilder"
)

type HasManyRelation struct {
	Model ModelInterface
}

func (HasManyRelation) Prepare(selector sqlbuilder.Selector) sqlbuilder.Selector {

	return selector
}