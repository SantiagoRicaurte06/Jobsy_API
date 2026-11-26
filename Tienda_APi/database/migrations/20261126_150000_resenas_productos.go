package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ResenasProductos_20261126_150000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ResenasProductos_20261126_150000{}
	m.Created = "20261126_150000"

	migration.Register("ResenasProductos_20261126_150000", m)
}

// Up crea la tabla tienda.resenas_productos.
func (m *ResenasProductos_20261126_150000) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS tienda.resenas_productos (
		id_resenas_productos SERIAL PRIMARY KEY,
		id_productos INT NOT NULL REFERENCES tienda.productos(id_productos),
		id_usuarios INT NOT NULL REFERENCES core.usuarios(id_usuarios),
		autor VARCHAR(120) NOT NULL,
		calificacion INT NOT NULL CHECK (calificacion BETWEEN 1 AND 5),
		titulo VARCHAR(120) NOT NULL,
		comentario TEXT NOT NULL,
		fotos TEXT NOT NULL DEFAULT '[]',
		verificado BOOLEAN NOT NULL DEFAULT false,
		activo BOOLEAN NOT NULL DEFAULT true,
		creado_en TIMESTAMP NOT NULL DEFAULT NOW(),
		actualizado_en TIMESTAMP NOT NULL DEFAULT NOW()
	);`)
	m.SQL(`CREATE UNIQUE INDEX IF NOT EXISTS ux_resenas_productos_usuario
		ON tienda.resenas_productos (id_productos, id_usuarios);`)
}

// Down revierte el Up.
func (m *ResenasProductos_20261126_150000) Down() {
	m.SQL(`DROP TABLE IF EXISTS tienda.resenas_productos;`)
}
