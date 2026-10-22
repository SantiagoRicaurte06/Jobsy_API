-- Deshace los inserts de semilla de tienda.
--
-- La version anterior borraba una sola fila por tabla (WHERE id = 1) usando
-- cuatro nombres de columna que no existen: id_carrito_item,
-- id_pedido_item, id_saldo y id_metodo_guardado. Ademas el orden violaba las
-- claves foraneas.
--
-- Ahora borra todas las filas sembradas, de hija a madre.

DELETE FROM tienda.transacciones;
DELETE FROM tienda.pagos;
DELETE FROM tienda.pedido_items;
DELETE FROM tienda.facturacion;
DELETE FROM tienda.tr_carrito_items;
DELETE FROM tienda.carritos;
DELETE FROM tienda.metodos_pago_guardados;
DELETE FROM tienda.saldo_jobsy;
DELETE FROM tienda.cupones;
DELETE FROM tienda.productos;
DELETE FROM tienda.categorias_tienda;
DELETE FROM tienda.tipo_transaccion;
DELETE FROM tienda.metodo_pago;
