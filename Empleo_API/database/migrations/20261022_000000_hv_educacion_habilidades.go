package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type HvEducacionHabilidades_20261022_000000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &HvEducacionHabilidades_20261022_000000{}
	m.Created = "20261022_000000"

	migration.Register("HvEducacionHabilidades_20261022_000000", m)
}

// Run the migrations
//
// La hoja de vida del front tiene educacion y habilidades, pero el esquema no
// tenia donde guardarlas y se perdian al recargar. Se guardan como JSON en la
// propia hoja: son listas cortas que siempre se leen y escriben enteras.
func (m *HvEducacionHabilidades_20261022_000000) Up() {
	m.SQL(`ALTER TABLE empleo.hojas_vida
		ADD COLUMN IF NOT EXISTS educacion TEXT,
		ADD COLUMN IF NOT EXISTS habilidades TEXT;`)
}

// Reverse the migrations
func (m *HvEducacionHabilidades_20261022_000000) Down() {
	m.SQL(`ALTER TABLE empleo.hojas_vida
		DROP COLUMN IF EXISTS educacion,
		DROP COLUMN IF EXISTS habilidades;`)
}
