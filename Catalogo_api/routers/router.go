// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Api_Josby/Catalogo_API/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		// Generos
		beego.NSNamespace("/generos",
			beego.NSInclude(
				&controllers.GenerosController{},
			),
		),

		// Modalidades de Trabajo
		beego.NSNamespace("/modalidades_trabajo",
			beego.NSInclude(
				&controllers.ModalidadesTrabajoController{},
			),
		),

		// Modalidades de Suscripcion
		beego.NSNamespace("/modalidades_suscripcion",
			beego.NSInclude(
				&controllers.ModalidadesSuscripcionController{},
			),
		),

		// Niveles de Experiencia
		beego.NSNamespace("/niveles_experiencia",
			beego.NSInclude(
				&controllers.NivelesExperienciaController{},
			),
		),

		// Turnos
		beego.NSNamespace("/turnos",
			beego.NSInclude(
				&controllers.TurnosController{},
			),
		),

		// Dias de la Semana
		beego.NSNamespace("/dias_semana",
			beego.NSInclude(
				&controllers.DiasSemanaController{},
			),
		),

		// Estados
		beego.NSNamespace("/estados",
			beego.NSInclude(
				&controllers.EstadosController{},
			),
		),

		// Tipos
		beego.NSNamespace("/tipos",
			beego.NSInclude(
				&controllers.TiposController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
