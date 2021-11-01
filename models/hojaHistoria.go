package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type HojaHistoria struct {
	IdHojaHistoria    int              `orm:"column(id_hoja_historia);pk;auto"`
	IdHistoriaClinica *HistoriaClinica `orm:"column(id_historia_clinica);rel(fk);null"`
	IdDiagnostico     *Diagnostico     `orm:"column(id_diagnostico);rel(fk);null"`
	FechaConsulta     *time.Time       `orm:"column(fecha_consulta);type(timestamp without time zone);null"`
	Motivo            string           `orm:"column(motivo);null"`
	Observacion       string           `orm:"column(observacion);null"`
	Evolucion         string           `orm:"column(evolucion);type(json);null"`
	IdEspecialidad    int              `orm:"column(id_especialidad);type(bigint);null"`
	IdProfesional     float64          `orm:"column(id_profesional);type(bigint);null"`
	IdPersona         float64          `orm:"column(id_persona);type(bigint);null"`
}

func (t *HojaHistoria) TableName() string {
	return "hojahistoria"
}
func init() {
	orm.RegisterModel(new(HojaHistoria))
}

// AddHojaHistoria inserta un registro en la tabla hojahistoria
// Último registro insertado con éxito
func AddHojaHistoria(m *HojaHistoria) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHojaHistoriaById obtiene un registro de la tabla hojahistoria por su id
// Id no existe
func GetHojaHistoriaById(id int) (v *HojaHistoria, err error) {
	o := orm.NewOrm()
	v = &HojaHistoria{IdHojaHistoria: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHojaHistoria obtiene todos los registros de la tabla hojahistoria
// No existen registros
func GetAllHojaHistoria(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HojaHistoria)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []HojaHistoria
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
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

// UpdateHojaHistoria actualiza un registro de la tabla hojahistoria
// El registro a actualizar no existe
func UpdateHojaHistoria(m *HojaHistoria) (err error) {
	o := orm.NewOrm()
	v := HojaHistoria{IdHojaHistoria: m.IdHojaHistoria}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Número de registros actualizados de la base de datos:", num)
		}
	}
	return
}

// DeleteHojaHistoria elimina un registro de la tabla
// El registro a eliminar no existe
func DeleteHojaHistoria(id int) (err error) {
	o := orm.NewOrm()
	v := HojaHistoria{IdHojaHistoria: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HojaHistoria{IdHojaHistoria: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
