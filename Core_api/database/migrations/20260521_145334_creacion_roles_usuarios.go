package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionRolesUsuarios_20260521_145334 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionRolesUsuarios_20260521_145334{}
	m.Created = "20260521_145334"

	migration.Register("CreacionRolesUsuarios_20260521_145334", m)
}

// Run the migrations
func (m *CreacionRolesUsuarios_20260521_145334) Up() {
	m.SQL(`INSERT INTO core.roles_usuarios (nombre, descripcion) VALUES
    ('empleado',      'Usuario que ofrece servicios del hogar'),
    ('empleador',     'Usuario que publica ofertas y contrata empleados'),
    ('administrador', 'Administrador con acceso total a la plataforma');`)

}

// Reverse the migrations
func (m *CreacionRolesUsuarios_20260521_145334) Down() {
	m.SQL(`DELETE FROM core.roles_usuarios WHERE nombre IN ('empleado', 'empleador', 'administrador');`)

}
