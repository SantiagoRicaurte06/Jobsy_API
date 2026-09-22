-- Semilla de los catalogos maestros que no tenian ninguna.
--
-- Los scripts de semilla de Empleo, Tienda, Soporte, Suscripciones y Usuarios
-- resuelven sus claves foraneas con subconsultas del tipo:
--
--     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='pago' AND nombre='tienda')
--
-- Contra una tabla vacia eso devuelve NULL y el INSERT falla por NOT NULL.
--
-- QUE CUBRE ESTE ARCHIVO
--   catalogo.tipos                       (no se sembraba en ninguna parte)
--   catalogo.generos                     (idem)
--   catalogo.niveles_experiencia         (idem)
--   catalogo.modalidades_trabajo         (idem)
--   catalogo.turnos                      (idem)
--   catalogo.dias_semana                 (idem)
--   catalogo.modalidades_suscripcion     (idem)
--
-- QUE NO CUBRE, A PROPOSITO
--   catalogo.estados se siembra en la migracion
--   database/migrations/20260520_221642_estados.go, con 39 filas que incluyen
--   los 18 estados que el resto de servicios necesita. Duplicarlas aqui
--   crearia dos fuentes de verdad que pueden divergir.
--
--   Ojo: esa migracion no compilaba (le faltaba la comilla invertida que
--   cierra el literal de Up), asi que sus INSERT nunca se habian ejecutado.
--   Se corrigio junto con este archivo.
--
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
