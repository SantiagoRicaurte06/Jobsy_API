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

// ejecutarScript carga un .sql y encola cada sentencia.
//
// Antes esto leia un archivo con un nombre que no existe y solo imprimia el
// error, asi que la migracion "corria" ejecutando SQL vacio y no creaba
// ninguna tabla. Ahora un script ilegible aborta: es preferible fallar de
// entrada a creerse que la base quedo montada.
func ejecutarScript(m *CreacionApiCore_20260521_144404, ruta string) {
	file, err := ioutil.ReadFile(ruta)
	if err != nil {
		panic(fmt.Sprintf("no se pudo leer la migracion %s: %v", ruta, err))
	}

	for _, sentencia := range strings.Split(string(file), ";") {
		// El split deja un resto vacio despues del ultimo ';'.
		if strings.TrimSpace(sentencia) == "" {
			continue
		}
		m.SQL(sentencia)
	}
}

// Run the migrations
func (m *CreacionApiCore_20260521_144404) Up() {
	ejecutarScript(m, "../scripts/20260521_144404creacion_db_up.sql")
}

// Reverse the migrations
func (m *CreacionApiCore_20260521_144404) Down() {
	ejecutarScript(m, "../scripts/20260521_144404_creacion_db_down.sql")
}
