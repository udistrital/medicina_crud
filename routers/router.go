// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
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
		), beego.NSNamespace("/Sistemas",
			beego.NSInclude(
				&controllers.SistemasController{},
			),
		),
		beego.NSNamespace("/TipoAntecedente",
			beego.NSInclude(
				&controllers.TipoAntecedenteController{},
			),
		),
		beego.NSNamespace("/TipoExamen",
			beego.NSInclude(
				&controllers.TipoExamenController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
