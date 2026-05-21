package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionSchemaSuscripciones_20260521_150638 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionSchemaSuscripciones_20260521_150638{}
	m.Created = "20260521_150638"

	migration.Register("CreacionSchemaSuscripciones_20260521_150638", m)
}

// Run the migrations
func (m *CreacionSchemaSuscripciones_20260521_150638) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreacionSchemaSuscripciones_20260521_150638) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
