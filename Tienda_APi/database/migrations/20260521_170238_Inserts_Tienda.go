package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertsTienda_20260521_170238 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsTienda_20260521_170238{}
	m.Created = "20260521_170238"

	migration.Register("InsertsTienda_20260521_170238", m)
}

// Run the migrations
func (m *InsertsTienda_20260521_170238) Up() {
	file, err := ioutil.ReadFile("../20260521_170238_Inserts_Tienda_up.sql")
	if err != nil {
		fmt.Println(err)

	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}

}

// Reverse the migrations
func (m *InsertsTienda_20260521_170238) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../20260521_170238_Inserts_Tienda_down.sql")
	if err != nil {
		fmt.Println(err)

	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}

}
