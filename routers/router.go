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
		beego.NSNamespace("/medicinaAntecedente",
			beego.NSInclude(
				&controllers.MedicinaAntecedenteController{},
			),
		),
		beego.NSNamespace("/medicinaConsultaFisioterapia",
			beego.NSInclude(
				&controllers.MedicinaConsultaFisioterapiaController{},
			),
		),
		beego.NSNamespace("/medicinaDiagnostico",
			beego.NSInclude(
				&controllers.MedicinaDiagnosticoController{},
			),
		),
		beego.NSNamespace("/medicinaExamen",
			beego.NSInclude(
				&controllers.MedicinaExamenController{},
			),
		), beego.NSNamespace("/medicinaHistoriaClinica",
			beego.NSInclude(
				&controllers.MedicinaHistoriaClinicaController{},
			),
		), beego.NSNamespace("/medicinaHojaHistoria",
			beego.NSInclude(
				&controllers.MedicinaHojaHistoriaController{},
			),
		), beego.NSNamespace("/medicinaSistemas",
			beego.NSInclude(
				&controllers.MedicinaSistemasController{},
			),
		),
		beego.NSNamespace("/medicinaTipoAntecedente",
			beego.NSInclude(
				&controllers.MedicinaTipoAntecedenteController{},
			),
		),
		beego.NSNamespace("/medicinaTipoExamen",
			beego.NSInclude(
				&controllers.MedicinaTipoExamenController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
