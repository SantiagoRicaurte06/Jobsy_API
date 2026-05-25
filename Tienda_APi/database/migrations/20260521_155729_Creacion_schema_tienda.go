package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionSchemaTienda_20260521_155729 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionSchemaTienda_20260521_155729{}
	m.Created = "20260521_155729"

	migration.Register("CreacionSchemaTienda_20260521_155729", m)
}

// Run the migrations
func (m *CreacionSchemaTienda_20260521_155729) Up() {
	file, err := ioutil.ReadFile("../20260521_155729_Creacion_schema_tienda_up.sql")
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
func (m *CreacionSchemaTienda_20260521_155729) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../20260521_155729_Creacion_schema_tienda_down.sql")
	if err != nil {
		fmt.Println(err)

	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}

}
