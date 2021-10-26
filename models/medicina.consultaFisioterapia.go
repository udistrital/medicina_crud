package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MedicinaConsultaFisioterapia struct {
	IdMedicinaConsultaFisioterapia int    `orm:"column(id_consulta_fisioterapia);pk;auto"`
	IdHojaHistoria                 *int   `orm:"column(id_hoja_historia);rel(fk);null"`
	Motivo_consulta                string `orm:"column(motivo_consulta);null"`
	Valoracion                     string `orm:"column(valoracion);null"`
	PlanManejo                     string `orm:"column(plan_manejo);null"`
	Evolucion                      string `orm:"column(evolucion);null"`
	Observaciones                  string `orm:"column(observaciones);null"`
}

func (t *MedicinaConsultaFisioterapia) TableName() string {
	return "consultafisioterapia"
}
func init() {
	orm.RegisterModel(new(MedicinaConsultaFisioterapia))
}

// AddMedicinaConsultaFisioterapia inserta un registro en la tabla consultafisioterapia
// Último registro insertado con éxito
func AddMedicinaConsultaFisioterapia(m *MedicinaConsultaFisioterapia) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMedicinaConsultaFisioterapiaById obtiene un registro de la tabla consultafisioterapia por su id
// Id no existe
func GetMedicinaConsultaFisioterapiaById(id int) (v *MedicinaConsultaFisioterapia, err error) {
	o := orm.NewOrm()
	v = &MedicinaConsultaFisioterapia{IdMedicinaConsultaFisioterapia: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMedicinaConsultaFisioterapia obtiene todos los registros de la tabla consultafisioterapia
// No existen registros
func GetAllMedicinaConsultaFisioterapia(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MedicinaConsultaFisioterapia))
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Orden inválido, debe ser del tipo [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Orden inválido, debe ser del tipo [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("error: los tamaños de 'sortby', 'order' no coinciden o el tamaño de 'order' no es 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("error: campos de 'order' no utilizados")
		}
	}
	var l []MedicinaConsultaFisioterapia
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateMedicinaConsultaFisioterapia actualiza un registro de la tabla consultafisioterapia
// El registro a actualizar no existe
func UpdateMedicinaConsultaFisioterapia(m *MedicinaConsultaFisioterapia) (err error) {
	o := orm.NewOrm()
	v := MedicinaConsultaFisioterapia{IdMedicinaConsultaFisioterapia: m.IdMedicinaConsultaFisioterapia}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMedicinaConsultaFisioterapia  elimina un registro de la tabla consultafisioterapia
// El registro a eliminar no existe
func DeleteMedicinaConsultaFisioterapia(id int) (err error) {
	o := orm.NewOrm()
	v := MedicinaConsultaFisioterapia{IdMedicinaConsultaFisioterapia: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MedicinaConsultaFisioterapia{IdMedicinaConsultaFisioterapia: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
