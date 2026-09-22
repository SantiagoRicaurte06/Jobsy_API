-- Semilla de los catalogos maestros.
--
-- Este archivo no existia. Los scripts de semilla de Empleo, Tienda, Soporte,
-- Suscripciones y Usuarios resuelven sus claves foraneas con subconsultas del
-- tipo:
--
--     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='pago' AND nombre='tienda')
--
-- Como catalogo.tipos y catalogo.estados estaban vacias, esas subconsultas
-- devolvian NULL y el INSERT fallaba contra la restriccion NOT NULL. Es decir:
-- ningun servicio podia sembrar datos.
--
-- Las filas de abajo son exactamente las que referencian los demas scripts.
-- Si anades una subconsulta nueva a un seed, anade aqui la fila que necesita.
--
-- ORDEN DE EJECUCION: despues de crear el esquema catalogo y antes de
-- cualquier otra semilla.

-- ---- catalogo.generos ------------------------------------------
-- Referenciado por core.usuarios.id_genero.
INSERT INTO catalogo.generos (nombre) VALUES
    ('Masculino'), ('Femenino'), ('Otro')
ON CONFLICT (nombre) DO NOTHING;

-- ---- catalogo.niveles_experiencia ------------------------------
-- Referenciado por empleo.perfiles_empleado.id_nivel_experiencia.
INSERT INTO catalogo.niveles_experiencia (nombre, descripcion) VALUES
    ('sin_experiencia', 'Sin experiencia previa comprobable'),
    ('intermedio',      'Entre uno y tres anos de experiencia'),
    ('experto',         'Mas de tres anos de experiencia')
ON CONFLICT (nombre) DO NOTHING;

-- ---- catalogo.modalidades_trabajo ------------------------------
-- OJO: el frontend solo modela cuatro modalidades y con otros nombres
-- (por_horas, jornada_completa, fin_de_semana, presencial). El middleware
-- traduce entre ambos vocabularios; 'presencial' no tiene equivalente aqui
-- porque es una modalidad de lugar, no de jornada.
INSERT INTO catalogo.modalidades_trabajo (nombre, descripcion) VALUES
    ('por_horas',        'Servicio cobrado por hora trabajada'),
    ('tiempo_completo',  'Jornada laboral completa'),
    ('por_jornadas',     'Jornadas sueltas acordadas con el empleador'),
    ('fines_semana',     'Solo sabados y domingos'),
    ('interno',          'El empleado reside en el lugar de trabajo'),
    ('dias_especificos', 'Dias concretos pactados de antemano')
ON CONFLICT (nombre) DO NOTHING;

-- ---- catalogo.turnos -------------------------------------------
INSERT INTO catalogo.turnos (nombre, descripcion) VALUES
    ('manana',   'Desde la manana hasta el mediodia'),
    ('tarde',    'Desde el mediodia hasta la noche'),
    ('noche',    'Horario nocturno'),
    ('completo', 'Todo el dia'),
    ('jornada',  'Jornada pactada sin franja fija')
ON CONFLICT (nombre) DO NOTHING;

-- ---- catalogo.dias_semana --------------------------------------
INSERT INTO catalogo.dias_semana (nombre, nombre_completo, orden) VALUES
    ('lun', 'Lunes',     1),
    ('mar', 'Martes',    2),
    ('mie', 'Miercoles', 3),
    ('jue', 'Jueves',    4),
    ('vie', 'Viernes',   5),
    ('sab', 'Sabado',    6),
    ('dom', 'Domingo',   7)
ON CONFLICT (nombre) DO NOTHING;

-- ---- catalogo.modalidades_suscripcion --------------------------
INSERT INTO catalogo.modalidades_suscripcion (nombre, descripcion) VALUES
    ('mensual', 'Cobro cada mes'),
    ('anual',   'Cobro cada ano')
ON CONFLICT (nombre) DO NOTHING;

-- ---- catalogo.estados ------------------------------------------
-- Tabla polimorfica: la unicidad es (contexto, nombre).
INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
    ('carrito',       'activo',      'Carrito en uso'),
    ('carrito',       'pagado',      'Carrito convertido en pedido'),
    ('carrito',       'abandonado',  'Carrito sin actividad'),

    ('pedido',        'pendiente',   'Pedido creado, pago no confirmado'),
    ('pedido',        'pagado',      'Pago confirmado'),
    ('pedido',        'enviado',     'Pedido despachado'),
    ('pedido',        'entregado',   'Pedido recibido por el cliente'),
    ('pedido',        'cancelado',   'Pedido anulado'),

    ('pago',          'pendiente',   'Pago iniciado, sin confirmar'),
    ('pago',          'aprobado',    'Pago confirmado'),
    ('pago',          'rechazado',   'Pago rechazado'),

    ('oferta',        'activa',      'Oferta visible y admitiendo postulaciones'),
    ('oferta',        'cerrada',     'Oferta que ya no admite postulaciones'),
    ('oferta',        'expirada',    'Oferta pasada de su fecha de expiracion'),

    ('postulacion',   'pendiente',   'Postulacion sin revisar'),
    ('postulacion',   'revisando',   'Postulacion en revision por el empleador'),
    ('postulacion',   'aceptada',    'Postulacion aceptada'),
    ('postulacion',   'rechazada',   'Postulacion rechazada'),

    ('contratacion',  'activa',      'Contratacion en curso'),
    ('contratacion',  'completada',  'Servicio finalizado'),
    ('contratacion',  'cancelada',   'Contratacion anulada'),

    ('certificacion', 'pendiente',   'Certificacion sin validar'),
    ('certificacion', 'validado',    'Certificacion verificada'),
    ('certificacion', 'rechazado',   'Certificacion no valida'),

    ('documento',     'pendiente',   'Documento sin verificar'),
    ('documento',     'verificado',  'Documento verificado'),
    ('documento',     'rechazado',   'Documento no valido'),

    ('reporte',       'abierto',     'Reporte recien creado'),
    ('reporte',       'en_revision', 'Reporte siendo atendido'),
    ('reporte',       'resuelto',    'Reporte resuelto'),
    ('reporte',       'cerrado',     'Reporte cerrado'),

    ('suscripcion',   'activa',      'Suscripcion vigente'),
    ('suscripcion',   'vencida',     'Suscripcion caducada'),
    ('suscripcion',   'cancelada',   'Suscripcion cancelada por el usuario')
ON CONFLICT (contexto, nombre) DO NOTHING;

-- ---- catalogo.tipos --------------------------------------------
INSERT INTO catalogo.tipos (contexto, nombre, descripcion) VALUES
    ('producto_tienda',   'insumos_aseo',          'Productos de limpieza'),
    ('producto_tienda',   'equipo',                'Equipos y herramientas'),
    ('producto_tienda',   'proteccion',            'Elementos de proteccion personal'),
    ('producto_tienda',   'uniforme',              'Uniformes y ropa de trabajo'),

    ('pago',              'tienda',                'Pago de un pedido de la tienda'),
    ('pago',              'contratacion',          'Pago de un servicio contratado'),
    ('pago',              'suscripcion',           'Pago de una suscripcion'),

    ('cuenta_bancaria',   'ahorros',               'Cuenta de ahorros'),
    ('cuenta_bancaria',   'corriente',             'Cuenta corriente'),

    ('documento_usuario', 'foto',                  'Fotografia del documento de identidad'),
    ('documento_usuario', 'antecedentes',          'Certificado de antecedentes'),
    ('documento_usuario', 'hoja_vida',             'Hoja de vida en PDF'),

    ('mascota',           'ninguna',               'Sin mascotas en la vivienda'),
    ('mascota',           'perros',                'Perros en la vivienda'),
    ('mascota',           'gatos',                 'Gatos en la vivienda'),
    ('mascota',           'varias',                'Varias especies'),

    ('experiencia_hv',    'hogar_residencial',     'Experiencia en hogares'),
    ('experiencia_hv',    'oficina_comercial',     'Experiencia en oficinas o locales'),

    ('forma_cobro',       'efectivo',              'Cobro en efectivo'),
    ('forma_cobro',       'transferencia',         'Cobro por transferencia'),
    ('forma_cobro',       'saldo_jobsy',           'Cobro al saldo de la plataforma'),

    ('resena',            'empleador_a_empleado',  'Resena escrita por el empleador'),
    ('resena',            'empleado_a_empleador',  'Resena escrita por el empleado'),

    ('reporte',           'tecnico',               'Problema tecnico de la plataforma'),
    ('reporte',           'usuario',               'Reporte sobre otro usuario'),
    ('reporte',           'pago',                  'Problema con un cobro o transferencia'),
    ('reporte',           'sugerencia',            'Sugerencia de mejora'),

    ('notificacion',      'sistema',               'Aviso generado por la plataforma'),
    ('notificacion',      'postulacion',           'Novedad sobre una postulacion'),
    ('notificacion',      'mensaje',               'Mensaje recibido'),
    ('notificacion',      'pago',                  'Novedad sobre un pago')
ON CONFLICT (contexto, nombre) DO NOTHING;
