package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertsSchemaSuscripciones_20260521_162957 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsSchemaSuscripciones_20260521_162957{}
	m.Created = "20260521_162957"

	migration.Register("InsertsSchemaSuscripciones_20260521_162957", m)
}

// Run the migrations
func (m *InsertsSchemaSuscripciones_20260521_162957) Up() {
	m.SQL(`
		INSERT INTO suscripciones.planes
			(nombre, precio_mensual, precio_anual, descripcion)
		VALUES
		    ('Básico',      0.00,      0.00,      'Plan gratuito con funcionalidades esenciales.'),
		    ('Pro',         29900.00,  299000.00, 'Plan profesional con acceso completo.'),
		    ('Empresarial', 89900.00,  899000.00, 'Plan para empresas con múltiples usuarios.');
	`)

	m.SQL(`
		INSERT INTO suscripciones.plan_features
		    (id_planes, feature, incluida)
		VALUES
		    (1, 'Postulación a 3 ofertas por mes', TRUE),
		    (1, 'Perfil público básico',            TRUE),
		    (2, 'Postulaciones ilimitadas',         TRUE),
		    (2, 'Verificación de documentos',       TRUE),
		    (3, 'Panel de administración',          TRUE),
		    (3, 'Soporte prioritario 24/7',         TRUE);
	`)

	m.SQL(`
		INSERT INTO suscripciones.usuarios_suscripciones
		    (id_usuarios, id_planes, id_modalidad_suscripcion, id_estado, inicio_en, fin_en, auto_renovar)
		VALUES
		    (1,
		     (SELECT id_planes FROM suscripciones.planes WHERE nombre='Pro'),
		     (SELECT id_modalidad_suscripcion FROM catalogo.modalidades_suscripcion WHERE nombre='mensual'),
		     (SELECT id_estado FROM catalogo.estados WHERE contexto='suscripcion' AND nombre='activa'),
		     NOW(), NOW() + INTERVAL '30 days', TRUE),

		    (2,
		     (SELECT id_planes FROM suscripciones.planes WHERE nombre='Básico'),
		     (SELECT id_modalidad_suscripcion FROM catalogo.modalidades_suscripcion WHERE nombre='mensual'),
		     (SELECT id_estado FROM catalogo.estados WHERE contexto='suscripcion' AND nombre='activa'),
		     NOW(), NULL, FALSE),

		    (3,
		     (SELECT id_planes FROM suscripciones.planes WHERE nombre='Empresarial'),
		     (SELECT id_modalidad_suscripcion FROM catalogo.modalidades_suscripcion WHERE nombre='anual'),
		     (SELECT id_estado FROM catalogo.estados WHERE contexto='suscripcion' AND nombre='activa'),
		     NOW(), NOW() + INTERVAL '365 days', TRUE);
	`)
}

// Reverse the migrations
func (m *InsertsSchemaSuscripciones_20260521_162957) Down() {
	m.SQL(`
		DELETE FROM suscripciones.usuarios_suscripciones
		    WHERE id_usuarios IN (1, 2, 3);
	`)

	m.SQL(`
		DELETE FROM suscripciones.plan_features
		    WHERE feature IN (
		        'Postulación a 3 ofertas por mes',
		        'Perfil público básico',
		        'Postulaciones ilimitadas',
		        'Verificación de documentos',
		        'Panel de administración',
		        'Soporte prioritario 24/7'
		    );
	`)

	m.SQL(`
		DELETE FROM suscripciones.planes
		    WHERE nombre IN ('Básico', 'Pro', 'Empresarial');
	`)
}
