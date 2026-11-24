#!/bin/sh
# Crea el usuario con el que se conectan las 7 APIs. Solo lee y escribe
# datos: no puede crear ni borrar tablas (las APIs no lo necesitan, el
# esquema sale de las migraciones). Postgres corre este script una sola vez,
# al crear el volumen, despues de los .sql de esta carpeta.
set -e

psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d "$POSTGRES_DB" \
     -v usuario="$JOBSY_DB_USUARIO" -v clave="$JOBSY_DB_CLAVE" <<'SQL'
CREATE ROLE :"usuario" LOGIN PASSWORD :'clave';
-- psql no sustituye variables dentro de $$...$$: el nombre pasa por un GUC.
SELECT set_config('jobsy.usuario', :'usuario', false);

DO $$
DECLARE
    e text;
BEGIN
    FOREACH e IN ARRAY ARRAY['catalogo','core','empleo','soporte','suscripciones','tienda','usuarios'] LOOP
        EXECUTE format('GRANT USAGE ON SCHEMA %I TO %I', e, current_setting('jobsy.usuario'));
        EXECUTE format('GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA %I TO %I', e, current_setting('jobsy.usuario'));
        EXECUTE format('GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA %I TO %I', e, current_setting('jobsy.usuario'));
        -- Las tablas que creen migraciones futuras (como postgres) tambien.
        EXECUTE format('ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO %I', e, current_setting('jobsy.usuario'));
        EXECUTE format('ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO %I', e, current_setting('jobsy.usuario'));
    END LOOP;
END
$$;
SQL
