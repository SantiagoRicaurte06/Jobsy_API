CREATE SCHEMA IF NOT EXISTS suscripciones;

CREATE TABLE suscripciones.planes (
    id_planes           SERIAL          PRIMARY KEY,
    nombre              VARCHAR(50)     NOT NULL UNIQUE,
    precio_mensual      DECIMAL(10,2)   NOT NULL,
    precio_anual        DECIMAL(10,2),
    descripcion         TEXT,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT NOW(),
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT NOW()
);

CREATE TABLE suscripciones.plan_features (
    id_plan_features    SERIAL          PRIMARY KEY,
    id_planes           INT             NOT NULL,
    feature             VARCHAR(150)    NOT NULL,
    incluida            BOOLEAN         NOT NULL DEFAULT TRUE,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT NOW(),
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_pf_planes FOREIGN KEY (id_planes) REFERENCES suscripciones.planes(id_planes)
);

CREATE TABLE suscripciones.usuarios_suscripciones (
    id_usuarios_suscripciones   SERIAL      PRIMARY KEY,
    id_usuarios                 INT         NOT NULL,
    id_planes                   INT         NOT NULL,
    id_modalidad_suscripcion    INT         NOT NULL,        -- FK → catalogo.modalidades_suscripcion
    id_estado                   INT         NOT NULL,        -- FK → catalogo.estados (contexto='suscripcion')
    inicio_en                   TIMESTAMP   NOT NULL,
    fin_en                      TIMESTAMP,
    auto_renovar                BOOLEAN     NOT NULL DEFAULT TRUE,
    activo                      BOOLEAN     NOT NULL DEFAULT TRUE,
    creado_en                   TIMESTAMP   NOT NULL DEFAULT NOW(),
    actualizado_en              TIMESTAMP   NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_us_usuarios  FOREIGN KEY (id_usuarios)             REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_us_planes    FOREIGN KEY (id_planes)               REFERENCES suscripciones.planes(id_planes),
    CONSTRAINT fk_us_modalidad FOREIGN KEY (id_modalidad_suscripcion) REFERENCES catalogo.modalidades_suscripcion(id_modalidad_suscripcion),
    CONSTRAINT fk_us_estado    FOREIGN KEY (id_estado)               REFERENCES catalogo.estados(id_estado)
);
