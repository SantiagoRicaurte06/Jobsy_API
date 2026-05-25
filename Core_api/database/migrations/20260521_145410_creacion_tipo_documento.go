package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionTipoDocumento_20260521_145410 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionTipoDocumento_20260521_145410{}
	m.Created = "20260521_145410"

	migration.Register("CreacionTipoDocumento_20260521_145410", m)
}

// Run the migrations
func (m *CreacionTipoDocumento_20260521_145410) Up() {
	m.SQL(`INSERT INTO core.tipo_documento (nombre, descripcion) VALUES
    ('cedula',              'Cédula de ciudadanía colombiana'),
    ('pasaporte',           'Pasaporte vigente nacional o internacional'),
    ('tarjeta_extranjeria', 'Tarjeta de extranjería para ciudadanos no colombianos');`)

}

// Reverse the migrations
func (m *CreacionTipoDocumento_20260521_145410) Down() {
	m.SQL(`DELETE FROM core.tipo_documento WHERE nombre IN ('cedula', 'pasaporte', 'tarjeta_extranjeria');`)

}
