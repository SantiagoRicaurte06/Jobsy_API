package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Insert_Schema_Soporte_20260521_132443 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Insert_Schema_Soporte_20260521_132443{}
	m.Created = "20260514_125215"
	migration.Register("CreacionDb_20260514_125215", m)
}

// Run the migrations
func (m *Insert_Schema_Soporte_20260521_132443) Up() {
	// Se lee el archivo SQL
	file, err := ioutil.ReadFile("../20260521_132443_Insert_Schema_Soporte_up.sql")

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
func (m *Insert_Schema_Soporte_20260521_132443) Down() {
	// Se lee el archivo SQL
	file, err := ioutil.ReadFile("../20260521_132443_Insert_Schema_Soporte_down.sql")

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

