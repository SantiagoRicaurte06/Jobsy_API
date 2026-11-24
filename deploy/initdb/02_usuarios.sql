-- =====================================================================
--  Esquema USUARIOS (lo sirve Usuarios_Api)
--
--  jobsy_supabase.sql no lo trae. Esto es copia del Up() de
--  Usuarios_Api/database/migrations/20260521_131549_jobsy_migrations.go en la
--  rama fix-ddl-desplegable (la de main no se puede ejecutar: tiene un ");"
--  de mas y referencia una tabla usuarios que no existe).
-- =====================================================================

CREATE SCHEMA IF NOT EXISTS usuarios;

CREATE TABLE IF NOT EXISTS usuarios.configuraciones (
    id_configuraciones SERIAL PRIMARY KEY,
    id_usuarios INT NOT NULL REFERENCES core.usuarios(id_usuarios),
    perfil_publico BOOLEAN NOT NULL DEFAULT false,
    mostrar_correo BOOLEAN NOT NULL DEFAULT false,
    mostrar_telefono BOOLEAN NOT NULL DEFAULT false,
    hv_visible BOOLEAN NOT NULL DEFAULT false,
    aparecer_busquedas BOOLEAN NOT NULL DEFAULT false,
    mostrar_ubicacion BOOLEAN NOT NULL DEFAULT false,
    doble_factor BOOLEAN NOT NULL DEFAULT false,
    notif_correo BOOLEAN NOT NULL DEFAULT false,
    notif_plataforma BOOLEAN NOT NULL DEFAULT false,
    notif_nuevas_ofertas BOOLEAN NOT NULL DEFAULT false,
    notif_actualizaciones BOOLEAN NOT NULL DEFAULT false,
    notif_mensajes BOOLEAN NOT NULL DEFAULT false,
    notif_postulaciones BOOLEAN NOT NULL DEFAULT false,
    activo BOOLEAN NOT NULL DEFAULT true,
    creado_en TIMESTAMP NOT NULL DEFAULT NOW(),
    actualizado_en TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Una fila de configuracion por usuario: con dos, no hay forma de saber
-- cual manda.
CREATE UNIQUE INDEX IF NOT EXISTS ux_configuraciones_usuario
    ON usuarios.configuraciones (id_usuarios);

CREATE TABLE IF NOT EXISTS usuarios.notificaciones (
    id_notificaciones SERIAL PRIMARY KEY,
    id_usuarios INT NOT NULL REFERENCES core.usuarios(id_usuarios),
    id_tipo INT NOT NULL REFERENCES catalogo.tipos(id_tipo),
    titulo VARCHAR(255) NOT NULL,
    cuerpo TEXT,
    leida BOOLEAN NOT NULL DEFAULT false,
    url_destino VARCHAR(300),
    activo BOOLEAN NOT NULL DEFAULT true,
    creado_en TIMESTAMP NOT NULL DEFAULT NOW(),
    actualizado_en TIMESTAMP NOT NULL DEFAULT NOW()
);

-- La bandeja se consulta siempre por usuario y por no leidas.
CREATE INDEX IF NOT EXISTS ix_notificaciones_usuario
    ON usuarios.notificaciones (id_usuarios);
