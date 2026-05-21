package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type NivelesExperiencia_20260520_221817 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &NivelesExperiencia_20260520_221817{}
	m.Created = "20260520_221817"

	migration.Register("NivelesExperiencia_20260520_221817", m)
}

// Run the migrations
func (m *NivelesExperiencia_20260520_221817) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *NivelesExperiencia_20260520_221817) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
