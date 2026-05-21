package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Generos_20260520_221804 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Generos_20260520_221804{}
	m.Created = "20260520_221804"

	migration.Register("Generos_20260520_221804", m)
}

// Run the migrations
func (m *Generos_20260520_221804) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Generos_20260520_221804) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
