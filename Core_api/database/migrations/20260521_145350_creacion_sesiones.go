package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionSesiones_20260521_145350 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionSesiones_20260521_145350{}
	m.Created = "20260521_145350"

	migration.Register("CreacionSesiones_20260521_145350", m)
}

// Run the migrations
func (m *CreacionSesiones_20260521_145350) Up() {
	m.SQL(`INSERT INTO core.sesiones
    (id_usuarios, token, dispositivo, ip, expira_en)
VALUES
    (1, 'tok_laura_abc123xyz',  'Chrome / Windows 11', '192.168.1.10', NOW() + INTERVAL '7 days'),
    (2, 'tok_carlos_def456uvw', 'Safari / iPhone 15',  '192.168.1.20', NOW() + INTERVAL '7 days'),
    (3, 'tok_vale_ghi789rst',   'Firefox / macOS 14',  '192.168.1.30', NOW() + INTERVAL '7 days');`)

}

// Reverse the migrations
func (m *CreacionSesiones_20260521_145350) Down() {
	m.SQL(`DELETE FROM core.sesiones WHERE token IN ('tok_laura_abc123xyz', 'tok_carlos_def456uvw', 'tok_vale_ghi789rst');`)	

}
