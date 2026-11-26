package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type SuscriptoresBoletin_20261126_150000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &SuscriptoresBoletin_20261126_150000{}
	m.Created = "20261126_150000"

	migration.Register("SuscriptoresBoletin_20261126_150000", m)
}

// Up crea la tabla usuarios.suscriptores_boletin.
func (m *SuscriptoresBoletin_20261126_150000) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS usuarios.suscriptores_boletin (
		id_suscriptores_boletin SERIAL PRIMARY KEY,
		correo VARCHAR(150) NOT NULL,
		activo BOOLEAN NOT NULL DEFAULT true,
		creado_en TIMESTAMP NOT NULL DEFAULT NOW(),
		actualizado_en TIMESTAMP NOT NULL DEFAULT NOW()
	);`)
	m.SQL(`CREATE UNIQUE INDEX IF NOT EXISTS ux_suscriptores_boletin_correo
		ON usuarios.suscriptores_boletin (lower(correo));`)
}

// Down revierte el Up.
func (m *SuscriptoresBoletin_20261126_150000) Down() {
	m.SQL(`DROP TABLE IF EXISTS usuarios.suscriptores_boletin;`)
}
