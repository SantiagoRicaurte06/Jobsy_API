-- Deshace la creacion del esquema core.
--
-- Antes apuntaba al esquema "core_api", que no existe: el up crea "core". Con
-- IF EXISTS eso no daba error, simplemente no borraba nada.
--
-- El orden importa: primero las tablas que apuntan a otras, luego las
-- apuntadas. Al reves, las claves foraneas lo impiden.

DROP TABLE IF EXISTS core.sesiones;
DROP TABLE IF EXISTS core.credenciales;
DROP TABLE IF EXISTS core.documentos_identidad;
DROP TABLE IF EXISTS core.usuarios;
DROP TABLE IF EXISTS core.tipo_documento;
DROP TABLE IF EXISTS core.roles_usuarios;

DROP SCHEMA IF EXISTS core;
