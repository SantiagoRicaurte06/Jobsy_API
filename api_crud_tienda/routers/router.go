// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api_crud_tienda/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/categorias_tienda",
			beego.NSInclude(
				&controllers.CategoriasTiendaController{},
			),
		),

		beego.NSNamespace("/productos",
			beego.NSInclude(
				&controllers.ProductosController{},
			),
		),

		beego.NSNamespace("/carritos",
			beego.NSInclude(
				&controllers.CarritosController{},
			),
		),

		beego.NSNamespace("/cupones",
			beego.NSInclude(
				&controllers.CuponesController{},
			),
		),

		beego.NSNamespace("/tr_carrito_items",
			beego.NSInclude(
				&controllers.TrCarritoItemsController{},
			),
		),

		beego.NSNamespace("/facturacion",
			beego.NSInclude(
				&controllers.FacturacionController{},
			),
		),

		beego.NSNamespace("/pedido_items",
			beego.NSInclude(
				&controllers.PedidoItemsController{},
			),
		),

		beego.NSNamespace("/pagos",
			beego.NSInclude(
				&controllers.PagosController{},
			),
		),

		beego.NSNamespace("/metodo_pago",
			beego.NSInclude(
				&controllers.MetodoPagoController{},
			),
		),

		beego.NSNamespace("/transacciones",
			beego.NSInclude(
				&controllers.TransaccionesController{},
			),
		),

		beego.NSNamespace("/tipo_transaccion",
			beego.NSInclude(
				&controllers.TipoTransaccionController{},
			),
		),

		beego.NSNamespace("/saldo_jobsy",
			beego.NSInclude(
				&controllers.SaldoJobsyController{},
			),
		),

		beego.NSNamespace("/metodos_pago_guardados",
			beego.NSInclude(
				&controllers.MetodosPagoGuardadosController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
