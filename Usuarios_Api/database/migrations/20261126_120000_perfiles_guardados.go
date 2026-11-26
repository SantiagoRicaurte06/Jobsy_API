package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type PerfilesGuardados_20261126_120000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PerfilesGuardados_20261126_120000{}
	m.Created = "20261126_120000"

	migration.Register("PerfilesGuardados_20261126_120000", m)
}

// Up crea la lista de perfiles de empleados que cada usuario guarda para
// verlos despues ("Guardar perfil" en el detalle del empleado).
func (m *PerfilesGuardados_20261126_120000) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS usuarios.perfiles_guardados (
		id_perfiles_guardados SERIAL PRIMARY KEY,
		id_usuarios INT NOT NULL REFERENCES core.usuarios(id_usuarios),
		id_empleado INT NOT NULL REFERENCES core.usuarios(id_usuarios),
		activo BOOLEAN NOT NULL DEFAULT true,
		creado_en TIMESTAMP NOT NULL DEFAULT NOW(),
		actualizado_en TIMESTAMP NOT NULL DEFAULT NOW()
	);`)
	// Un mismo perfil se guarda una sola vez por usuario.
	m.SQL(`CREATE UNIQUE INDEX IF NOT EXISTS ux_perfiles_guardados_usuario_empleado
		ON usuarios.perfiles_guardados (id_usuarios, id_empleado);`)
}

// Down revierte el Up.
func (m *PerfilesGuardados_20261126_120000) Down() {
	m.SQL(`DROP TABLE IF EXISTS usuarios.perfiles_guardados;`)
}
