package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Estados_20260520_221642 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Estados_20260520_221642{}
	m.Created = "20260520_221642"

	migration.Register("Estados_20260520_221642", m)
}

// Run the migrations
func (m *Estados_20260520_221642) Up() {

	m.SQL(`
		INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
		    ('documento', 'pendiente',  'Documento recibido, pendiente de revisión por el equipo'),
		    ('documento', 'verificado', 'Documento revisado y aprobado correctamente'),
		    ('documento', 'rechazado',  'Documento rechazado por inconsistencias o ilegibilidad');

		INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
		    ('certificacion', 'pendiente',  'Certificación enviada, en espera de validación'),
		    ('certificacion', 'validado',   'Certificación verificada y marcada como válida'),
		    ('certificacion', 'rechazado',  'Certificación no aprobada por la plataforma');

		INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
		    ('oferta', 'activa',   'Oferta publicada y visible para los empleados'),
		    ('oferta', 'pausada',  'Oferta temporalmente ocultada por el empleador'),
		    ('oferta', 'cerrada',  'Oferta cerrada: ya no acepta postulaciones'),
		    ('oferta', 'borrada',  'Oferta eliminada por el empleador o administrador');

		INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
		    ('postulacion', 'pendiente',  'Postulación enviada y en espera de revisión'),
		    ('postulacion', 'revisando',  'El empleador está evaluando activamente la postulación'),
		    ('postulacion', 'aceptado',   'Postulación aprobada, se procede a contratación'),
		    ('postulacion', 'rechazado',  'Postulación no fue aceptada por el empleador'),
		    ('postulacion', 'cancelado',  'El empleado retiró su postulación antes de recibir respuesta');

		INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
		    ('contratacion', 'activa',     'Contratación en curso, servicio siendo prestado'),
		    ('contratacion', 'completada', 'Servicio finalizado correctamente por ambas partes'),
		    ('contratacion', 'cancelada',  'Contratación cancelada antes o durante la prestación');

		INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
		    ('carrito', 'activo',     'Carrito con productos seleccionados pendientes de pago'),
		    ('carrito', 'pagado',     'Carrito procesado y convertido en pedido pagado'),
		    ('carrito', 'abandonado', 'Carrito sin actividad por más de 24 horas');

		INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
		    ('pedido', 'pendiente',   'Pedido generado, en espera de confirmación de pago'),
		    ('pedido', 'pagado',      'Pago confirmado, pedido listo para despacho'),
		    ('pedido', 'enviado',     'Pedido entregado al operador logístico'),
		    ('pedido', 'entregado',   'Pedido recibido satisfactoriamente por el cliente'),
		    ('pedido', 'cancelado',   'Pedido cancelado antes del despacho'),
		    ('pedido', 'reembolsado', 'Pago devuelto al cliente por devolución o inconveniente');

		INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
		    ('pago', 'pendiente',   'Pago iniciado, esperando respuesta de la pasarela'),
		    ('pago', 'procesando',  'Pago en proceso de verificación bancaria'),
		    ('pago', 'aprobado',    'Pago aprobado y fondos debitados exitosamente'),
		    ('pago', 'rechazado',   'Pago rechazado por la entidad financiera'),
		    ('pago', 'reembolsado', 'Monto devuelto al usuario por error o cancelación');

		INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
		    ('suscripcion', 'activa',    'Suscripción vigente con todos los beneficios del plan activos'),
		    ('suscripcion', 'cancelada', 'Suscripción cancelada por el usuario antes de su vencimiento'),
		    ('suscripcion', 'expirada',  'Suscripción vencida sin renovación automática');

		INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
		    ('reporte', 'abierto',     'Reporte registrado, pendiente de atención por soporte'),
		    ('reporte', 'en_revision', 'Un agente de soporte está analizando el caso'),
		    ('reporte', 'respondido',  'El equipo de soporte envió respuesta al usuario'),
		    ('reporte', 'cerrado',     'Caso resuelto y cerrado por soporte o el usuario');)
}


func (m *Estados_20260520_221642) Down() {

	m.SQL(` DELETE FROM catalogo.estados WHERE contexto IN (
			'documento',
			'certificacion',
			'oferta',
			'postulacion',
			'contratacion',
			'carrito',
			'pedido',
			'pago',
			'suscripcion',
			'reporte'
		);`)
}
