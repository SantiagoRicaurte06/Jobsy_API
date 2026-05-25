package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ModalidadesSuscripcion_20260520_221729 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModalidadesSuscripcion_20260520_221729{}
	m.Created = "20260520_221729"

	migration.Register("ModalidadesSuscripcion_20260520_221729", m)
}

// Run the migrations
func (m *ModalidadesSuscripcion_20260520_221729) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *ModalidadesSuscripcion_20260520_221729) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
