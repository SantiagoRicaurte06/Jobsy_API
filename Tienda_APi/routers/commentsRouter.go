package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CarritosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CarritosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CarritosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CarritosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CarritosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CarritosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CarritosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CarritosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CarritosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CarritosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CategoriasTiendaController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CategoriasTiendaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CategoriasTiendaController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CategoriasTiendaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CategoriasTiendaController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CategoriasTiendaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CategoriasTiendaController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CategoriasTiendaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CategoriasTiendaController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CategoriasTiendaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CuponesController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CuponesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CuponesController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CuponesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CuponesController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CuponesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CuponesController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CuponesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:CuponesController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:CuponesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:FacturacionController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:FacturacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:FacturacionController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:FacturacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:FacturacionController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:FacturacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:FacturacionController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:FacturacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:FacturacionController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:FacturacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodoPagoController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodoPagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodoPagoController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodoPagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodoPagoController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodoPagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodoPagoController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodoPagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodoPagoController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodoPagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodosPagoGuardadosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodosPagoGuardadosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodosPagoGuardadosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodosPagoGuardadosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodosPagoGuardadosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodosPagoGuardadosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodosPagoGuardadosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodosPagoGuardadosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodosPagoGuardadosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:MetodosPagoGuardadosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:PagosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:PagosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:PagosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:PagosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:PagosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:PagosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:PagosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:PagosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:PagosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:PagosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:PedidoItemsController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:PedidoItemsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:PedidoItemsController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:PedidoItemsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:PedidoItemsController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:PedidoItemsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:PedidoItemsController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:PedidoItemsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:PedidoItemsController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:PedidoItemsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:ProductosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:ProductosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:ProductosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:ProductosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:ProductosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:ProductosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:ProductosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:ProductosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:ProductosController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:ProductosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:SaldoJobsyController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:SaldoJobsyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:SaldoJobsyController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:SaldoJobsyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:SaldoJobsyController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:SaldoJobsyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:SaldoJobsyController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:SaldoJobsyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:SaldoJobsyController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:SaldoJobsyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TipoTransaccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TipoTransaccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TipoTransaccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TipoTransaccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TipoTransaccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TrCarritoItemsController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TrCarritoItemsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TrCarritoItemsController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TrCarritoItemsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TrCarritoItemsController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TrCarritoItemsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TrCarritoItemsController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TrCarritoItemsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TrCarritoItemsController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TrCarritoItemsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TransaccionesController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TransaccionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TransaccionesController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TransaccionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TransaccionesController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TransaccionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TransaccionesController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TransaccionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_tienda/controllers:TransaccionesController"] = append(beego.GlobalControllerRouter["api_crud_tienda/controllers:TransaccionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
