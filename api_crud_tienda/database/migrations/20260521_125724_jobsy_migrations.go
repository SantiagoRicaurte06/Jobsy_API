package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type JobsyMigrations_20260521_125724 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &JobsyMigrations_20260521_125724{}
	m.Created = "20260521_125724"

	migration.Register("JobsyMigrations_20260521_125724", m)
}

// Run the migrations
func (m *JobsyMigrations_20260521_125724) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS usuarios (
		id_usuarios SERIAL PRIMARY KEY,
		nombre VARCHAR(100) NOT NULL,
		correo VARCHAR(100) UNIQUE NOT NULL,
		contrasena VARCHAR(255) NOT NULL,
		telefono VARCHAR(20),
		activo BOOLEAN DEFAULT true,
		creado_en TIMESTAMP DEFAULT NOW(),
		actualizado_en TIMESTAMP DEFAULT NOW()
	);`)

	m.SQL(`CREATE TABLE IF NOT EXISTS configuraciones (
		id_configuraciones SERIAL PRIMARY KEY,
		id_usuarios INT REFERENCES usuarios(id_usuarios),
		perfil_publico BOOLEAN DEFAULT false,
		mostrar_correo BOOLEAN DEFAULT false,
		mostrar_telefono BOOLEAN DEFAULT false,
		hv_visible BOOLEAN DEFAULT false,
		aparecer_busquedas BOOLEAN DEFAULT false,
		mostrar_ubicacion BOOLEAN DEFAULT false,
		doble_factor BOOLEAN DEFAULT false,
		notif_correo BOOLEAN DEFAULT false,
		notif_plataforma BOOLEAN DEFAULT false,
		notif_nuevas_ofertas BOOLEAN DEFAULT false,
		notif_actualizaciones BOOLEAN DEFAULT false,
		notif_mensajes BOOLEAN DEFAULT false,
		notif_postulaciones BOOLEAN DEFAULT false,
		activo BOOLEAN DEFAULT true,
		creado_en TIMESTAMP DEFAULT NOW(),
		actualizado_en TIMESTAMP DEFAULT NOW()
	);`)

	m.SQL(`CREATE TABLE IF NOT EXISTS notificaciones (
		id_notificaciones SERIAL PRIMARY KEY,
		id_usuarios INT REFERENCES usuarios(id_usuarios),
		titulo VARCHAR(255),
		mensaje TEXT,
		leido BOOLEAN DEFAULT false,
		activo BOOLEAN DEFAULT true,
		creado_en TIMESTAMP DEFAULT NOW(),
		actualizado_en TIMESTAMP DEFAULT NOW()
	);`)
}

// Reverse the migrations
func (m *JobsyMigrations_20260521_125724) Down() {
	m.SQL(`DROP TABLE IF EXISTS configuraciones;`)
	m.SQL(`DROP TABLE IF EXISTS notificaciones;`)
	m.SQL(`DROP TABLE IF EXISTS usuarios;`)
}