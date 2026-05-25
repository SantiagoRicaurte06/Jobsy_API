package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionDocumentosIdentidad_20260521_145322 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDocumentosIdentidad_20260521_145322{}
	m.Created = "20260521_145322"

	migration.Register("CreacionDocumentosIdentidad_20260521_145322", m)
}

// Run the migrations
func (m *CreacionDocumentosIdentidad_20260521_145322) Up() {
	m.SQL(`INSERT INTO core.documentos_identidad
    (id_usuarios, id_tipo_documento, numero_documento, fecha_expedicion, lugar_expedicion)
VALUES
    (1, 1, '1020304050', '2013-06-15', 'Bogotá'),
    (2, 1, '1098765432', '2008-03-20', 'Medellín'),
    (3, 1, '1044556677', '2010-09-10', 'Cali');`)

}

// Reverse the migrations
func (m *CreacionDocumentosIdentidad_20260521_145322) Down() {
	m.SQL(`DELETE FROM core.documentos_identidad WHERE id_usuarios IN (1, 2, 3);`)

}
