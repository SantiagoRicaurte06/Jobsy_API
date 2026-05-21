package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:DocumentosIdentidadController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:DocumentosIdentidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:DocumentosIdentidadController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:DocumentosIdentidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:DocumentosIdentidadController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:DocumentosIdentidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:DocumentosIdentidadController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:DocumentosIdentidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:DocumentosIdentidadController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:DocumentosIdentidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:RolesUsuariosController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:RolesUsuariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:RolesUsuariosController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:RolesUsuariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:RolesUsuariosController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:RolesUsuariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:RolesUsuariosController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:RolesUsuariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:RolesUsuariosController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:RolesUsuariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:SesionesController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:SesionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:SesionesController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:SesionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:SesionesController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:SesionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:SesionesController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:SesionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:SesionesController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:SesionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:TipoDocumentoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["Jobsy_API/Core_api/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
