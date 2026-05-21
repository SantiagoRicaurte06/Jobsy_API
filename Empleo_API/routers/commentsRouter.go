package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["Api_Josby/controllers:CertificacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:CertificacionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:CertificacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:CertificacionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:CertificacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:CertificacionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:CertificacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:CertificacionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:CertificacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:CertificacionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:ContratacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:ContratacionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:ContratacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:ContratacionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:ContratacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:ContratacionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:ContratacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:ContratacionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:ContratacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:ContratacionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DiasTrabajadosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DiasTrabajadosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DiasTrabajadosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DiasTrabajadosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DiasTrabajadosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DiasTrabajadosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DiasTrabajadosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DiasTrabajadosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DiasTrabajadosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DiasTrabajadosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DisponibilidadEmpleadoController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DisponibilidadEmpleadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DisponibilidadEmpleadoController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DisponibilidadEmpleadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DisponibilidadEmpleadoController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DisponibilidadEmpleadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DisponibilidadEmpleadoController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DisponibilidadEmpleadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DisponibilidadEmpleadoController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DisponibilidadEmpleadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DocumentosUsuarioController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DocumentosUsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DocumentosUsuarioController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DocumentosUsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DocumentosUsuarioController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DocumentosUsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DocumentosUsuarioController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DocumentosUsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:DocumentosUsuarioController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:DocumentosUsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:EspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:EspecialidadesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:EspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:EspecialidadesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:EspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:EspecialidadesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:EspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:EspecialidadesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:EspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:EspecialidadesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HojasVidaController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HojasVidaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HojasVidaController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HojasVidaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HojasVidaController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HojasVidaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HojasVidaController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HojasVidaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HojasVidaController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HojasVidaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HorariosAgendadosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HorariosAgendadosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HorariosAgendadosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HorariosAgendadosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HorariosAgendadosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HorariosAgendadosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HorariosAgendadosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HorariosAgendadosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HorariosAgendadosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HorariosAgendadosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HvExperienciasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HvExperienciasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HvExperienciasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HvExperienciasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HvExperienciasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HvExperienciasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HvExperienciasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HvExperienciasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:HvExperienciasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:HvExperienciasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:OfertaFotosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:OfertaFotosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:OfertaFotosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:OfertaFotosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:OfertaFotosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:OfertaFotosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:OfertaFotosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:OfertaFotosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:OfertaFotosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:OfertaFotosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:OfertasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:OfertasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:OfertasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:OfertasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:OfertasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:OfertasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:OfertasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:OfertasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:OfertasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:OfertasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PerfilesEmpleadoController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PerfilesEmpleadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PerfilesEmpleadoController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PerfilesEmpleadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PerfilesEmpleadoController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PerfilesEmpleadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PerfilesEmpleadoController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PerfilesEmpleadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PerfilesEmpleadoController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PerfilesEmpleadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionChipsController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionChipsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionChipsController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionChipsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionChipsController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionChipsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionChipsController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionChipsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionChipsController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionChipsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:PostulacionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:ResenasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:ResenasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:ResenasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:ResenasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:ResenasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:ResenasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:ResenasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:ResenasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:ResenasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:ResenasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TamanoInmuebleController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TamanoInmuebleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TamanoInmuebleController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TamanoInmuebleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TamanoInmuebleController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TamanoInmuebleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TamanoInmuebleController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TamanoInmuebleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TamanoInmuebleController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TamanoInmuebleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TipoInmuebleController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TipoInmuebleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TipoInmuebleController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TipoInmuebleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TipoInmuebleController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TipoInmuebleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TipoInmuebleController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TipoInmuebleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TipoInmuebleController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TipoInmuebleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrContratacionEspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrContratacionEspecialidadesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrContratacionEspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrContratacionEspecialidadesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrContratacionEspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrContratacionEspecialidadesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrContratacionEspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrContratacionEspecialidadesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrContratacionEspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrContratacionEspecialidadesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrOfertaServiciosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrOfertaServiciosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrOfertaServiciosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrOfertaServiciosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrOfertaServiciosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrOfertaServiciosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrOfertaServiciosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrOfertaServiciosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrOfertaServiciosController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrOfertaServiciosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrPerfilEspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrPerfilEspecialidadesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrPerfilEspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrPerfilEspecialidadesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrPerfilEspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrPerfilEspecialidadesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrPerfilEspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrPerfilEspecialidadesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrPerfilEspecialidadesController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrPerfilEspecialidadesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrPostulacionDiasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrPostulacionDiasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrPostulacionDiasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrPostulacionDiasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrPostulacionDiasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrPostulacionDiasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrPostulacionDiasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrPostulacionDiasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Api_Josby/controllers:TrPostulacionDiasController"] = append(beego.GlobalControllerRouter["Api_Josby/controllers:TrPostulacionDiasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
