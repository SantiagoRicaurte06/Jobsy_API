-- Deshace la semilla de los catalogos maestros.
--
-- No se borran filas sueltas por id: al no ser SERIAL fijos, un DELETE por id
-- concreto es fragil. Se vacian las tablas enteras, que es lo que corresponde
-- a deshacer una semilla.
--
-- Cualquier otra tabla que apunte a estos catalogos debe borrarse antes.

-- catalogo.estados no se toca: lo siembra otra migracion
-- (20260520_221642_estados.go) y le corresponde a ella deshacerlo.
DELETE FROM catalogo.tipos;
DELETE FROM catalogo.modalidades_suscripcion;
DELETE FROM catalogo.dias_semana;
DELETE FROM catalogo.turnos;
DELETE FROM catalogo.modalidades_trabajo;
DELETE FROM catalogo.niveles_experiencia;
DELETE FROM catalogo.generos;
