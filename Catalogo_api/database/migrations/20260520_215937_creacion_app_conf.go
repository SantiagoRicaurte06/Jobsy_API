package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionAppConf_20260520_215937 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionAppConf_20260520_215937{}
	m.Created = "20260520_215937"

	migration.Register("CreacionAppConf_20260520_215937", m)
}

// Run the migrations
func (m *CreacionAppConf_20260520_215937) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreacionAppConf_20260520_215937) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
