package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionUsuarios_20260521_145400 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionUsuarios_20260521_145400{}
	m.Created = "20260521_145400"

	migration.Register("CreacionUsuarios_20260521_145400", m)
}

// Run the migrations
func (m *CreacionUsuarios_20260521_145400) Up() {
	m.SQL(`INSERT INTO core.usuarios
    (nombre, apellido, correo, telefono, ciudad, fecha_nacimiento, id_genero, nacionalidad, foto_url, id_roles_usuarios)
	VALUES
    ('Laura',    'Martínez', 'laura.martinez@email.com', '3101234567', 'Bogotá',   '1995-04-12', 2, 'Colombiana', 'https://cdn.jobsy.co/fotos/laura.jpg',  1),
    ('Carlos',   'Rincón',   'carlos.rincon@email.com',  '3209876543', 'Medellín', '1988-11-03', 1, 'Colombiana', 'https://cdn.jobsy.co/fotos/carlos.jpg', 2),
    ('Valentina','Gómez',    'vgomez@email.com',          '3156784321', 'Cali',     '1992-07-25', 2, 'Colombiana', 'https://cdn.jobsy.co/fotos/vale.jpg',   1);`)

}

// Reverse the migrations
func (m *CreacionUsuarios_20260521_145400) Down() {
	m.SQL(`DELETE FROM core.usuarios WHERE nombre IN ('Laura', 'Carlos', 'Valentina');`)

}
