package routers

import (
	"Jobsy_API/Core_api/controllers"

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

		beego.NSNamespace("/tipo_documento",
			beego.NSInclude(
				&controllers.TipoDocumentoController{},
			),
		),

		beego.NSNamespace("/documentos_identidad",
			beego.NSInclude(
				&controllers.DocumentosIdentidadController{},
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
	)

	beego.AddNamespace(ns)
}
