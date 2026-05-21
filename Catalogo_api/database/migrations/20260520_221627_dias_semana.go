package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type DiasSemana_20260520_221627 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DiasSemana_20260520_221627{}
	m.Created = "20260520_221627"

	migration.Register("DiasSemana_20260520_221627", m)
}

// Run the migrations
func (m *DiasSemana_20260520_221627) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *DiasSemana_20260520_221627) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
