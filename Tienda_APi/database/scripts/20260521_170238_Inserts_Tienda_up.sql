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
