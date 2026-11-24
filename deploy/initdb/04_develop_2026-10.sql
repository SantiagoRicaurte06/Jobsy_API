-- Cambios de esquema que develop de Jobsy_API trae despues del dump del 8 de
-- octubre. Son las migraciones Up() de:
--   Tienda_APi/database/migrations/20261022_220000_datos_envio_pedido.go
--   Empleo_API/database/migrations/20261022_000000_hv_educacion_habilidades.go
-- y las especialidades nuevas de Empleo_API/database/scripts/..._Insert_Api_Empleo_up.sql.
-- Todo es idempotente: se puede volver a correr sobre una base ya actualizada.

ALTER TABLE tienda.facturacion
    ADD COLUMN IF NOT EXISTS factura_barrio VARCHAR(100),
    ADD COLUMN IF NOT EXISTS factura_detalle VARCHAR(150),
    ADD COLUMN IF NOT EXISTS factura_notas VARCHAR(300);

ALTER TABLE empleo.hojas_vida
    ADD COLUMN IF NOT EXISTS educacion TEXT,
    ADD COLUMN IF NOT EXISTS habilidades TEXT;

INSERT INTO empleo.especialidades (nombre, descripcion, icono)
SELECT v.nombre, v.descripcion, v.icono
FROM (VALUES
    ('Limpieza profunda',  'Limpieza a fondo de cocina, baños y zonas difíciles', 'icon-sparkles'),
    ('Planchado',          'Planchado y organización de ropa',                     'icon-shirt'),
    ('Zonas verdes',       'Cuidado de jardines y zonas verdes',                   'icon-leaf'),
    ('Cuidado de adultos', 'Acompañamiento de adultos mayores',                    'icon-heart'),
    ('Mascotas',           'Paseo y cuidado de mascotas',                          'icon-paw'),
    ('Bioseguridad',       'Desinfección y protocolos de bioseguridad',            'icon-shield')
) AS v(nombre, descripcion, icono)
WHERE NOT EXISTS (SELECT 1 FROM empleo.especialidades e WHERE e.nombre = v.nombre);
