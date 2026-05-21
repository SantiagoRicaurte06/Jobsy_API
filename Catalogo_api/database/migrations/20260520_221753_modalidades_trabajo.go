package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ModalidadesTrabajo_20260520_221753 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModalidadesTrabajo_20260520_221753{}
	m.Created = "20260520_221753"

	migration.Register("ModalidadesTrabajo_20260520_221753", m)
}

// Run the migrations
func (m *ModalidadesTrabajo_20260520_221753) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *ModalidadesTrabajo_20260520_221753) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
