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

	)
	beego.AddNamespace(ns)
}
