-- =========================================
-- ROLES USUARIOS
-- =========================================

CREATE TABLE tienda.roles_usuarios (
    id_roles_usuarios SERIAL PRIMARY KEY,
    nombre            VARCHAR(50) NOT NULL UNIQUE,
    descripcion       VARCHAR(200),
    activo            BOOLEAN NOT NULL DEFAULT TRUE,
    creado_en         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =========================================
-- USUARIOS
-- =========================================

CREATE TABLE tienda.usuarios (
    id_usuarios         SERIAL PRIMARY KEY,
    nombre              VARCHAR(80) NOT NULL,
    apellido            VARCHAR(80) NOT NULL,
    correo              VARCHAR(150) NOT NULL UNIQUE,
    telefono            VARCHAR(20) NOT NULL,
    ciudad              VARCHAR(80) NOT NULL,
    fecha_nacimiento    DATE NOT NULL,
    id_genero           INT NOT NULL,
    nacionalidad        VARCHAR(60) NOT NULL,
    foto_url            VARCHAR(300) NOT NULL,
    id_roles_usuarios   INT NOT NULL,
    disponible          BOOLEAN NOT NULL DEFAULT TRUE,
    perfil_publico      BOOLEAN NOT NULL DEFAULT TRUE,
    activo              BOOLEAN NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_usuarios_roles
        FOREIGN KEY (id_roles_usuarios)
        REFERENCES tienda.roles_usuarios(id_roles_usuarios)
);

-- =========================================
-- TIPO DOCUMENTO
-- =========================================

CREATE TABLE tienda.tipo_documento (
    id_tipo_documento SERIAL PRIMARY KEY,
    nombre            VARCHAR(50) NOT NULL UNIQUE,
    descripcion       VARCHAR(200),
    activo            BOOLEAN NOT NULL DEFAULT TRUE,
    creado_en         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =========================================
-- DOCUMENTOS IDENTIDAD
-- =========================================

CREATE TABLE tienda.documentos_identidad (
    id_documentos_identidad SERIAL PRIMARY KEY,
    id_usuarios             INT NOT NULL,
    id_tipo_documento       INT NOT NULL,
    numero_documento        VARCHAR(30) NOT NULL UNIQUE,
    fecha_expedicion        DATE,
    lugar_expedicion        VARCHAR(100),
    activo                  BOOLEAN NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT uq_usuario_tipo_doc
        UNIQUE (id_usuarios, id_tipo_documento),

    CONSTRAINT fk_docid_usuarios
        FOREIGN KEY (id_usuarios)
        REFERENCES tienda.usuarios(id_usuarios),

    CONSTRAINT fk_docid_tipo
        FOREIGN KEY (id_tipo_documento)
        REFERENCES tienda.tipo_documento(id_tipo_documento)
);

-- =========================================
-- CREDENCIALES
-- =========================================

CREATE TABLE tienda.credenciales (
    id_credenciales   SERIAL PRIMARY KEY,
    id_usuarios       INT NOT NULL,
    password_hash     VARCHAR(255) NOT NULL,
    salt              VARCHAR(100),
    doble_factor      BOOLEAN NOT NULL DEFAULT FALSE,
    token_reset       VARCHAR(200),
    token_expira_en   TIMESTAMP,
    ultimo_login      TIMESTAMP,
    intentos_fallidos INT NOT NULL DEFAULT 0,
    bloqueado         BOOLEAN NOT NULL DEFAULT FALSE,
    activo            BOOLEAN NOT NULL DEFAULT TRUE,
    creado_en         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_cred_usuarios
        FOREIGN KEY (id_usuarios)
        REFERENCES tienda.usuarios(id_usuarios)
);

-- =========================================
-- SESIONES
-- =========================================

CREATE TABLE tienda.sesiones (
    id_sesiones      SERIAL PRIMARY KEY,
    id_usuarios      INT NOT NULL,
    token            VARCHAR(500) NOT NULL UNIQUE,
    dispositivo      VARCHAR(200),
    ip               VARCHAR(50),
    activa           BOOLEAN NOT NULL DEFAULT TRUE,
    expira_en        TIMESTAMP NOT NULL,
    creado_en        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_ses_usuarios
        FOREIGN KEY (id_usuarios)
        REFERENCES tienda.usuarios(id_usuarios)
);