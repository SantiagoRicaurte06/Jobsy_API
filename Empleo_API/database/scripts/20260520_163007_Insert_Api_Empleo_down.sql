-- ============================================================
-- DELETE seed data — esquema empleo
-- Orden inverso a los INSERTs para respetar FKs
-- ============================================================

DELETE FROM empleo.horarios_agendados
    WHERE hora_inicio IN ('07:30', '08:00', '07:00')
      AND id_contrataciones IN (1, 2, 3);

DELETE FROM empleo.TR_contratacion_especialidades
    WHERE id_contrataciones IN (1, 2, 3);

DELETE FROM empleo.contrataciones
    WHERE direccion_servicio IN (
        'Cra 76 # 35-12 Apto 401, Laureles, Medellín',
        'Cl 37 Sur # 43-60, El Trianón, Envigado',
        'Cra 84 # 65-22, Robledo, Medellín'
    );

DELETE FROM empleo.resenas
    WHERE comentario IN (
        'Excelente trabajo, muy puntual y responsable.',
        'Buen trabajo de limpieza, dejó todo muy ordenado.',
        'El empleador fue muy amable y el espacio era cómodo.'
    );

DELETE FROM empleo.postulacion_chips
    WHERE texto IN ('Con referencias', 'Disponible inmediato', 'Productos propios');

DELETE FROM empleo.TR_postulacion_dias
    WHERE id_postulaciones IN (1, 2, 3);

DELETE FROM empleo.postulaciones
    WHERE mensaje_empleador IN (
        'Tengo experiencia en apartamentos similares, soy muy organizada.',
        'Especializada en limpieza profunda, traigo mis propios productos.',
        'Tengo experiencia cuidando niños y puedo empezar de inmediato.'
    );

DELETE FROM empleo.oferta_fotos
    WHERE url_foto IN (
        'https://cdn.jobsy.co/ofertas/1_sala.jpg',
        'https://cdn.jobsy.co/ofertas/1_cocina.jpg',
        'https://cdn.jobsy.co/ofertas/2_fachada.jpg'
    );

DELETE FROM empleo.TR_oferta_servicios
    WHERE id_ofertas IN (1, 2, 3);

DELETE FROM empleo.ofertas
    WHERE titulo IN (
        'Aseo apartamento Laureles',
        'Limpieza profunda casa Envigado',
        'Asistente de hogar tiempo completo'
    );

DELETE FROM empleo.hv_experiencias
    WHERE empresa IN ('Familia Rodríguez', 'Familia Ospina', 'Familia Torres');

DELETE FROM empleo.hojas_vida
    WHERE nombre IN (
        'HV Principal Laura',
        'HV Principal Valentina',
        'HV Secundaria Laura'
    );

DELETE FROM empleo.documentos_usuario
    WHERE url_archivo IN (
        'https://cdn.jobsy.co/docs/laura_antec.pdf',
        'https://cdn.jobsy.co/docs/vale_antec.pdf',
        'https://cdn.jobsy.co/docs/laura_foto.jpg'
    );

DELETE FROM empleo.certificaciones
    WHERE nombre IN (
        'Manipulación de alimentos',
        'Primeros auxilios básicos',
        'Limpieza y desinfección hogar'
    );

DELETE FROM empleo.disponibilidad_empleado
    WHERE hora_inicio IN ('07:00', '13:00')
      AND id_perfiles_empleado IN (1, 2);

DELETE FROM empleo.TR_perfil_especialidades
    WHERE id_perfiles_empleado IN (1, 2);

DELETE FROM empleo.perfiles_empleado
    WHERE barrio_zona IN ('Chapinero', 'El Poblado');

DELETE FROM empleo.dias_trabajados
    WHERE id_dia_semana IN (1, 2, 5);

DELETE FROM empleo.tamano_inmueble
    WHERE nombre IN ('pequeno', 'mediano', 'grande');

DELETE FROM empleo.tipo_inmueble
    WHERE nombre IN ('casa', 'apartamento', 'oficina_local');

DELETE FROM empleo.especialidades
    WHERE nombre IN (
        'Limpieza general',
        'Cocina y alimentación',
        'Cuidado de niños'
    );