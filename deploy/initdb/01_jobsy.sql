-- =====================================================================
--  Jobsy - carga completa para Supabase (PostgreSQL)
--  Generado a partir de la rama fix-ddl-desplegable de Jobsy_API,
--  con erratas de sintaxis corregidas y una semilla minima de datos.
--  Ejecutar completo, de una vez, en el SQL Editor de Supabase.
-- =====================================================================

CREATE SCHEMA IF NOT EXISTS catalogo;
CREATE SCHEMA IF NOT EXISTS core;
CREATE SCHEMA IF NOT EXISTS empleo;
CREATE SCHEMA IF NOT EXISTS soporte;
CREATE SCHEMA IF NOT EXISTS suscripciones;
CREATE SCHEMA IF NOT EXISTS tienda;


-- =====================================================================
-- Esquema EMPLEADOS/CATALOGO
-- =====================================================================
--CREATE DATABASE IF NOT EXISTS tienda_db;

--Creacion de Schema
CREATE SCHEMA IF NOT EXISTS catalogo;


CREATE TABLE catalogo.generos (
    id_genero       SERIAL          PRIMARY KEY,
    nombre          VARCHAR(30)     NOT NULL UNIQUE,   -- Masculino | Femenino | Otro
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en       TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE catalogo.niveles_experiencia (
    id_nivel_experiencia  SERIAL      PRIMARY KEY,
    nombre          VARCHAR(40)     NOT NULL UNIQUE,   -- sin_experiencia | intermedio | experto
    descripcion     VARCHAR(200),
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en       TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE catalogo.modalidades_trabajo (
    id_modalidad_trabajo  SERIAL      PRIMARY KEY,
    nombre          VARCHAR(40)     NOT NULL UNIQUE,   -- por_horas | tiempo_completo | por_jornadas | fines_semana | interno | dias_especificos
    descripcion     VARCHAR(200),
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en       TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE catalogo.turnos (
    id_turno        SERIAL          PRIMARY KEY,
    nombre          VARCHAR(30)     NOT NULL UNIQUE,   -- manana | tarde | noche | completo | jornada
    descripcion     VARCHAR(200),
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en       TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE catalogo.dias_semana (
    id_dia_semana   SERIAL          PRIMARY KEY,
    nombre          VARCHAR(10)     NOT NULL UNIQUE,   -- lun | mar | mie | jue | vie | sab | dom
    nombre_completo VARCHAR(20),                       -- Lunes | Martes | ...
    orden           INT             NOT NULL,          -- 1–7 para ordenar
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en       TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

  
CREATE TABLE catalogo.estados (
    id_estado       SERIAL          PRIMARY KEY,
    contexto        VARCHAR(40)     NOT NULL,          -- ver contextos válidos en cabecera
    nombre          VARCHAR(30)     NOT NULL,
    descripcion     VARCHAR(200),
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en       TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (contexto, nombre)
);


CREATE TABLE catalogo.tipos (
    id_tipo         SERIAL          PRIMARY KEY,
    contexto        VARCHAR(40)     NOT NULL,          -- ver contextos válidos en cabecera
    nombre          VARCHAR(50)     NOT NULL,
    descripcion     VARCHAR(200),
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en       TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (contexto, nombre)
);


CREATE TABLE catalogo.modalidades_suscripcion (
    id_modalidad_suscripcion SERIAL  PRIMARY KEY,
    nombre          VARCHAR(20)     NOT NULL UNIQUE,   -- mensual | anual
    descripcion     VARCHAR(200),
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en       TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- =====================================================================
-- Esquema CORE (usuarios)
-- =====================================================================
-- Esquema core: usuarios, roles, documentos de identidad, credenciales y sesiones.
--
-- ORDEN DE EJECUCION: este script depende de catalogo.generos (ver la FK de
-- core.usuarios), asi que el esquema catalogo debe crearse ANTES que este.

CREATE SCHEMA IF NOT EXISTS core;

CREATE TABLE core.roles_usuarios (
    id_roles_usuarios   SERIAL          PRIMARY KEY,
    nombre              VARCHAR(50)     NOT NULL UNIQUE,     -- admin | empleado | empleador
    descripcion         VARCHAR(200),
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE core.usuarios (
    id_usuarios         SERIAL          PRIMARY KEY,
    nombre              VARCHAR(80)     NOT NULL,
    apellido            VARCHAR(80)     NOT NULL,
    correo              VARCHAR(150)    NOT NULL UNIQUE,
    telefono            VARCHAR(20)     NOT NULL,
    ciudad              VARCHAR(80)     NOT NULL,
    fecha_nacimiento    DATE            NOT NULL,
    id_genero           INT             NOT NULL,            -- FK → catalogo.generos
    nacionalidad        VARCHAR(60)     NOT NULL,
    foto_url            VARCHAR(300)    NOT NULL,
    id_roles_usuarios   INT             NOT NULL,
    disponible          BOOLEAN         NOT NULL DEFAULT TRUE,
    perfil_publico      BOOLEAN         NOT NULL DEFAULT TRUE,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_usuarios_roles  FOREIGN KEY (id_roles_usuarios) REFERENCES core.roles_usuarios(id_roles_usuarios),
    CONSTRAINT fk_usuarios_genero FOREIGN KEY (id_genero)         REFERENCES catalogo.generos(id_genero)
);

CREATE TABLE core.tipo_documento (
    id_tipo_documento   SERIAL          PRIMARY KEY,
    nombre              VARCHAR(50)     NOT NULL UNIQUE,     -- cedula | pasaporte | tarjeta_extranjeria | nit | otro
    descripcion         VARCHAR(200),
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE core.documentos_identidad (
    id_documentos_identidad SERIAL      PRIMARY KEY,
    id_usuarios             INT         NOT NULL,
    id_tipo_documento       INT         NOT NULL,
    numero_documento        VARCHAR(30) NOT NULL UNIQUE,
    fecha_expedicion        DATE,
    lugar_expedicion        VARCHAR(100),
    activo                  BOOLEAN     NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (id_usuarios, id_tipo_documento),
    CONSTRAINT fk_docid_usuarios FOREIGN KEY (id_usuarios)       REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_docid_tipo     FOREIGN KEY (id_tipo_documento) REFERENCES core.tipo_documento(id_tipo_documento)
);

CREATE TABLE core.credenciales (
    id_credenciales     SERIAL          PRIMARY KEY,
    id_usuarios         INT             NOT NULL,
    password_hash       VARCHAR(255)    NOT NULL,
    salt                VARCHAR(100),
    doble_factor        BOOLEAN         NOT NULL DEFAULT FALSE,
    token_reset         VARCHAR(200),
    token_expira_en     TIMESTAMP,
    ultimo_login        TIMESTAMP,
    intentos_fallidos   INT             NOT NULL DEFAULT 0,
    bloqueado           BOOLEAN         NOT NULL DEFAULT FALSE,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_cred_usuarios FOREIGN KEY (id_usuarios) REFERENCES core.usuarios(id_usuarios)
);

-- Una credencial por usuario: sin esto, dos filas para el mismo id_usuarios
-- hacen ambiguo contra cual se valida el login.
CREATE UNIQUE INDEX ux_credenciales_usuario ON core.credenciales (id_usuarios);

CREATE TABLE core.sesiones (
    id_sesiones         SERIAL          PRIMARY KEY,
    id_usuarios         INT             NOT NULL,
    token               VARCHAR(500)    NOT NULL UNIQUE,
    dispositivo         VARCHAR(200),
    ip                  VARCHAR(50),
    activa              BOOLEAN         NOT NULL DEFAULT TRUE,
    expira_en           TIMESTAMP       NOT NULL,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_ses_usuarios FOREIGN KEY (id_usuarios) REFERENCES core.usuarios(id_usuarios)
);

-- Indices sobre las FK. PostgreSQL no los crea solo, y sin ellos cada join y
-- cada borrado en cascada hace recorrido secuencial.
CREATE INDEX ix_usuarios_rol         ON core.usuarios (id_roles_usuarios);
CREATE INDEX ix_usuarios_genero      ON core.usuarios (id_genero);
CREATE INDEX ix_usuarios_correo      ON core.usuarios (correo);
CREATE INDEX ix_docid_usuario        ON core.documentos_identidad (id_usuarios);
CREATE INDEX ix_docid_tipo           ON core.documentos_identidad (id_tipo_documento);
CREATE INDEX ix_sesiones_usuario     ON core.sesiones (id_usuarios);
CREATE INDEX ix_sesiones_expira      ON core.sesiones (expira_en) WHERE activa;


-- =====================================================================
-- Esquema EMPLEO
-- =====================================================================
CREATE SCHEMA IF NOT EXISTS empleo;
-- ============================================================
-- MÓDULO EMPLEO
-- ============================================================

CREATE TABLE empleo.especialidades (
    id_especialidades   SERIAL          PRIMARY KEY,
    nombre              VARCHAR(80)     NOT NULL UNIQUE,
    descripcion         VARCHAR(200),
    icono               VARCHAR(50),
    activa              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE empleo.tipo_inmueble (
    id_tipo_inmueble    SERIAL          PRIMARY KEY,
    nombre              VARCHAR(80)     NOT NULL UNIQUE,    -- casa | apartamento | oficina_local | finca_villa
    descripcion         VARCHAR(200)    NOT NULL,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE empleo.tamano_inmueble (
    id_tamano_inmueble  SERIAL          PRIMARY KEY,
    nombre              VARCHAR(80)     NOT NULL UNIQUE,    -- pequeno | mediano | grande | lujoso
    descripcion         VARCHAR(200)    NOT NULL,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE empleo.dias_trabajados (
    id_dias_trabajados  SERIAL          PRIMARY KEY,
    id_dia_semana       INT             NOT NULL,           -- FK → catalogo.dias_semana
    horas_trabajadas    VARCHAR(20),
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_dias_trab_dia FOREIGN KEY (id_dia_semana) REFERENCES catalogo.dias_semana(id_dia_semana)
);

CREATE TABLE empleo.perfiles_empleado (
    id_perfiles_empleado    SERIAL          PRIMARY KEY,
    id_usuarios             INT             NOT NULL UNIQUE,
    bio                     TEXT,
    id_nivel_experiencia    INT             NOT NULL,        -- FK → catalogo.niveles_experiencia
    anios_experiencia       INT,
    tarifa_hora             DECIMAL(10,2),
    tarifa_negociable       BOOLEAN         NOT NULL DEFAULT FALSE,
    id_modalidad_trabajo    INT             NOT NULL,        -- FK → catalogo.modalidades_trabajo
    id_turno_preferido      INT             NOT NULL,        -- FK → catalogo.turnos
    distancia_max_km        INT,
    barrio_zona             VARCHAR(100),
    hv_visible              BOOLEAN         NOT NULL DEFAULT TRUE,
    aparecer_busquedas      BOOLEAN         NOT NULL DEFAULT TRUE,
    rating_promedio         DECIMAL(3,2)    NOT NULL DEFAULT 0,
    total_resenas           INT             NOT NULL DEFAULT 0,
    total_trabajos          INT             NOT NULL DEFAULT 0,
    verificado              BOOLEAN         NOT NULL DEFAULT FALSE,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_perfil_usuarios  FOREIGN KEY (id_usuarios)          REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_perfil_nivel     FOREIGN KEY (id_nivel_experiencia)  REFERENCES catalogo.niveles_experiencia(id_nivel_experiencia),
    CONSTRAINT fk_perfil_modalidad FOREIGN KEY (id_modalidad_trabajo)  REFERENCES catalogo.modalidades_trabajo(id_modalidad_trabajo),
    CONSTRAINT fk_perfil_turno     FOREIGN KEY (id_turno_preferido)    REFERENCES catalogo.turnos(id_turno)
);

CREATE TABLE empleo.TR_perfil_especialidades (
    id_TR_perfil_especialidades SERIAL      PRIMARY KEY,
    id_perfiles_empleado        INT         NOT NULL,
    id_especialidades           INT         NOT NULL,
    activo                      BOOLEAN     NOT NULL DEFAULT TRUE,
    creado_en                   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en              TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (id_perfiles_empleado, id_especialidades),
    CONSTRAINT fk_perfil_esp_perfil FOREIGN KEY (id_perfiles_empleado) REFERENCES empleo.perfiles_empleado(id_perfiles_empleado),
    CONSTRAINT fk_perfil_esp_esp    FOREIGN KEY (id_especialidades)     REFERENCES empleo.especialidades(id_especialidades)
);

CREATE TABLE empleo.disponibilidad_empleado (
    id_disponibilidad_empleado  SERIAL      PRIMARY KEY,
    id_perfiles_empleado        INT         NOT NULL,
    id_dia_semana               INT         NOT NULL,        -- FK → catalogo.dias_semana
    id_turno                    INT         NOT NULL,        -- FK → catalogo.turnos
    hora_inicio                 VARCHAR(8),
    hora_fin                    VARCHAR(8),
    activo                      BOOLEAN     NOT NULL DEFAULT TRUE,
    creado_en                   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en              TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (id_perfiles_empleado, id_dia_semana, id_turno),
    CONSTRAINT fk_disp_perfil FOREIGN KEY (id_perfiles_empleado) REFERENCES empleo.perfiles_empleado(id_perfiles_empleado),
    CONSTRAINT fk_disp_dia    FOREIGN KEY (id_dia_semana)        REFERENCES catalogo.dias_semana(id_dia_semana),
    CONSTRAINT fk_disp_turno  FOREIGN KEY (id_turno)             REFERENCES catalogo.turnos(id_turno)
);

CREATE TABLE empleo.certificaciones (
    id_certificaciones  SERIAL          PRIMARY KEY,
    id_perfiles_empleado INT            NOT NULL,
    nombre              VARCHAR(150)    NOT NULL,
    entidad_emisora     VARCHAR(150),
    anio_emision        INT,
    url_documento       VARCHAR(300),
    id_estado           INT             NOT NULL,            -- FK → catalogo.estados (contexto='certificacion')
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_cert_perfil  FOREIGN KEY (id_perfiles_empleado) REFERENCES empleo.perfiles_empleado(id_perfiles_empleado),
    CONSTRAINT fk_cert_estado  FOREIGN KEY (id_estado)            REFERENCES catalogo.estados(id_estado)
);
 
CREATE TABLE empleo.documentos_usuario (
    id_documentos_usuario   SERIAL      PRIMARY KEY,
    id_usuarios             INT         NOT NULL,
    id_tipo                 INT         NOT NULL,            -- FK → catalogo.tipos (contexto='documento_usuario')
    url_archivo             VARCHAR(300) NOT NULL,
    id_estado               INT         NOT NULL,            -- FK → catalogo.estados (contexto='documento')
    notas                   TEXT,
    activo                  BOOLEAN     NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    verificado_en           TIMESTAMP,
    CONSTRAINT fk_docu_usuarios FOREIGN KEY (id_usuarios) REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_docu_tipo     FOREIGN KEY (id_tipo)     REFERENCES catalogo.tipos(id_tipo),
    CONSTRAINT fk_docu_estado   FOREIGN KEY (id_estado)   REFERENCES catalogo.estados(id_estado)
);
 
CREATE TABLE empleo.ofertas (
    id_ofertas              SERIAL          PRIMARY KEY,
    id_empleador            INT             NOT NULL,
    id_dias_trabajados      INT,
    titulo                  VARCHAR(150)    NOT NULL,
    descripcion             TEXT,
    id_tipo_inmueble        INT             NOT NULL,
    id_tamano_inmueble      INT             NOT NULL,
    num_habitaciones        INT,
    num_banos               INT,
    barrio                  VARCHAR(100),
    direccion               VARCHAR(200)    NOT NULL,
    ciudad                  VARCHAR(80)     NOT NULL,
    tarifa_hora             DECIMAL(10,2)   NOT NULL,
    tarifa_jornada          DECIMAL(10,2),
    horario_descripcion     VARCHAR(200),
    id_modalidad_trabajo    INT             NOT NULL,        -- FK → catalogo.modalidades_trabajo
    insumos_cliente         BOOLEAN         NOT NULL DEFAULT TRUE,
    vigilancia              BOOLEAN         NOT NULL DEFAULT FALSE,
    id_tipo_mascota         INT             NOT NULL,        -- FK → catalogo.tipos (contexto='mascota')
    condiciones_extra       TEXT,
    id_estado               INT             NOT NULL,        -- FK → catalogo.estados (contexto='oferta')
    destacada               BOOLEAN         NOT NULL DEFAULT FALSE,
    nueva                   BOOLEAN         NOT NULL DEFAULT TRUE,
    vistas                  INT             NOT NULL DEFAULT 0,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expira_en               TIMESTAMP       NOT NULL,
    CONSTRAINT fk_ofertas_empleador  FOREIGN KEY (id_empleador)       REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_ofertas_tipo_inm   FOREIGN KEY (id_tipo_inmueble)   REFERENCES empleo.tipo_inmueble(id_tipo_inmueble),
    CONSTRAINT fk_ofertas_tamano     FOREIGN KEY (id_tamano_inmueble) REFERENCES empleo.tamano_inmueble(id_tamano_inmueble),
    CONSTRAINT fk_ofertas_dias       FOREIGN KEY (id_dias_trabajados) REFERENCES empleo.dias_trabajados(id_dias_trabajados),
    CONSTRAINT fk_ofertas_modalidad  FOREIGN KEY (id_modalidad_trabajo) REFERENCES catalogo.modalidades_trabajo(id_modalidad_trabajo),
    CONSTRAINT fk_ofertas_mascota    FOREIGN KEY (id_tipo_mascota)    REFERENCES catalogo.tipos(id_tipo),
    CONSTRAINT fk_ofertas_estado     FOREIGN KEY (id_estado)          REFERENCES catalogo.estados(id_estado)
);
 
CREATE TABLE empleo.TR_oferta_servicios (
    id_TR_oferta_servicios  SERIAL          PRIMARY KEY,
    id_ofertas              INT             NOT NULL,
    id_especialidades       INT             NOT NULL,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (id_ofertas, id_especialidades),
    CONSTRAINT fk_oferta_srv_oferta FOREIGN KEY (id_ofertas)        REFERENCES empleo.ofertas(id_ofertas),
    CONSTRAINT fk_oferta_srv_esp    FOREIGN KEY (id_especialidades) REFERENCES empleo.especialidades(id_especialidades)
);
 
CREATE TABLE empleo.oferta_fotos (
    id_oferta_fotos     SERIAL          PRIMARY KEY,
    id_ofertas          INT             NOT NULL,
    url_foto            VARCHAR(300)    NOT NULL,
    orden               INT             NOT NULL DEFAULT 0,
    es_principal        BOOLEAN         NOT NULL DEFAULT FALSE,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_oferta_fotos_oferta FOREIGN KEY (id_ofertas) REFERENCES empleo.ofertas(id_ofertas)
);
 
CREATE TABLE empleo.hojas_vida (
    id_hojas_vida       SERIAL          PRIMARY KEY,
    id_usuarios         INT             NOT NULL,
    nombre              VARCHAR(150)    NOT NULL,
    resumen             TEXT,
    es_predeterminada   BOOLEAN         NOT NULL DEFAULT FALSE,
    url_pdf             VARCHAR(300),
    es_publica          BOOLEAN         NOT NULL DEFAULT FALSE,
    progreso_pct        INT             NOT NULL DEFAULT 0,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_hv_usuarios FOREIGN KEY (id_usuarios) REFERENCES core.usuarios(id_usuarios)
);

CREATE TABLE empleo.hv_experiencias (
    id_hv_experiencias  SERIAL          PRIMARY KEY,
    id_hojas_vida       INT             NOT NULL,
    id_tipo             INT             NOT NULL,            -- FK → catalogo.tipos (contexto='experiencia_hv')
    descripcion         TEXT,
    empresa             VARCHAR(150),
    cargo               VARCHAR(100),
    fecha_inicio        DATE            NOT NULL,
    fecha_fin           DATE,
    activo              BOOLEAN         NOT NULL DEFAULT FALSE, -- true = trabajo actual
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_hv_exp_hv   FOREIGN KEY (id_hojas_vida) REFERENCES empleo.hojas_vida(id_hojas_vida),
    CONSTRAINT fk_hv_exp_tipo FOREIGN KEY (id_tipo)       REFERENCES catalogo.tipos(id_tipo)
);

CREATE TABLE empleo.postulaciones (
    id_postulaciones        SERIAL          PRIMARY KEY,
    id_ofertas              INT             NOT NULL,
    id_empleado             INT             NOT NULL,
    id_hojas_vida           INT             NOT NULL,
    id_estado               INT             NOT NULL,        -- FK → catalogo.estados (contexto='postulacion')
    disponible_desde        DATE            NOT NULL,
    id_modalidad_disp       INT             NOT NULL,        -- FK → catalogo.modalidades_trabajo
    id_turno_postulacion    INT             NOT NULL,        -- FK → catalogo.turnos
    nota_disponibilidad     TEXT,
    id_forma_cobro          INT             NOT NULL,        -- FK → catalogo.tipos (contexto='forma_cobro')
    tarifa_min              DECIMAL(10,2),
    tarifa_max              DECIMAL(10,2),
    mensaje_empleador       TEXT,
    acepta_terminos         BOOLEAN         NOT NULL DEFAULT FALSE,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    postulado_en            TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    respondido_en           TIMESTAMP,
    nota_rechazo            TEXT,
    origen_postulacion      BOOLEAN,                         -- true=empleador | false=empleado
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_post_oferta       FOREIGN KEY (id_ofertas)          REFERENCES empleo.ofertas(id_ofertas),
    CONSTRAINT fk_post_empleado     FOREIGN KEY (id_empleado)         REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_post_hv           FOREIGN KEY (id_hojas_vida)       REFERENCES empleo.hojas_vida(id_hojas_vida),
    CONSTRAINT fk_post_forma_cobro  FOREIGN KEY (id_forma_cobro)      REFERENCES catalogo.tipos(id_tipo),
    CONSTRAINT fk_post_estado       FOREIGN KEY (id_estado)           REFERENCES catalogo.estados(id_estado),
    CONSTRAINT fk_post_modalidad    FOREIGN KEY (id_modalidad_disp)   REFERENCES catalogo.modalidades_trabajo(id_modalidad_trabajo),
    CONSTRAINT fk_post_turno        FOREIGN KEY (id_turno_postulacion) REFERENCES catalogo.turnos(id_turno)
);
 
CREATE TABLE empleo.TR_postulacion_dias (
    id_TR_postulacion_dias  SERIAL          PRIMARY KEY,
    id_postulaciones        INT             NOT NULL,
    id_dia_semana           INT             NOT NULL,        -- FK → catalogo.dias_semana
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (id_postulaciones, id_dia_semana),
    CONSTRAINT fk_post_dias_post FOREIGN KEY (id_postulaciones) REFERENCES empleo.postulaciones(id_postulaciones),
    CONSTRAINT fk_post_dias_dia  FOREIGN KEY (id_dia_semana)    REFERENCES catalogo.dias_semana(id_dia_semana)
);
 
CREATE TABLE empleo.postulacion_chips (
    id_postulacion_chips    SERIAL          PRIMARY KEY,
    id_postulaciones        INT             NOT NULL,
    texto                   VARCHAR(100)    NOT NULL,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_chips_post FOREIGN KEY (id_postulaciones) REFERENCES empleo.postulaciones(id_postulaciones)
);
 
CREATE TABLE empleo.resenas (
    id_resenas          SERIAL          PRIMARY KEY,
    id_postulaciones    INT             NOT NULL,
    id_reviewer         INT             NOT NULL,
    id_reviewed         INT             NOT NULL,
    calificacion        DECIMAL(2,1)    NOT NULL CHECK (calificacion BETWEEN 1.0 AND 5.0),
    comentario          TEXT,
    id_tipo             INT             NOT NULL,            -- FK → catalogo.tipos (contexto='resena')
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_resenas_post     FOREIGN KEY (id_postulaciones) REFERENCES empleo.postulaciones(id_postulaciones),
    CONSTRAINT fk_resenas_reviewer FOREIGN KEY (id_reviewer)      REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_resenas_reviewed FOREIGN KEY (id_reviewed)      REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_resenas_tipo     FOREIGN KEY (id_tipo)          REFERENCES catalogo.tipos(id_tipo)
);
 
CREATE TABLE empleo.contrataciones (
    id_contrataciones       SERIAL          PRIMARY KEY,
    id_empleador            INT             NOT NULL,
    id_empleado             INT             NOT NULL,
    id_ofertas              INT             NOT NULL,
    id_postulaciones        INT             NOT NULL,
    id_estado               INT             NOT NULL,        -- FK → catalogo.estados (contexto='contratacion')
    direccion_servicio      VARCHAR(200)    NOT NULL,
    fecha_inicio            DATE            NOT NULL,
    hora_inicio             VARCHAR(8)          NOT NULL,
    duracion_horas          DECIMAL(4,1),
    id_modalidad            INT             NOT NULL,        -- FK → catalogo.modalidades_trabajo
    tarifa_acordada         DECIMAL(10,2)   NOT NULL,
    total_estimado          DECIMAL(12,2),
    descripcion_trabajo     TEXT,
    nota_adicional          TEXT,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    finalizado_en           TIMESTAMP,
    CONSTRAINT fk_cont_empleador FOREIGN KEY (id_empleador)   REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_cont_empleado  FOREIGN KEY (id_empleado)    REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_cont_oferta    FOREIGN KEY (id_ofertas)     REFERENCES empleo.ofertas(id_ofertas),
    CONSTRAINT fk_cont_post      FOREIGN KEY (id_postulaciones) REFERENCES empleo.postulaciones(id_postulaciones),
    CONSTRAINT fk_cont_estado    FOREIGN KEY (id_estado)      REFERENCES catalogo.estados(id_estado),
    CONSTRAINT fk_cont_modalidad FOREIGN KEY (id_modalidad)   REFERENCES catalogo.modalidades_trabajo(id_modalidad_trabajo)
);
 
CREATE TABLE empleo.TR_contratacion_especialidades (
    id_TR_contratacion_especialidades SERIAL  PRIMARY KEY,
    id_contrataciones   INT             NOT NULL,
    id_especialidades   INT             NOT NULL,
    activo              BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en           TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (id_contrataciones, id_especialidades),
    CONSTRAINT fk_cont_esp_cont FOREIGN KEY (id_contrataciones) REFERENCES empleo.contrataciones(id_contrataciones),
    CONSTRAINT fk_cont_esp_esp  FOREIGN KEY (id_especialidades) REFERENCES empleo.especialidades(id_especialidades)
);
 
CREATE TABLE empleo.horarios_agendados (
    id_horarios_agendados   SERIAL          PRIMARY KEY,
    id_contrataciones       INT             NOT NULL,
    id_dia_semana           INT             NOT NULL,        -- FK → catalogo.dias_semana
    id_turno                INT             NOT NULL,        -- FK → catalogo.turnos
    hora_inicio             VARCHAR(8)          NOT NULL,
    hora_fin                VARCHAR(8)          NOT NULL,
    nota                    TEXT,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_horarios_cont  FOREIGN KEY (id_contrataciones) REFERENCES empleo.contrataciones(id_contrataciones),
    CONSTRAINT fk_horarios_dia   FOREIGN KEY (id_dia_semana)     REFERENCES catalogo.dias_semana(id_dia_semana),
    CONSTRAINT fk_horarios_turno FOREIGN KEY (id_turno)          REFERENCES catalogo.turnos(id_turno)
);


-- =====================================================================
-- Esquema SOPORTE
-- =====================================================================
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


-- =====================================================================
-- Esquema SUSCRIPCIONES
-- =====================================================================
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


-- =====================================================================
-- Esquema TIENDA
-- =====================================================================
-- Esquema tienda: catalogo de productos, carrito, pedidos, pagos y saldo.
--
-- Este archivo contenia una copia del esquema de core (roles_usuarios,
-- usuarios, credenciales, sesiones...), de modo que ninguna de las 13 tablas
-- que los modelos de Tienda_APi esperan llegaba a existir. El script de
-- inserts, que si apunta a las tablas correctas, fallaba entero.
--
-- Las definiciones de abajo se derivan de los structs de Tienda_APi/models/,
-- respetando el nombre de columna de cada tag orm y la nulabilidad declarada.
--
-- ORDEN DE EJECUCION: depende de los esquemas catalogo y core, que deben
-- crearse antes.

CREATE SCHEMA IF NOT EXISTS tienda;

CREATE TABLE tienda.categorias_tienda (
    id_categorias_tienda    SERIAL          PRIMARY KEY,
    nombre                  VARCHAR(100)    NOT NULL,
    icono                   VARCHAR(50),
    activa                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tienda.productos (
    id_productos            SERIAL          PRIMARY KEY,
    id_categorias_tienda    INT             NOT NULL,
    nombre                  VARCHAR(150)    NOT NULL,
    descripcion             TEXT,
    id_tipo                 INT             NOT NULL,   -- FK → catalogo.tipos (contexto='producto_tienda')
    precio                  DECIMAL(12,2)   NOT NULL,
    precio_original         DECIMAL(12,2),
    descuento_pct           INT             NOT NULL DEFAULT 0,
    stock                   INT             NOT NULL DEFAULT 0,
    imagen_url              VARCHAR(300),
    es_nuevo                BOOLEAN         NOT NULL DEFAULT FALSE,
    destacado               BOOLEAN         NOT NULL DEFAULT FALSE,
    envio_gratis            BOOLEAN         NOT NULL DEFAULT FALSE,
    rating_promedio         DECIMAL(3,2)    NOT NULL DEFAULT 0,
    total_resenas           INT             NOT NULL DEFAULT 0,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_productos_categoria FOREIGN KEY (id_categorias_tienda) REFERENCES tienda.categorias_tienda(id_categorias_tienda),
    CONSTRAINT fk_productos_tipo      FOREIGN KEY (id_tipo)              REFERENCES catalogo.tipos(id_tipo),
    -- El backend no valida nada; estas restricciones son la unica defensa
    -- contra un precio negativo o un stock negativo.
    CONSTRAINT ck_productos_precio    CHECK (precio >= 0),
    CONSTRAINT ck_productos_stock     CHECK (stock >= 0),
    CONSTRAINT ck_productos_descuento CHECK (descuento_pct BETWEEN 0 AND 100)
);

CREATE TABLE tienda.cupones (
    id_cupones              SERIAL          PRIMARY KEY,
    codigo                  VARCHAR(50)     NOT NULL UNIQUE,
    descuento_pct           INT,
    descuento_fijo          DECIMAL(12,2),
    usos_maximos            INT,
    usos_actuales           INT             NOT NULL DEFAULT 0,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    expira_en               TIMESTAMP,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT ck_cupones_pct   CHECK (descuento_pct IS NULL OR descuento_pct BETWEEN 0 AND 100),
    CONSTRAINT ck_cupones_fijo  CHECK (descuento_fijo IS NULL OR descuento_fijo >= 0),
    CONSTRAINT ck_cupones_usos  CHECK (usos_maximos IS NULL OR usos_actuales <= usos_maximos)
);

CREATE TABLE tienda.carritos (
    id_carritos             SERIAL          PRIMARY KEY,
    id_usuarios             INT             NOT NULL,   -- FK → core.usuarios
    id_estado               INT             NOT NULL,   -- FK → catalogo.estados (contexto='carrito')
    id_cupones              INT,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_carritos_usuarios FOREIGN KEY (id_usuarios) REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_carritos_estado   FOREIGN KEY (id_estado)   REFERENCES catalogo.estados(id_estado),
    CONSTRAINT fk_carritos_cupon    FOREIGN KEY (id_cupones)  REFERENCES tienda.cupones(id_cupones)
);

CREATE TABLE tienda.tr_carrito_items (
    id_tr_carrito_items     SERIAL          PRIMARY KEY,
    id_carritos             INT             NOT NULL,
    id_productos            INT             NOT NULL,
    cantidad                INT             NOT NULL,
    precio_unitario         DECIMAL(12,2)   NOT NULL,
    seleccionado            BOOLEAN         NOT NULL DEFAULT TRUE,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (id_carritos, id_productos),
    CONSTRAINT fk_carrito_items_carrito  FOREIGN KEY (id_carritos)  REFERENCES tienda.carritos(id_carritos),
    CONSTRAINT fk_carrito_items_producto FOREIGN KEY (id_productos) REFERENCES tienda.productos(id_productos),
    CONSTRAINT ck_carrito_items_cantidad CHECK (cantidad > 0),
    CONSTRAINT ck_carrito_items_precio   CHECK (precio_unitario >= 0)
);

-- OJO con el nombre: la entidad se llama Facturacion pero su PK es id_pedidos
-- y las demas tablas la referencian como "pedidos". Son la misma cosa con tres
-- nombres; se conserva asi porque los modelos Go dependen de estos nombres.
CREATE TABLE tienda.facturacion (
    id_pedidos              SERIAL          PRIMARY KEY,
    id_usuarios             INT             NOT NULL,   -- FK → core.usuarios
    id_carritos             INT             NOT NULL,
    id_estado               INT             NOT NULL,   -- FK → catalogo.estados (contexto='pedido')
    subtotal                DECIMAL(12,2)   NOT NULL DEFAULT 0,
    descuento_total         DECIMAL(12,2)   NOT NULL DEFAULT 0,
    cupon_descuento         DECIMAL(12,2)   NOT NULL DEFAULT 0,
    envio                   DECIMAL(12,2)   NOT NULL DEFAULT 0,
    total                   DECIMAL(12,2)   NOT NULL DEFAULT 0,
    factura_nombre          VARCHAR(150),
    factura_doc             VARCHAR(30),
    factura_correo          VARCHAR(150),
    factura_telefono        VARCHAR(20),
    factura_direccion       VARCHAR(200),
    factura_ciudad          VARCHAR(80),
    factura_dpto            VARCHAR(80),
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_facturacion_usuarios FOREIGN KEY (id_usuarios) REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_facturacion_carrito  FOREIGN KEY (id_carritos) REFERENCES tienda.carritos(id_carritos),
    CONSTRAINT fk_facturacion_estado   FOREIGN KEY (id_estado)   REFERENCES catalogo.estados(id_estado),
    -- El servidor no calcula estos importes: los recibe del cliente. Estas
    -- comprobaciones evitan al menos los valores imposibles.
    CONSTRAINT ck_facturacion_montos   CHECK (subtotal >= 0 AND total >= 0 AND envio >= 0)
);

CREATE TABLE tienda.pedido_items (
    id_pedido_items         SERIAL          PRIMARY KEY,
    id_pedidos              INT             NOT NULL,
    id_productos            INT             NOT NULL,
    cantidad                INT             NOT NULL,
    precio_unitario         DECIMAL(12,2)   NOT NULL,
    subtotal                DECIMAL(12,2)   NOT NULL,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_pedido_items_pedido   FOREIGN KEY (id_pedidos)   REFERENCES tienda.facturacion(id_pedidos),
    CONSTRAINT fk_pedido_items_producto FOREIGN KEY (id_productos) REFERENCES tienda.productos(id_productos),
    CONSTRAINT ck_pedido_items_cantidad CHECK (cantidad > 0)
);

CREATE TABLE tienda.metodo_pago (
    id_metodo_pago          SERIAL          PRIMARY KEY,
    nombre                  VARCHAR(50)     NOT NULL,   -- tarjeta | pse | nequi | daviplata | saldo_jobsy
    descripcion             VARCHAR(200)    NOT NULL,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tienda.metodos_pago_guardados (
    id_metodos_pago_guardados SERIAL        PRIMARY KEY,
    id_usuarios             INT             NOT NULL,   -- FK → core.usuarios
    id_metodo_pago          INT             NOT NULL,
    alias                   VARCHAR(80),
    numero_cuenta           VARCHAR(50)     NOT NULL,
    id_tipo_cuenta          INT             NOT NULL,   -- FK → catalogo.tipos (contexto='tipo_cuenta')
    es_principal            BOOLEAN         NOT NULL DEFAULT FALSE,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_mpg_usuarios FOREIGN KEY (id_usuarios)     REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_mpg_metodo   FOREIGN KEY (id_metodo_pago)  REFERENCES tienda.metodo_pago(id_metodo_pago),
    CONSTRAINT fk_mpg_tipo_cta FOREIGN KEY (id_tipo_cuenta)  REFERENCES catalogo.tipos(id_tipo)
);

-- Un solo metodo principal por usuario.
CREATE UNIQUE INDEX ux_mpg_principal ON tienda.metodos_pago_guardados (id_usuarios) WHERE es_principal AND activo;

CREATE TABLE tienda.pagos (
    id_pagos                SERIAL          PRIMARY KEY,
    id_pedidos              INT,
    -- Puente hacia empleo.contrataciones. Se deja como entero suelto y
    -- NULLABLE a proposito: un pago es de un pedido de tienda O de una
    -- contratacion, nunca de ambos, y el modelo Go los declara a los dos como
    -- obligatorios, lo que haria imposible registrar cualquiera de los dos
    -- casos. Es un problema de diseno del modelo, no del DDL.
    id_contrataciones       INT,
    id_usuarios             INT             NOT NULL,   -- FK → core.usuarios
    id_metodo_pago          INT             NOT NULL,
    id_tipo_pago            INT             NOT NULL,   -- FK → catalogo.tipos (contexto='tipo_pago')
    id_estado               INT             NOT NULL,   -- FK → catalogo.estados (contexto='pago')
    -- Nullable, aunque el modelo Go lo declare obligatorio: un pago con
    -- tarjeta no tiene tipo de cuenta bancaria, y los propios inserts de
    -- semilla omiten esta columna. Con NOT NULL, el seed no entra.
    id_tipo_cuenta          INT,            -- FK → catalogo.tipos (contexto='cuenta_bancaria')
    monto                   DECIMAL(12,2)   NOT NULL,
    referencia_ext          VARCHAR(100),
    banco                   VARCHAR(100),
    titular_tarjeta         VARCHAR(150),
    ultimos4                VARCHAR(4),
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_pagos_pedido   FOREIGN KEY (id_pedidos)     REFERENCES tienda.facturacion(id_pedidos),
    CONSTRAINT fk_pagos_usuarios FOREIGN KEY (id_usuarios)    REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_pagos_metodo   FOREIGN KEY (id_metodo_pago) REFERENCES tienda.metodo_pago(id_metodo_pago),
    CONSTRAINT fk_pagos_tipo     FOREIGN KEY (id_tipo_pago)   REFERENCES catalogo.tipos(id_tipo),
    CONSTRAINT fk_pagos_estado   FOREIGN KEY (id_estado)      REFERENCES catalogo.estados(id_estado),
    CONSTRAINT fk_pagos_tipo_cta FOREIGN KEY (id_tipo_cuenta) REFERENCES catalogo.tipos(id_tipo),
    CONSTRAINT ck_pagos_monto    CHECK (monto >= 0),
    -- Un pago cuelga de un pedido o de una contratacion, no de ninguno ni de
    -- los dos.
    CONSTRAINT ck_pagos_origen   CHECK ((id_pedidos IS NULL) <> (id_contrataciones IS NULL))
);

CREATE TABLE tienda.tipo_transaccion (
    id_tipo_transaccion     SERIAL          PRIMARY KEY,
    nombre                  VARCHAR(50)     NOT NULL,   -- ingreso | egreso | reembolso
    descripcion             VARCHAR(200)    NOT NULL,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tienda.transacciones (
    id_transacciones        SERIAL          PRIMARY KEY,
    id_usuarios             INT             NOT NULL,   -- FK → core.usuarios
    id_pagos                INT             NOT NULL,
    id_tipo_transaccion     INT             NOT NULL,
    monto                   DECIMAL(12,2)   NOT NULL,
    descripcion             TEXT,
    referencia              VARCHAR(100),
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_trans_usuarios FOREIGN KEY (id_usuarios)          REFERENCES core.usuarios(id_usuarios),
    CONSTRAINT fk_trans_pagos    FOREIGN KEY (id_pagos)             REFERENCES tienda.pagos(id_pagos),
    CONSTRAINT fk_trans_tipo     FOREIGN KEY (id_tipo_transaccion)  REFERENCES tienda.tipo_transaccion(id_tipo_transaccion)
);

CREATE TABLE tienda.saldo_jobsy (
    id_saldo_jobsy          SERIAL          PRIMARY KEY,
    id_usuarios             INT             NOT NULL,   -- FK → core.usuarios
    saldo_disponible        DECIMAL(12,2)   NOT NULL DEFAULT 0,
    total_ganado            DECIMAL(12,2)   NOT NULL DEFAULT 0,
    total_retirado          DECIMAL(12,2)   NOT NULL DEFAULT 0,
    proximo_pago            DATE,
    activo                  BOOLEAN         NOT NULL DEFAULT TRUE,
    creado_en               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actualizado_en          TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_saldo_usuarios FOREIGN KEY (id_usuarios) REFERENCES core.usuarios(id_usuarios),
    -- El endpoint PUT /v1/saldo_jobsy/:id permite fijar cualquier saldo sin
    -- autenticacion. Mientras eso siga asi, esto al menos impide un saldo
    -- negativo.
    CONSTRAINT ck_saldo_no_negativo CHECK (saldo_disponible >= 0 AND total_ganado >= 0 AND total_retirado >= 0)
);

-- Un unico monedero por usuario.
CREATE UNIQUE INDEX ux_saldo_usuario ON tienda.saldo_jobsy (id_usuarios);

-- Indices sobre las claves foraneas y las columnas por las que se filtra.
CREATE INDEX ix_productos_categoria   ON tienda.productos (id_categorias_tienda);
CREATE INDEX ix_productos_activo      ON tienda.productos (activo) WHERE activo;
CREATE INDEX ix_carritos_usuario      ON tienda.carritos (id_usuarios);
CREATE INDEX ix_carrito_items_carrito ON tienda.tr_carrito_items (id_carritos);
CREATE INDEX ix_facturacion_usuario   ON tienda.facturacion (id_usuarios);
CREATE INDEX ix_facturacion_estado    ON tienda.facturacion (id_estado);
CREATE INDEX ix_pedido_items_pedido   ON tienda.pedido_items (id_pedidos);
CREATE INDEX ix_pagos_pedido          ON tienda.pagos (id_pedidos);
CREATE INDEX ix_pagos_usuario         ON tienda.pagos (id_usuarios);
CREATE INDEX ix_pagos_contratacion    ON tienda.pagos (id_contrataciones);
CREATE INDEX ix_transacciones_usuario ON tienda.transacciones (id_usuarios);
CREATE INDEX ix_mpg_usuario           ON tienda.metodos_pago_guardados (id_usuarios);


-- =====================================================================
-- DATOS: catalogo
-- =====================================================================
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


-- =====================================================================
-- DATOS: semilla minima (roles, usuarios, estados)
-- =====================================================================
-- Semilla minima de desarrollo (NO son datos reales).
-- Cubre los dos huecos de la rama: roles+usuarios en core y los estados de
-- catalogo que los inserts resuelven por subconsulta (contexto, nombre).

-- ---- core.roles_usuarios ----
INSERT INTO core.roles_usuarios (id_roles_usuarios, nombre, descripcion) VALUES
    (1, 'admin',     'Administrador de la plataforma'),
    (2, 'empleado',  'Persona que ofrece servicios'),
    (3, 'empleador', 'Persona que contrata servicios');

-- ---- core.usuarios (ids fijos: los inserts los referencian como 1,2,3) ----
INSERT INTO core.usuarios
    (id_usuarios, nombre, apellido, correo, telefono, ciudad, fecha_nacimiento,
     id_genero, nacionalidad, foto_url, id_roles_usuarios) VALUES
    (1, 'Laura',     'Martínez', 'laura@jobsy.co',  '3001112233', 'Medellín', '1995-04-12',
        2, 'Colombiana', 'https://cdn.jobsy.co/users/laura.jpg',  2),
    (2, 'Carlos',    'Gómez',    'carlos@jobsy.co', '3004445566', 'Medellín', '1988-09-30',
        1, 'Colombiana', 'https://cdn.jobsy.co/users/carlos.jpg', 3),
    (3, 'Valentina', 'Ospina',   'vale@jobsy.co',   '3007778899', 'Envigado', '1992-01-25',
        2, 'Colombiana', 'https://cdn.jobsy.co/users/vale.jpg',   2);

-- ---- catalogo.estados (las 17 parejas que buscan los inserts) ----
INSERT INTO catalogo.estados (contexto, nombre, descripcion) VALUES
    ('carrito',       'activo',      'Carrito en uso'),
    ('carrito',       'pagado',      'Carrito ya pagado'),
    ('certificacion', 'pendiente',   'Certificación por validar'),
    ('certificacion', 'validado',    'Certificación validada'),
    ('contratacion',  'activa',      'Contratación en curso'),
    ('contratacion',  'completada',  'Contratación finalizada'),
    ('documento',     'pendiente',   'Documento por verificar'),
    ('documento',     'verificado',  'Documento verificado'),
    ('oferta',        'activa',      'Oferta publicada y vigente'),
    ('pago',          'aprobado',    'Pago aprobado'),
    ('pago',          'pendiente',   'Pago pendiente'),
    ('pedido',        'pagado',      'Pedido pagado'),
    ('pedido',        'pendiente',   'Pedido pendiente de pago'),
    ('postulacion',   'pendiente',   'Postulación pendiente'),
    ('postulacion',   'revisando',   'Postulación en revisión'),
    ('reporte',       'abierto',     'Reporte abierto'),
    ('reporte',       'en_revision', 'Reporte en revisión');


-- =====================================================================
-- DATOS: empleo
-- =====================================================================
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


-- =====================================================================
-- DATOS: soporte
-- =====================================================================
INSERT INTO soporte.reportes
    (id_usuarios, id_tipo, descripcion, id_estado)
VALUES
    (1,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='reporte' AND nombre='tecnico'),
     'El botón de postulación no responde en móvil Chrome.',
     (SELECT id_estado FROM catalogo.estados WHERE contexto='reporte' AND nombre='abierto')),
    (2,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='reporte' AND nombre='pago'),
     'Se me realizó un cobro duplicado por la suscripción Pro.',
     (SELECT id_estado FROM catalogo.estados WHERE contexto='reporte' AND nombre='en_revision')),
    (3,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='reporte' AND nombre='usuario'),
     'Un usuario me envió mensajes ofensivos en el chat.',
     (SELECT id_estado FROM catalogo.estados WHERE contexto='reporte' AND nombre='abierto'));
 
-- ---- soporte.reporte_mensajes --------------------------------
INSERT INTO soporte.reporte_mensajes
    (id_reportes, id_remitente, mensaje)
VALUES
    (1, 1, 'El problema ocurre desde ayer, lo probé en iOS y Android también.'),
    (2, 2, 'Adjunto captura del extracto bancario con el doble cargo.'),
    (2, 2, 'Ya hablé con mi banco y me confirmaron que salieron dos débitos.');


-- =====================================================================
-- DATOS: tienda
-- =====================================================================
-- ---- tienda.metodo_pago --------------------------------------
INSERT INTO tienda.metodo_pago (nombre, descripcion) VALUES
    ('tarjeta',     'Tarjeta débito o crédito Visa / Mastercard'),
    ('pse',         'Pago en línea PSE desde cuenta bancaria'),
    ('nequi',       'Pago desde billetera digital Nequi'),
    ('daviplata',   'Pago desde billetera digital Daviplata'),
    ('saldo_jobsy', 'Pago descontado del saldo disponible en la plataforma');

-- ---- tienda.tipo_transaccion ---------------------------------
INSERT INTO tienda.tipo_transaccion (nombre, descripcion) VALUES
    ('ingreso',   'Entrada de dinero a la cuenta del usuario'),
    ('egreso',    'Salida de dinero de la cuenta del usuario'),
    ('reembolso', 'Devolución de fondos al usuario por cancelación');

-- ---- tienda.categorias_tienda --------------------------------
INSERT INTO tienda.categorias_tienda (nombre, icono) VALUES
    ('Limpieza del hogar',    'icon-broom'),
    ('Equipos y herramientas','icon-wrench'),
    ('Uniformes y ropa',      'icon-shirt');


-- ---- tienda.productos ----------------------------------------
-- id_cat: 1=Limpieza | id_tipo: producto_tienda
INSERT INTO tienda.productos
    (id_categorias_tienda, nombre, descripcion, id_tipo,
     precio, precio_original, descuento_pct, stock, es_nuevo, destacado)
VALUES
    (1, 'Desengrasante multiusos 1L',
        'Limpiador potente para cocinas y superficies grasosas.',
        (SELECT id_tipo FROM catalogo.tipos WHERE contexto='producto_tienda' AND nombre='insumos_aseo'),
        18900.00, 22000.00, 14, 150, FALSE, TRUE),

    (2, 'Aspiradora portátil Jobsy Pro',
        'Aspiradora inalámbrica liviana, ideal para escaleras y sofás.',
        (SELECT id_tipo FROM catalogo.tipos WHERE contexto='producto_tienda' AND nombre='equipo'),
        189000.00, 210000.00, 10, 30, TRUE, TRUE),

    (1, 'Kit guantes de nitrilo x3 pares',
        'Guantes desechables resistentes para limpieza profunda.',
        (SELECT id_tipo FROM catalogo.tipos WHERE contexto='producto_tienda' AND nombre='proteccion'),
        12500.00, NULL, 0, 200, FALSE, FALSE);

-- ---- tienda.cupones ------------------------------------------
INSERT INTO tienda.cupones
    (codigo, descuento_pct, usos_maximos, expira_en)
VALUES
    ('BIENVENIDO10', 10, 500, NOW() + INTERVAL '60 days'),
    ('VERANO20',     20, 200, NOW() + INTERVAL '30 days'),
    ('JOBSYPRO15',   15, 100, NOW() + INTERVAL '90 days');

-- ---- tienda.carritos -----------------------------------------
INSERT INTO tienda.carritos (id_usuarios, id_estado, id_cupones)
VALUES
    (1, (SELECT id_estado FROM catalogo.estados WHERE contexto='carrito' AND nombre='activo'),    NULL),
    (2, (SELECT id_estado FROM catalogo.estados WHERE contexto='carrito' AND nombre='activo'),    1),
    (3, (SELECT id_estado FROM catalogo.estados WHERE contexto='carrito' AND nombre='pagado'),    NULL);

-- ---- tienda.TR_carrito_items ---------------------------------
INSERT INTO tienda.TR_carrito_items
    (id_carritos, id_productos, cantidad, precio_unitario)
VALUES
    (1, 1, 2, 18900.00),
    (2, 2, 1, 189000.00),
    (3, 3, 3, 12500.00);

-- ---- tienda.facturacion --------------------------------------
INSERT INTO tienda.facturacion
    (id_usuarios, id_carritos, id_estado,
     subtotal, descuento_total, total,
     factura_nombre, factura_doc, factura_correo, factura_ciudad)
VALUES
    (1, 1,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='pedido' AND nombre='pendiente'),
     37800.00,  0.00,     37800.00,  'Laura Martínez',  '1020304050', 'laura.martinez@email.com', 'Bogotá'),
    (2, 2,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='pedido' AND nombre='pagado'),
     189000.00, 18900.00, 170100.00, 'Carlos Rincón',   '1098765432', 'carlos.rincon@email.com',  'Medellín'),
    (3, 3,
     (SELECT id_estado FROM catalogo.estados WHERE contexto='pedido' AND nombre='pagado'),
     37500.00,  0.00,     37500.00,  'Valentina Gómez', '1044556677', 'vgomez@email.com',         'Cali');

-- ---- tienda.pedido_items -------------------------------------
INSERT INTO tienda.pedido_items
    (id_pedidos, id_productos, cantidad, precio_unitario, subtotal)
VALUES
    (1, 1, 2, 18900.00,  37800.00),
    (2, 2, 1, 189000.00, 189000.00),
    (3, 3, 3, 12500.00,  37500.00);

-- ---- tienda.pagos --------------------------------------------
-- id_tipo_pago contexto='pago' | id_estado contexto='pago'
INSERT INTO tienda.pagos
    (id_pedidos, id_usuarios, id_metodo_pago, id_tipo_pago, id_estado, monto)
VALUES
    (1, 1, 3,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='pago' AND nombre='tienda'),
     (SELECT id_estado FROM catalogo.estados WHERE contexto='pago' AND nombre='pendiente'),
     37800.00),
    (2, 2, 3,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='pago' AND nombre='tienda'),
     (SELECT id_estado FROM catalogo.estados WHERE contexto='pago' AND nombre='aprobado'),
     170100.00),
    (3, 3, 1,
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='pago' AND nombre='tienda'),
     (SELECT id_estado FROM catalogo.estados WHERE contexto='pago' AND nombre='aprobado'),
     37500.00);

-- ---- tienda.transacciones ------------------------------------
-- id_tipo_transaccion: 1=ingreso 2=egreso
INSERT INTO tienda.transacciones
    (id_usuarios, id_pagos, id_tipo_transaccion, monto, descripcion)
VALUES
    (2, 2, 1, 170100.00, 'Compra tienda - Aspiradora Jobsy Pro'),
    (3, 3, 1, 37500.00,  'Compra tienda - Kit guantes nitrilo x3'),
    (1, 1, 2, 37800.00,  'Pago pendiente - Desengrasante multiusos');

-- ---- tienda.saldo_jobsy --------------------------------------
INSERT INTO tienda.saldo_jobsy
    (id_usuarios, saldo_disponible, total_ganado, total_retirado, proximo_pago)
VALUES
    (1, 80000.00,  160000.00, 80000.00,  '2025-06-01'),
    (3, 108000.00, 216000.00, 108000.00, '2025-06-01'),
    (2, 0.00,      0.00,      0.00,      NULL);

-- ---- tienda.metodos_pago_guardados ---------------------------
-- id_metodo_pago: 3=nequi 4=daviplata | id_tipo_cuenta: contexto='cuenta_bancaria'
INSERT INTO tienda.metodos_pago_guardados
    (id_usuarios, id_metodo_pago, alias, numero_cuenta, id_tipo_cuenta, es_principal)
VALUES
    (1, 3, 'Mi Nequi',         '3101234567',
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='cuenta_bancaria' AND nombre='ahorros'), TRUE),
    (2, 4, 'Daviplata Carlos', '3209876543',
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='cuenta_bancaria' AND nombre='ahorros'), TRUE),
    (3, 3, 'Nequi Valentina',  '3156784321',
     (SELECT id_tipo FROM catalogo.tipos WHERE contexto='cuenta_bancaria' AND nombre='ahorros'), TRUE);
