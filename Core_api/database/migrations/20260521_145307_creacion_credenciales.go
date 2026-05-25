package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionCredenciales_20260521_145307 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionCredenciales_20260521_145307{}
	m.Created = "20260521_145307"

	migration.Register("CreacionCredenciales_20260521_145307", m)
}

// Run the migrations
func (m *CreacionCredenciales_20260521_145307) Up() {
	m.SQL(`INSERT INTO core.roles_usuarios (nombre, descripcion) VALUES
    ('empleado',      'Usuario que ofrece servicios del hogar'),
    ('empleador',     'Usuario que publica ofertas y contrata empleados'),
    ('administrador', 'Administrador con acceso total a la plataforma');`)

}

// Reverse the migrations
func (m *CreacionCredenciales_20260521_145307) Down() {
	m.SQL(`DELETE FROM core.roles_usuarios WHERE nombre IN ('empleado', 'empleador', 'administrador');`)

}
