-- =====================================================================
--  Ajuste de secuencias tras la semilla
--
--  01_jobsy.sql inserta filas con ids explicitos (1, 2, 3...) y eso no
--  avanza las secuencias SERIAL: el primer INSERT desde una API intentaba
--  usar el id 1 y fallaba con "duplicate key value violates unique
--  constraint ..._pkey". Esto pone cada secuencia en el maximo id de su tabla.
-- =====================================================================

DO $$
DECLARE
    c record;
BEGIN
    FOR c IN
        SELECT n.nspname AS esquema, t.relname AS tabla, a.attname AS columna,
               pg_get_serial_sequence(format('%I.%I', n.nspname, t.relname), a.attname) AS secuencia
        FROM pg_class t
        JOIN pg_namespace n ON n.oid = t.relnamespace
        JOIN pg_attribute a ON a.attrelid = t.oid AND a.attnum > 0 AND NOT a.attisdropped
        WHERE t.relkind = 'r'
          AND n.nspname IN ('catalogo', 'core', 'empleo', 'soporte', 'suscripciones', 'tienda', 'usuarios')
          AND pg_get_serial_sequence(format('%I.%I', n.nspname, t.relname), a.attname) IS NOT NULL
    LOOP
        EXECUTE format(
            'SELECT setval(%L, COALESCE((SELECT MAX(%I) FROM %I.%I), 0) + 1, false)',
            c.secuencia, c.columna, c.esquema, c.tabla);
    END LOOP;
END $$;
