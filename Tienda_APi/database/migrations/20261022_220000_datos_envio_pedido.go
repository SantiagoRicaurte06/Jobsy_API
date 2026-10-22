package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type DatosEnvioPedido_20261022_220000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DatosEnvioPedido_20261022_220000{}
	m.Created = "20261022_220000"

	migration.Register("DatosEnvioPedido_20261022_220000", m)
}

// Run the migrations
//
// El pedido ya tenia nombre, telefono, direccion y ciudad de facturacion,
// pero para entregar hace falta el barrio, el detalle de la vivienda (apto,
// torre, casa) y las indicaciones para el repartidor.
func (m *DatosEnvioPedido_20261022_220000) Up() {
	m.SQL(`ALTER TABLE tienda.facturacion
		ADD COLUMN IF NOT EXISTS factura_barrio VARCHAR(100),
		ADD COLUMN IF NOT EXISTS factura_detalle VARCHAR(150),
		ADD COLUMN IF NOT EXISTS factura_notas VARCHAR(300);`)
}

// Reverse the migrations
func (m *DatosEnvioPedido_20261022_220000) Down() {
	m.SQL(`ALTER TABLE tienda.facturacion
		DROP COLUMN IF EXISTS factura_barrio,
		DROP COLUMN IF EXISTS factura_detalle,
		DROP COLUMN IF EXISTS factura_notas;`)
}
