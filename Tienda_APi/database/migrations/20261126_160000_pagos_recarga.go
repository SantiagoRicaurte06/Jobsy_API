package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type PagosRecarga_20261126_160000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PagosRecarga_20261126_160000{}
	m.Created = "20261126_160000"

	migration.Register("PagosRecarga_20261126_160000", m)
}

// Up permite pagos sin pedido ni contratacion: las recargas del saldo Jobsy.
// Sigue prohibido que un pago cuelgue de ambos a la vez.
func (m *PagosRecarga_20261126_160000) Up() {
	m.SQL(`ALTER TABLE tienda.pagos DROP CONSTRAINT IF EXISTS ck_pagos_origen;`)
	m.SQL(`ALTER TABLE tienda.pagos ADD CONSTRAINT ck_pagos_origen
		CHECK (id_pedidos IS NULL OR id_contrataciones IS NULL);`)
}

// Down vuelve a exigir exactamente un origen (falla si ya hay recargas).
func (m *PagosRecarga_20261126_160000) Down() {
	m.SQL(`ALTER TABLE tienda.pagos DROP CONSTRAINT IF EXISTS ck_pagos_origen;`)
	m.SQL(`ALTER TABLE tienda.pagos ADD CONSTRAINT ck_pagos_origen
		CHECK ((id_pedidos IS NULL) <> (id_contrataciones IS NULL));`)
}
