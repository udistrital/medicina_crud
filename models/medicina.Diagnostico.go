package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"time"

	"github.com/astaxie/beego/orm"
)

type MedicinaDiagnostico struct {
	IdDiagnosticoMedicina int       `orm:"column(id_diagnostico_medicina);pk;auto"`
	Nombre                string    `orm:"column(nombre);null"`
	Descripcion           string    `orm:"column(descripcion);null"`
	Numero                int       `orm:"column(numero);null"`
	Activo                bool      `orm:"column(activo);null"`
	FechaCreacion         time.Time `orm:"column(fecha_creacion);type(date)"`
	FechaModificacion     time.Time `orm:"column(fecha_modificacion);type(date)"`
	PlanManejo            string    `orm:"column(plan_manejo);null"`
	Analisis              string    `orm:"column(analisis);null"`
}

func (t *MedicinaDiagnostico) TableName() string {
	return "diagnosticomedicina"
}
func init() {
	orm.RegisterModel(new(MedicinaDiagnostico))
}

// AddMedicinaDiagnostico inserta un registro en la tabla diagnosticomedicina
// Último registro insertado con éxito
func AddMedicinaDiagnostico(m *MedicinaDiagnostico) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMedicinaDiagnosticoById obtiene un registro de la tabla diagnostico medicina por su id
// Id no existe
func GetMedicinaDiagnosticoById(id int) (v *MedicinaDiagnostico, err error) {
	o := orm.NewOrm()
	v = &MedicinaDiagnostico{IdDiagnosticoMedicina: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMedicinaDiagnostico obtiene todos los registros de la tabla diagnostico medicina
// No existen registros
func GetAllMedicinaDiagnostico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MedicinaDiagnostico))
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
	var l []MedicinaDiagnostico
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

// UpdateMedicinaDiagnostico actualiza un registro de la tabla diagnosticomedicina
// El registro a actualizar no existe
func UpdateMedicinaDiagnostico(m *MedicinaDiagnostico) (err error) {
	o := orm.NewOrm()
	v := MedicinaDiagnostico{IdDiagnosticoMedicina: m.IdDiagnosticoMedicina}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteMedicinaDiagnostico elimina un registro de la tabla diagnosticomedicina
// El registro a eliminar no existe
func DeleteMedicinaDiagnostico(id int) (err error) {
	o := orm.NewOrm()
	v := MedicinaDiagnostico{IdDiagnosticoMedicina: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MedicinaDiagnostico{IdDiagnosticoMedicina: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
