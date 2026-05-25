package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionApiCore_20260521_144404 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionApiCore_20260521_144404{}
	m.Created = "20260521_144404"

	migration.Register("CreacionApiCore_20260521_144404", m)
}

// Run the migrations
func (m *CreacionApiCore_20260521_144404) Up() {
	file, err:= ioutil.ReadFile("../scripts/20260514_125038_creacion_db_up.sql")
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
func (m *CreacionApiCore_20260521_144404) Down() {
	file, err:= ioutil.ReadFile("../scripts/20260514_125038_creacion_db_down.sql")
	if err != nil {
		fmt.Println(err)
	}
	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}		

}
