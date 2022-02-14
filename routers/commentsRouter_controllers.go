package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["Medicina/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["Medicina/controllers:AntecedenteController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["Medicina/controllers:AntecedenteController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["Medicina/controllers:AntecedenteController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["Medicina/controllers:AntecedenteController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:AntecedenteController"] = append(beego.GlobalControllerRouter["Medicina/controllers:AntecedenteController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["Medicina/controllers:EspecialidadController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	
	beego.GlobalControllerRouter["Medicina/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["Medicina/controllers:EspecialidadController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	
	beego.GlobalControllerRouter["Medicina/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["Medicina/controllers:EspecialidadController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	
	beego.GlobalControllerRouter["Medicina/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["Medicina/controllers:EspecialidadController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	
	beego.GlobalControllerRouter["Medicina/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["Medicina/controllers:EspecialidadController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:ConsultaFisioterapiaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:ConsultaFisioterapiaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:ConsultaFisioterapiaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:ConsultaFisioterapiaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:ConsultaFisioterapiaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:ConsultaFisioterapiaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:ConsultaFisioterapiaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:ConsultaFisioterapiaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:ConsultaFisioterapiaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:ConsultaFisioterapiaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Medicina/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Medicina/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Medicina/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Medicina/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Medicina/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:ExamenController"] = append(beego.GlobalControllerRouter["Medicina/controllers:ExamenController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:ExamenController"] = append(beego.GlobalControllerRouter["Medicina/controllers:ExamenController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:ExamenController"] = append(beego.GlobalControllerRouter["Medicina/controllers:ExamenController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:ExamenController"] = append(beego.GlobalControllerRouter["Medicina/controllers:ExamenController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:ExamenController"] = append(beego.GlobalControllerRouter["Medicina/controllers:ExamenController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:HistoriaClinicaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:HistoriaClinicaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:HistoriaClinicaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:HistoriaClinicaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:HistoriaClinicaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:HistoriaClinicaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:HistoriaClinicaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:HistoriaClinicaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:HistoriaClinicaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:HistoriaClinicaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:HojaHistoriaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:HojaHistoriaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:HojaHistoriaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:HojaHistoriaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:HojaHistoriaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:HojaHistoriaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:HojaHistoriaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:HojaHistoriaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:HojaHistoriaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:HojaHistoriaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	
	beego.GlobalControllerRouter["Medicina/controllers:NotasEnfermeriaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:NotasEnfermeriaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	
	beego.GlobalControllerRouter["Medicina/controllers:NotasEnfermeriaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:NotasEnfermeriaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	
	beego.GlobalControllerRouter["Medicina/controllers:NotasEnfermeriaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:NotasEnfermeriaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	
	beego.GlobalControllerRouter["Medicina/controllers:NotasEnfermeriaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:NotasEnfermeriaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	
	beego.GlobalControllerRouter["Medicina/controllers:NotasEnfermeriaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:NotasEnfermeriaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:SistemaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:SistemaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:SistemaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:SistemaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:SistemaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:SistemaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:SistemaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:SistemaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Medicina/controllers:SistemaController"] = append(beego.GlobalControllerRouter["Medicina/controllers:SistemaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}