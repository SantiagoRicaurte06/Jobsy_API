package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Estados_20260520_221642 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Estados_20260520_221642{}
	m.Created = "20260520_221642"

	migration.Register("Estados_20260520_221642", m)
}

// Run the migrations
func (m *Estados_20260520_221642) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Estados_20260520_221642) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
