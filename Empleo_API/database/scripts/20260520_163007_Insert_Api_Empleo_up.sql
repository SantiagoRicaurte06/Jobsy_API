-- ---- empleo.especialidades -----------------------------------
INSERT INTO empleo.especialidades (nombre, descripcion, icono) VALUES
    ('Limpieza general',     'Aseo y limpieza de espacios del hogar',      'icon-spray'),
    ('Cocina y alimentación','Preparación de alimentos y limpieza de cocina','icon-chef'),
    ('Cuidado de niños',     'Acompañamiento y cuidado de menores de edad', 'icon-baby');
 
-- ---- empleo.tipo_inmueble ------------------------------------
INSERT INTO empleo.tipo_inmueble (nombre, descripcion) VALUES
    ('casa',          'Casa unifamiliar independiente'),
    ('apartamento',   'Unidad residencial en conjunto o edificio'),
    ('oficina_local', 'Espacio de trabajo comercial o empresarial');
 
-- ---- empleo.tamano_inmueble ----------------------------------
INSERT INTO empleo.tamano_inmueble (nombre, descripcion) VALUES
    ('pequeno', 'Menos de 60 m² aproximadamente'),
    ('mediano', 'Entre 60 m² y 120 m²'),
    ('grande',  'Más de 120 m² o propiedades amplias');
 
-- ---- empleo.dias_trabajados ----------------------------------
-- id_dia_semana: 1=lun, 2=mar, 5=vie
INSERT INTO empleo.dias_trabajados (id_dia_semana, horas_trabajadas) VALUES
    (1, '8'),
    (2, '6'),
    (5, '4');
 
-- ---- empleo.perfiles_empleado --------------------------------
-- id_nivel: 2=intermedio 3=experto | id_modalidad: 1=por_horas | id_turno: 1=manana 2=tarde
INSERT INTO empleo.perfiles_empleado
    (id_usuarios, bio, id_nivel_experiencia, anios_experiencia, tarifa_hora,
     id_modalidad_trabajo, id_turno_preferido, distancia_max_km, barrio_zona)
VALUES
    (1, 'Soy responsable, puntual y con experiencia en hogares.', 2, 2, 18000.00, 1, 1, 10, 'Chapinero'),
    (3, 'Experta en limpieza y organización de apartamentos.',     3, 5, 22000.00, 1, 2, 5,  'El Poblado');
 
-- ---- empleo.TR_perfil_especialidades -------------------------
INSERT INTO empleo.TR_perfil_especialidades (id_perfiles_empleado, id_especialidades) VALUES
    (1, 1),
    (1, 2),
    (2, 1);
 
-- ---- empleo.disponibilidad_empleado --------------------------
-- id_dia: 1=lun 3=mie | id_turno: 1=manana 2=tarde
INSERT INTO empleo.disponibilidad_empleado
    (id_perfiles_empleado, id_dia_semana, id_turno, hora_inicio, hora_fin)
VALUES
    (1, 1, 1, '07:00', '12:00'),
    (1, 3, 1, '07:00', '12:00'),
    (2, 2, 2, '13:00', '18:00');
 
-- ---- empleo.certificaciones ----------------------------------
-- catalogo.estados: buscar contexto='certificacion'
-- pendiente→id dinámico; usamos subquery para no hardcodear ids
INSERT INTO empleo.certificaciones
    (id_perfiles_empleado, nombre, entidad_emisora, anio_emision, id_estado)
VALUES
    (1, 'Manipulación de alimentos',     'SENA',      2022,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='certificacion' AND nombre='validado')),
    (1, 'Primeros auxilios básicos',     'Cruz Roja', 2023,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='certificacion' AND nombre='pendiente')),
    (2, 'Limpieza y desinfección hogar', 'SENA',      2021,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='certificacion' AND nombre='validado'));
 
-- ---- empleo.documentos_usuario -------------------------------
INSERT INTO empleo.documentos_usuario
    (id_usuarios, id_tipo, url_archivo, id_estado)
VALUES
    (1,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='documento_usuario' AND nombre='antecedentes'),
     'https://cdn.jobsy.co/docs/laura_antec.pdf',
     (SELECT id_estado FROM catalogo.estados WHERE contexto='documento' AND nombre='verificado')),
    (3,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='documento_usuario' AND nombre='antecedentes'),
     'https://cdn.jobsy.co/docs/vale_antec.pdf',
     (SELECT id_estado FROM catalogo.estados WHERE contexto='documento' AND nombre='pendiente')),
    (1,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='documento_usuario' AND nombre='foto'),
     'https://cdn.jobsy.co/docs/laura_foto.jpg',
     (SELECT id_estado FROM catalogo.estados WHERE contexto='documento' AND nombre='verificado'));
 
-- ---- empleo.hojas_vida ---------------------------------------
INSERT INTO empleo.hojas_vida
    (id_usuarios, nombre, resumen, es_predeterminada, es_publica, progreso_pct)
VALUES
    (1, 'HV Principal Laura',    'Empleada con 2 años de experiencia en hogares bogotanos.', TRUE,  TRUE,  85),
    (3, 'HV Principal Valentina','Especialista en limpieza profunda y organización.',         TRUE,  TRUE,  90),
    (1, 'HV Secundaria Laura',   'Perfil enfocado en cuidado de niños y cocina.',             FALSE, FALSE, 60);
 
-- ---- empleo.hv_experiencias ----------------------------------
INSERT INTO empleo.hv_experiencias
    (id_hojas_vida, id_tipo, descripcion, empresa, cargo, fecha_inicio, fecha_fin)
VALUES
    (1,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='experiencia_hv' AND nombre='hogar_residencial'),
     'Aseo general y cocina en familia de 4 personas.', 'Familia Rodríguez', 'Empleada doméstica', '2022-01-10', '2023-06-30'),
    (2,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='experiencia_hv' AND nombre='hogar_residencial'),
     'Limpieza profunda y plancha semanal.',            'Familia Ospina',    'Empleada doméstica', '2019-03-01', '2022-12-31'),
    (1,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='experiencia_hv' AND nombre='hogar_residencial'),
     'Cuidado de dos niños en edad escolar.',           'Familia Torres',    'Niñera',             '2021-07-15', '2022-01-05');
 
-- ---- empleo.ofertas ------------------------------------------
-- id_empleador: 2=Carlos | id_tipo_inmueble: 2=apto | id_tamano: 2=mediano
-- id_modalidad: 1=por_horas | mascota: 1=ninguna 2=perros | estado oferta: activa
INSERT INTO empleo.ofertas
    (id_empleador, titulo, descripcion, id_tipo_inmueble, id_tamano_inmueble,
     num_habitaciones, num_banos, barrio, direccion, ciudad,
     tarifa_hora, id_modalidad_trabajo, insumos_cliente, id_tipo_mascota,
     id_estado, expira_en)
VALUES
    (2, 'Aseo apartamento Laureles',
        'Necesito aseo general los miércoles, incluye cocina y baños.',
        2, 2, 3, 2, 'Laureles', 'Cra 76 # 35-12 Apto 401', 'Medellín',
        20000.00, 1, TRUE,
        (SELECT id_tipo FROM catalogo.tipos WHERE contexto='mascota' AND nombre='ninguna'),
        (SELECT id_estado FROM catalogo.estados WHERE contexto='oferta' AND nombre='activa'),
        NOW() + INTERVAL '30 days'),
 
    (2, 'Limpieza profunda casa Envigado',
        'Limpieza profunda mensual en casa de dos pisos.',
        1, 3, 4, 3, 'El Trianón', 'Cl 37 Sur # 43-60', 'Envigado',
        18000.00, 3, FALSE,
        (SELECT id_tipo FROM catalogo.tipos WHERE contexto='mascota' AND nombre='perros'),
        (SELECT id_estado FROM catalogo.estados WHERE contexto='oferta' AND nombre='activa'),
        NOW() + INTERVAL '20 days'),
 
    (2, 'Asistente de hogar tiempo completo',
        'Busco empleada de tiempo completo para hogar con dos niños.',
        1, 2, 3, 2, 'Robledo', 'Cra 84 # 65-22', 'Medellín',
        1500000.00, 2, TRUE,
        (SELECT id_tipo FROM catalogo.tipos WHERE contexto='mascota' AND nombre='perros'),
        (SELECT id_estado FROM catalogo.estados WHERE contexto='oferta' AND nombre='activa'),
        NOW() + INTERVAL '45 days');
 
-- ---- empleo.TR_oferta_servicios ------------------------------
INSERT INTO empleo.TR_oferta_servicios (id_ofertas, id_especialidades) VALUES
    (1, 1),
    (2, 1),
    (3, 2);
 
-- ---- empleo.oferta_fotos -------------------------------------
INSERT INTO empleo.oferta_fotos (id_ofertas, url_foto, orden, es_principal) VALUES
    (1, 'https://cdn.jobsy.co/ofertas/1_sala.jpg',    1, TRUE),
    (1, 'https://cdn.jobsy.co/ofertas/1_cocina.jpg',  2, FALSE),
    (2, 'https://cdn.jobsy.co/ofertas/2_fachada.jpg', 1, TRUE);
 
-- ---- empleo.postulaciones ------------------------------------
-- id_estado postulacion: pendiente=1 revisando=2
-- id_forma_cobro: efectivo
INSERT INTO empleo.postulaciones
    (id_ofertas, id_empleado, id_hojas_vida, id_estado,
     disponible_desde, id_modalidad_disp, id_turno_postulacion,
     id_forma_cobro, tarifa_min, tarifa_max, mensaje_empleador, acepta_terminos)
VALUES
    (1, 1, 1,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='postulacion' AND nombre='pendiente'),
     '2025-05-01', 1, 1,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='forma_cobro' AND nombre='efectivo'),
     18000, 22000, 'Tengo experiencia en apartamentos similares, soy muy organizada.', TRUE),
 
    (2, 3, 2,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='postulacion' AND nombre='revisando'),
     '2025-05-10', 3, 4,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='forma_cobro' AND nombre='efectivo'),
     17000, 20000, 'Especializada en limpieza profunda, traigo mis propios productos.', TRUE),
 
    (3, 1, 3,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='postulacion' AND nombre='pendiente'),
     '2025-05-15', 2, 5,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='forma_cobro' AND nombre='efectivo'),
     1400000, 1600000, 'Tengo experiencia cuidando niños y puedo empezar de inmediato.', TRUE);
 
-- ---- empleo.TR_postulacion_dias ------------------------------
-- id_dia: 1=lun 3=mie
INSERT INTO empleo.TR_postulacion_dias (id_postulaciones, id_dia_semana) VALUES
    (1, 3),
    (2, 1),
    (3, 1);
 
-- ---- empleo.postulacion_chips --------------------------------
INSERT INTO empleo.postulacion_chips (id_postulaciones, texto) VALUES
    (1, 'Con referencias'),
    (1, 'Disponible inmediato'),
    (2, 'Productos propios');
 
-- ---- empleo.resenas ------------------------------------------
INSERT INTO empleo.resenas
    (id_postulaciones, id_reviewer, id_reviewed, calificacion, comentario, id_tipo)
VALUES
    (1, 2, 1, 4.8, 'Excelente trabajo, muy puntual y responsable.',
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='resena' AND nombre='empleador_a_empleado')),
    (2, 2, 3, 4.5, 'Buen trabajo de limpieza, dejó todo muy ordenado.',
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='resena' AND nombre='empleador_a_empleado')),
    (1, 1, 2, 5.0, 'El empleador fue muy amable y el espacio era cómodo.',
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='resena' AND nombre='empleado_a_empleador'));
 
-- ---- empleo.contrataciones -----------------------------------
INSERT INTO empleo.contrataciones
    (id_empleador, id_empleado, id_ofertas, id_postulaciones,
     id_estado, direccion_servicio, fecha_inicio, hora_inicio,
     duracion_horas, id_modalidad, tarifa_acordada, total_estimado)
VALUES
    (2, 1, 1, 1,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='contratacion' AND nombre='activa'),
     'Cra 76 # 35-12 Apto 401, Laureles, Medellín',
     '2025-05-07', '07:30', 4.0, 1, 20000.00, 80000.00),
 
    (2, 3, 2, 2,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='contratacion' AND nombre='completada'),
     'Cl 37 Sur # 43-60, El Trianón, Envigado',
     '2025-04-20', '08:00', 6.0, 3, 18000.00, 108000.00),
 
    (2, 1, 3, 3,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='contratacion' AND nombre='activa'),
     'Cra 84 # 65-22, Robledo, Medellín',
     '2025-05-20', '07:00', 8.0, 2, 187500.00, 1500000.00);
 
-- ---- empleo.TR_contratacion_especialidades -------------------
INSERT INTO empleo.TR_contratacion_especialidades (id_contrataciones, id_especialidades) VALUES
    (1, 1),
    (2, 1),
    (3, 2);
 
-- ---- empleo.horarios_agendados ------------------------------
-- id_dia: 3=mie 1=lun | id_turno: 1=manana 4=completo 5=jornada
INSERT INTO empleo.horarios_agendados
    (id_contrataciones, id_dia_semana, id_turno, hora_inicio, hora_fin)
VALUES
    (1, 3, 1, '07:30', '11:30'),
    (2, 1, 4, '08:00', '14:00'),
    (3, 1, 5, '07:00', '15:00');