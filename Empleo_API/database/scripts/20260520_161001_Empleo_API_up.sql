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