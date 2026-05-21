-- ============================================================
-- MÓDULO SOPORTE
-- ============================================================

CREATE SCHEMA IF NOT EXISTS soporte;

CREATE TABLE soporte.reportes (
    id_reportes         SERIAL          PRIMARY KEY,
    id_usuarios         INT             NOT NULL,
    id_tipo             INT             NOT NULL,            -- FK → catalogo.tipos (contexto='reporte')
    descripcion         TEXT            NOT NULL,
    archivo_url         VARCHAR(300),
    id_estado           INT             NOT NULL,            -- FK → catalogo.estados (contexto='reporte')
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT NOW(),
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_rep_usuarios FOREIGN KEY (id_usuarios) REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_rep_tipo     FOREIGN KEY (id_tipo)     REFERENCES catalogo.tipos(id_tipo),
    CONSTRAINT fk_rep_estado   FOREIGN KEY (id_estado)   REFERENCES catalogo.estados(id_estado)
);

CREATE TABLE soporte.reporte_mensajes (
    id_reporte_mensajes SERIAL          PRIMARY KEY,
    id_reportes         INT             NOT NULL,
    id_remitente        INT             NOT NULL,
    mensaje             TEXT            NOT NULL,
    archivo_url         VARCHAR(300),
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT NOW(),
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_rm_reportes  FOREIGN KEY (id_reportes)  REFERENCES soporte.reportes(id_reportes),
    CONSTRAINT fk_rm_remitente FOREIGN KEY (id_remitente) REFERENCES core.usuarios(id_usuarios)
);
