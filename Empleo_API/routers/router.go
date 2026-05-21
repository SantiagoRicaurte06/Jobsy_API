// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Api_Josby/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

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

		beego.NSNamespace("/documentos_usuario",
			beego.NSInclude(
				&controllers.DocumentosUsuarioController{},
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

	)
	beego.AddNamespace(ns)
}
