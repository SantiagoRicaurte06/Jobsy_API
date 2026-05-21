// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Suscripciones_Api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/roles_usuarios",
			beego.NSInclude(
				&controllers.RolesUsuariosController{},
			),
		),

		beego.NSNamespace("/usuarios",
			beego.NSInclude(
				&controllers.UsuariosController{},
			),
		),

		beego.NSNamespace("/generos",
			beego.NSInclude(
				&controllers.GenerosController{},
			),
		),

		beego.NSNamespace("/documentos_identidad",
			beego.NSInclude(
				&controllers.DocumentosIdentidadController{},
			),
		),

		beego.NSNamespace("/tipo_documento",
			beego.NSInclude(
				&controllers.TipoDocumentoController{},
			),
		),

		beego.NSNamespace("/credenciales",
			beego.NSInclude(
				&controllers.CredencialesController{},
			),
		),

		beego.NSNamespace("/sesiones",
			beego.NSInclude(
				&controllers.SesionesController{},
			),
		),

		beego.NSNamespace("/dias_semana",
			beego.NSInclude(
				&controllers.DiasSemanaController{},
			),
		),

		beego.NSNamespace("/dias_trabajados",
			beego.NSInclude(
				&controllers.DiasTrabajadosController{},
			),
		),

		beego.NSNamespace("/perfiles_empleado",
			beego.NSInclude(
				&controllers.PerfilesEmpleadoController{},
			),
		),

		beego.NSNamespace("/niveles_experiencia",
			beego.NSInclude(
				&controllers.NivelesExperienciaController{},
			),
		),

		beego.NSNamespace("/modalidades_trabajo",
			beego.NSInclude(
				&controllers.ModalidadesTrabajoController{},
			),
		),

		beego.NSNamespace("/turnos",
			beego.NSInclude(
				&controllers.TurnosController{},
			),
		),

		beego.NSNamespace("/tr_perfil_especialidades",
			beego.NSInclude(
				&controllers.TrPerfilEspecialidadesController{},
			),
		),

		beego.NSNamespace("/especialidades",
			beego.NSInclude(
				&controllers.EspecialidadesController{},
			),
		),

		beego.NSNamespace("/disponibilidad_empleado",
			beego.NSInclude(
				&controllers.DisponibilidadEmpleadoController{},
			),
		),

		beego.NSNamespace("/certificaciones",
			beego.NSInclude(
				&controllers.CertificacionesController{},
			),
		),

		beego.NSNamespace("/estados",
			beego.NSInclude(
				&controllers.EstadosController{},
			),
		),

		beego.NSNamespace("/documentos_usuario",
			beego.NSInclude(
				&controllers.DocumentosUsuarioController{},
			),
		),

		beego.NSNamespace("/tipos",
			beego.NSInclude(
				&controllers.TiposController{},
			),
		),

		beego.NSNamespace("/ofertas",
			beego.NSInclude(
				&controllers.OfertasController{},
			),
		),

		beego.NSNamespace("/tipo_inmueble",
			beego.NSInclude(
				&controllers.TipoInmuebleController{},
			),
		),

		beego.NSNamespace("/tamano_inmueble",
			beego.NSInclude(
				&controllers.TamanoInmuebleController{},
			),
		),

		beego.NSNamespace("/tr_oferta_servicios",
			beego.NSInclude(
				&controllers.TrOfertaServiciosController{},
			),
		),

		beego.NSNamespace("/oferta_fotos",
			beego.NSInclude(
				&controllers.OfertaFotosController{},
			),
		),

		beego.NSNamespace("/hojas_vida",
			beego.NSInclude(
				&controllers.HojasVidaController{},
			),
		),

		beego.NSNamespace("/hv_experiencias",
			beego.NSInclude(
				&controllers.HvExperienciasController{},
			),
		),

		beego.NSNamespace("/postulaciones",
			beego.NSInclude(
				&controllers.PostulacionesController{},
			),
		),

		beego.NSNamespace("/tr_postulacion_dias",
			beego.NSInclude(
				&controllers.TrPostulacionDiasController{},
			),
		),

		beego.NSNamespace("/postulacion_chips",
			beego.NSInclude(
				&controllers.PostulacionChipsController{},
			),
		),

		beego.NSNamespace("/resenas",
			beego.NSInclude(
				&controllers.ResenasController{},
			),
		),

		beego.NSNamespace("/contrataciones",
			beego.NSInclude(
				&controllers.ContratacionesController{},
			),
		),

		beego.NSNamespace("/tr_contratacion_especialidades",
			beego.NSInclude(
				&controllers.TrContratacionEspecialidadesController{},
			),
		),

		beego.NSNamespace("/horarios_agendados",
			beego.NSInclude(
				&controllers.HorariosAgendadosController{},
			),
		),

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

		beego.NSNamespace("/planes",
			beego.NSInclude(
				&controllers.PlanesController{},
			),
		),

		beego.NSNamespace("/plan_features",
			beego.NSInclude(
				&controllers.PlanFeaturesController{},
			),
		),

		beego.NSNamespace("/usuarios_suscripciones",
			beego.NSInclude(
				&controllers.UsuariosSuscripcionesController{},
			),
		),

		beego.NSNamespace("/modalidades_suscripcion",
			beego.NSInclude(
				&controllers.ModalidadesSuscripcionController{},
			),
		),

		beego.NSNamespace("/configuraciones",
			beego.NSInclude(
				&controllers.ConfiguracionesController{},
			),
		),

		beego.NSNamespace("/notificaciones",
			beego.NSInclude(
				&controllers.NotificacionesController{},
			),
		),

		beego.NSNamespace("/reportes",
			beego.NSInclude(
				&controllers.ReportesController{},
			),
		),

		beego.NSNamespace("/reporte_mensajes",
			beego.NSInclude(
				&controllers.ReporteMensajesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
