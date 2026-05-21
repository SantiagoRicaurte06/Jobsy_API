package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
	"fmt"
	"io/ioutil"
	"strings"
)

// DO NOT MODIFY
type EmpleoAPI_20260520_161001 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EmpleoAPI_20260520_161001{}
	m.Created = "20260520_161001"

	migration.Register("EmpleoAPI_20260520_161001", m)
}

// Run the migrations
func (m *EmpleoAPI_20260520_161001) Up() {
	// Se lee el archivo SQL
	file, err := ioutil.ReadFile("../20260520_161001_Empleo_API_up.sql")

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
func (m *EmpleoAPI_20260520_161001) Down() {
	// Se lee el archivo SQL
	file, err := ioutil.ReadFile("../20260520_161001_Empleo_API_down.sql")

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
