package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertsSchemaUsuario_20260521_141000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsSchemaUsuario_20260521_141000{}
	m.Created = "20260521_141000"

	migration.Register("InsertsSchemaUsuario_20260521_141000", m)
}

// Run the migrations
func (m *InsertsSchemaUsuario_20260521_141000) Up() {
	m.SQL(`INSERT INTO usuarios.configuraciones 
    (id_usuarios, perfil_publico, mostrar_correo, mostrar_telefono, hv_visible, 
    aparecer_busquedas, mostrar_ubicacion, doble_factor, notif_correo, notif_plataforma, 
    notif_nuevas_ofertas, notif_actualizaciones, notif_mensajes, notif_postulaciones, activo) 
    VALUES (1, false, false, false, false, false, false, false, false, false, false, false, false, false, true);`)

	// Las columnas son cuerpo y leida (no mensaje y leido): estos inserts
	// seguian los nombres del DDL antiguo, que no coincidian con el modelo Go.
	// id_tipo se resuelve por nombre para no depender de un id concreto.
	m.SQL(`INSERT INTO usuarios.notificaciones
    (id_usuarios, id_tipo, titulo, cuerpo, leida, activo)
    VALUES (1,
        (SELECT id_tipo FROM catalogo.tipos WHERE contexto='notificacion' AND nombre='sistema'),
        'Bienvenido a Jobsy', 'Tu cuenta ha sido creada exitosamente.', false, true);`)

}

// Reverse the migrations
func (m *InsertsSchemaUsuario_20260521_141000) Down() {
	m.SQL(`DELETE FROM usuarios.configuraciones WHERE id_configuraciones = 1;`)
	m.SQL(`DELETE FROM usuarios.notificaciones WHERE id_notificaciones = 1;`)
}
