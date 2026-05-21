INSERT INTO soporte.reportes
    (id_usuarios, id_tipo, descripcion, id_estado)
VALUES
    (1,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='reporte' AND nombre='tecnico'),
     'El botón de postulación no responde en móvil Chrome.',
     (SELECT id_estado FROM catalogo.estados WHERE contexto='reporte' AND nombre='abierto')),
    (2,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='reporte' AND nombre='pago'),
     'Se me realizó un cobro duplicado por la suscripción Pro.',
     (SELECT id_estado FROM catalogo.estados WHERE contexto='reporte' AND nombre='en_revision')),
    (3,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='reporte' AND nombre='usuario'),
     'Un usuario me envió mensajes ofensivos en el chat.',
     (SELECT id_estado FROM catalogo.estados WHERE contexto='reporte' AND nombre='abierto'));
 
-- ---- soporte.reporte_mensajes --------------------------------
INSERT INTO soporte.reporte_mensajes
    (id_reportes, id_remitente, mensaje)
VALUES
    (1, 1, 'El problema ocurre desde ayer, lo probé en iOS y Android también.'),
    (2, 2, 'Adjunto captura del extracto bancario con el doble cargo.'),
    (2, 2, 'Ya hablé con mi banco y me confirmaron que salieron dos débitos.');