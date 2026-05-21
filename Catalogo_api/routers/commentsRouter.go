package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:DiasSemanaController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:DiasSemanaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:DiasSemanaController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:DiasSemanaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:DiasSemanaController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:DiasSemanaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:DiasSemanaController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:DiasSemanaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:DiasSemanaController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:DiasSemanaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:EstadosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:EstadosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:EstadosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:EstadosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:EstadosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:EstadosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:EstadosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:EstadosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:EstadosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:EstadosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:GenerosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:GenerosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:GenerosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:GenerosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:GenerosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:GenerosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:GenerosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:GenerosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:GenerosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:GenerosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesSuscripcionController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesSuscripcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesSuscripcionController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesSuscripcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesSuscripcionController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesSuscripcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesSuscripcionController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesSuscripcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesSuscripcionController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesSuscripcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesTrabajoController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesTrabajoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesTrabajoController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesTrabajoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesTrabajoController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesTrabajoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesTrabajoController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesTrabajoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesTrabajoController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:ModalidadesTrabajoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:NivelesExperienciaController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:NivelesExperienciaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:NivelesExperienciaController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:NivelesExperienciaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:NivelesExperienciaController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:NivelesExperienciaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:NivelesExperienciaController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:NivelesExperienciaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:NivelesExperienciaController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:NivelesExperienciaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TiposController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TiposController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TiposController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TiposController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TiposController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TiposController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TiposController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TiposController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TiposController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TiposController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TurnosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TurnosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TurnosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TurnosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TurnosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TurnosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TurnosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TurnosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TurnosController"] = append(beego.GlobalControllerRouter["Api_Josby/Catalogo_API/controllers:TurnosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
