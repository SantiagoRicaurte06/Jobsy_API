package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)
// DO NOT MODIFY
type CreacionSchemaSuscripciones_20260521_150638 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionSchemaSuscripciones_20260521_150638{}
	m.Created = "20260521_150638"

	migration.Register("CreacionSchemaSuscripciones_20260521_150638", m)
}

// Run the migrations
func (m *CreacionSchemaSuscripciones_20260521_150638) Up() {
	// Se lee el archivo SQL
	file, err := ioutil.ReadFile("../20260521_150638_Creacion_Schema_Suscripciones_up.sql")

	// Se verifica si hubo un error al leer el archivo
	if err != nil {
		fmt.Println(err)
	}

	// Se separa cada linea del archivo SQL y cuando encuentre un ; lo divide
	requests := strings.Split(string(file), ";")

	// Se ejecuta cada linea SQL
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *CreacionSchemaSuscripciones_20260521_150638) Down() {
	// Se lee el archivo SQL
	file, err := ioutil.ReadFile("../20260521_150638_Creacion_Schema_Suscripciones_down.sql")

	// Se verifica si hubo un error al leer el archivo
	if err != nil {
		fmt.Println(err)
	}

	// Se separa cada linea del archivo SQL y cuando encuentre un ; lo divide
	requests := strings.Split(string(file), ";")

	// Se ejecuta cada linea SQL
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}