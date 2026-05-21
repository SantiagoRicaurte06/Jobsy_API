package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Turnos_20260520_221826 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Turnos_20260520_221826{}
	m.Created = "20260520_221826"

	migration.Register("Turnos_20260520_221826", m)
}

// Run the migrations
func (m *Turnos_20260520_221826) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Turnos_20260520_221826) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
