-- ==============================
-- DELETE POR ID - tienda.metodo_pago
-- ==============================
DELETE FROM tienda.metodo_pago
WHERE id_metodo_pago = 1;

-- ==============================
-- DELETE POR ID - tienda.tipo_transaccion
-- ==============================
DELETE FROM tienda.tipo_transaccion
WHERE id_tipo_transaccion = 1;

-- ==============================
-- DELETE POR ID - tienda.categorias_tienda
-- ==============================
DELETE FROM tienda.categorias_tienda
WHERE id_categorias_tienda = 1;

-- ==============================
-- DELETE POR ID - tienda.productos
-- ==============================
DELETE FROM tienda.productos
WHERE id_productos = 1;

-- ==============================
-- DELETE POR ID - tienda.cupones
-- ==============================
DELETE FROM tienda.cupones
WHERE id_cupones = 1;

-- ==============================
-- DELETE POR ID - tienda.carritos
-- ==============================
DELETE FROM tienda.carritos
WHERE id_carritos = 1;

-- ==============================
-- DELETE POR ID - tienda.TR_carrito_items
-- ==============================
DELETE FROM tienda.TR_carrito_items
WHERE id_carrito_item = 1;

-- ==============================
-- DELETE POR ID - tienda.facturacion
-- ==============================
DELETE FROM tienda.facturacion
WHERE id_pedidos = 1;

-- ==============================
-- DELETE POR ID - tienda.pedido_items
-- ==============================
DELETE FROM tienda.pedido_items
WHERE id_pedido_item = 1;

-- ==============================
-- DELETE POR ID - tienda.pagos
-- ==============================
DELETE FROM tienda.pagos
WHERE id_pagos = 1;

-- ==============================
-- DELETE POR ID - tienda.transacciones
-- ==============================
DELETE FROM tienda.transacciones
WHERE id_transacciones = 1;

-- ==============================
-- DELETE POR ID - tienda.saldo_jobsy
-- ==============================
DELETE FROM tienda.saldo_jobsy
WHERE id_saldo = 1;

-- ==============================
-- DELETE POR ID - tienda.metodos_pago_guardados
-- ==============================
DELETE FROM tienda.metodos_pago_guardados
WHERE id_metodo_guardado = 1;