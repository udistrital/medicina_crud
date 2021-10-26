package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MedicinaHistoriaClinica struct {
	IdHistoriaClinica int `orm:"column(id_historia_clinica);pk;auto"`
	IdTercero         int `orm:"column(id_tercero);null"`
}

func (t *MedicinaHistoriaClinica) TableName() string {
	return "historiaclinica"
}
func init() {
	orm.RegisterModel(new(MedicinaHistoriaClinica))
}

// AddMedicinaHistoriaClinica inserta un registro en la tabla  historiaclinica
// Último registro insertado con éxito
func AddMedicinaHistoriaClinica(m *MedicinaHistoriaClinica) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMedicinaHistoriaClinicaById obtiene un registro de la tabla historiaclinica por su id
// Id no existe
func GetMedicinaHistoriaClinicaById(id int) (v *MedicinaHistoriaClinica, err error) {
	o := orm.NewOrm()
	v = &MedicinaHistoriaClinica{IdHistoriaClinica: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMedicinaHistoriaClinica obtiene todos los registros de la tabla historiaclinica
// No existen registros
func GetAllMedicinaHistoriaClinica(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MedicinaHistoriaClinica))
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
	var l []MedicinaHistoriaClinica
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

// UpdateMedicinaHistoriaClinica actualiza un registro de la tabla historiaclinica
// El registro a actualizar no existe
func UpdateMedicinaHistoriaClinica(m *MedicinaHistoriaClinica) (err error) {
	o := orm.NewOrm()
	v := MedicinaHistoriaClinica{IdHistoriaClinica: m.IdHistoriaClinica}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteMedicinaHistoriaClinica elimina un registro de la tabla historiaclinica
// El registro a eliminar no existe
func DeleteMedicinaHistoriaClinica(id int) (err error) {
	o := orm.NewOrm()
	v := MedicinaHistoriaClinica{IdHistoriaClinica: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MedicinaHistoriaClinica{IdHistoriaClinica: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
