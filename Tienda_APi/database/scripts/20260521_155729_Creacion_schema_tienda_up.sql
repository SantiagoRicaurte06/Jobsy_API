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
