package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanFeaturesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanFeaturesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanFeaturesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanFeaturesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanFeaturesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanFeaturesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanFeaturesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanFeaturesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanFeaturesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanFeaturesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:PlanesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:UsuariosSuscripcionesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:UsuariosSuscripcionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:UsuariosSuscripcionesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:UsuariosSuscripcionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:UsuariosSuscripcionesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:UsuariosSuscripcionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:UsuariosSuscripcionesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:UsuariosSuscripcionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Suscripciones_Api/controllers:UsuariosSuscripcionesController"] = append(beego.GlobalControllerRouter["Suscripciones_Api/controllers:UsuariosSuscripcionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
