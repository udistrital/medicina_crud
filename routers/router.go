// @APIVersion 1.0
// @Title Medicina
// @Description Api para Medicina dentro del módulo de Salud del proyecto SIBUD
// @Contact baluist@correo.udistrital.edu.co
// @TermsOfServiceUrl http://www.udistrital.edu.co/politicas-de-privacidad.pdf
// @License BSD-3-Clause
// @LicenseUrl http://opensource.org/licenses/BSD-3-Clause
// @BasePath /Medicina/v1
package routers

import (
	"Medicina/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/Antecedente",
			beego.NSInclude(
				&controllers.AntecedenteController{},
			),
		),
		beego.NSNamespace("/Especialidad",
			beego.NSInclude(
				&controllers.EspecialidadController{},
			),
		),
		beego.NSNamespace("/ConsultaFisioterapia",
			beego.NSInclude(
				&controllers.ConsultaFisioterapiaController{},
			),
		),
		beego.NSNamespace("/Diagnostico",
			beego.NSInclude(
				&controllers.DiagnosticoController{},
			),
		),
		beego.NSNamespace("/Examen",
			beego.NSInclude(
				&controllers.ExamenController{},
			),
		), beego.NSNamespace("/HistoriaClinica",
			beego.NSInclude(
				&controllers.HistoriaClinicaController{},
			),
		), beego.NSNamespace("/HojaHistoria",
			beego.NSInclude(
				&controllers.HojaHistoriaController{},
			),
		),
		beego.NSNamespace("/NotasEnfermeria",
			beego.NSInclude(
				&controllers.NotasEnfermeriaController{},
			),
		), 
		beego.NSNamespace("/Sistema",
			beego.NSInclude(
				&controllers.SistemaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
