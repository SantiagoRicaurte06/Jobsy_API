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
S
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