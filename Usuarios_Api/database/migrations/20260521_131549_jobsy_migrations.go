package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type JobsyMigrations_20260521_131549 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &JobsyMigrations_20260521_131549{}
	m.Created = "20260521_131549"

	migration.Register("JobsyMigrations_20260521_131549", m)
}

// Run the migrations
//
// ORDEN DE EJECUCION: depende de los esquemas core y catalogo, que deben
// crearse antes.
//
// Se corrigieron tres cosas respecto a la version anterior:
//
//  1. "CREATE SCHEMA IF NOT EXISTS usuarios\n);" tenia un parentesis de cierre
//     suelto: era un error de sintaxis y el esquema no se creaba.
//  2. Las FK decian REFERENCES usuarios(id_usuarios), que apunta a una TABLA
//     llamada usuarios en el search_path. La tabla real es core.usuarios; el
//     nombre coincidia con el del esquema y por eso pasaba desapercibido.
//  3. La tabla notificaciones declaraba las columnas mensaje y leido, pero el
//     modelo Go (models/notificaciones.go) usa cuerpo y leida, y ademas espera
//     id_tipo y url_destino, que no se creaban. Cualquier consulta del ORM
//     fallaba con "column does not exist".
func (m *JobsyMigrations_20260521_131549) Up() {
	m.SQL(`CREATE SCHEMA IF NOT EXISTS usuarios;`)

	m.SQL(`CREATE TABLE IF NOT EXISTS usuarios.configuraciones (
		id_configuraciones SERIAL PRIMARY KEY,
		id_usuarios INT NOT NULL REFERENCES core.usuarios(id_usuarios),
		perfil_publico BOOLEAN NOT NULL DEFAULT false,
		mostrar_correo BOOLEAN NOT NULL DEFAULT false,
		mostrar_telefono BOOLEAN NOT NULL DEFAULT false,
		hv_visible BOOLEAN NOT NULL DEFAULT false,
		aparecer_busquedas BOOLEAN NOT NULL DEFAULT false,
		mostrar_ubicacion BOOLEAN NOT NULL DEFAULT false,
		doble_factor BOOLEAN NOT NULL DEFAULT false,
		notif_correo BOOLEAN NOT NULL DEFAULT false,
		notif_plataforma BOOLEAN NOT NULL DEFAULT false,
		notif_nuevas_ofertas BOOLEAN NOT NULL DEFAULT false,
		notif_actualizaciones BOOLEAN NOT NULL DEFAULT false,
		notif_mensajes BOOLEAN NOT NULL DEFAULT false,
		notif_postulaciones BOOLEAN NOT NULL DEFAULT false,
		activo BOOLEAN NOT NULL DEFAULT true,
		creado_en TIMESTAMP NOT NULL DEFAULT NOW(),
		actualizado_en TIMESTAMP NOT NULL DEFAULT NOW()
	);`)

	// Una fila de configuracion por usuario: con dos, no hay forma de saber
	// cual manda.
	m.SQL(`CREATE UNIQUE INDEX IF NOT EXISTS ux_configuraciones_usuario
		ON usuarios.configuraciones (id_usuarios);`)

	m.SQL(`CREATE TABLE IF NOT EXISTS usuarios.notificaciones (
		id_notificaciones SERIAL PRIMARY KEY,
		id_usuarios INT NOT NULL REFERENCES core.usuarios(id_usuarios),
		id_tipo INT NOT NULL REFERENCES catalogo.tipos(id_tipo),
		titulo VARCHAR(255) NOT NULL,
		cuerpo TEXT,
		leida BOOLEAN NOT NULL DEFAULT false,
		url_destino VARCHAR(300),
		activo BOOLEAN NOT NULL DEFAULT true,
		creado_en TIMESTAMP NOT NULL DEFAULT NOW(),
		actualizado_en TIMESTAMP NOT NULL DEFAULT NOW()
	);`)

	// La bandeja se consulta siempre por usuario y por no leidas.
	m.SQL(`CREATE INDEX IF NOT EXISTS ix_notificaciones_usuario
		ON usuarios.notificaciones (id_usuarios);`)
	m.SQL(`CREATE INDEX IF NOT EXISTS ix_notificaciones_no_leidas
		ON usuarios.notificaciones (id_usuarios) WHERE NOT leida AND activo;`)
}

// Reverse the migrations
//
// Los nombres iban sin esquema, asi que dependian del search_path y podian
// borrar una tabla equivocada (o ninguna).
func (m *JobsyMigrations_20260521_131549) Down() {
	m.SQL(`DROP TABLE IF EXISTS usuarios.notificaciones;`)
	m.SQL(`DROP TABLE IF EXISTS usuarios.configuraciones;`)
	m.SQL(`DROP SCHEMA IF EXISTS usuarios;`)
}
