package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Sistema struct {
	Id              int              `orm:"column(id_sistema);pk;auto"`
	HistoriaClinica *HistoriaClinica `orm:"column(id_historia_clinica);rel(fk);null"`
	HojaHistoria    *HojaHistoria    `orm:"column(id_hoja_historia);rel(fk);null"`
	TipoSistema     *TipoSistema     `orm:"column(id_tipo_sistema);rel(fk);null"`
	Observacion     string           `orm:"column(observacion);null"`
}

func (p *Sistema) TableName() string {
	return "sistema"
}
func init() {
	orm.RegisterModel(new(Sistema))
}

// AddSistema inserta un registro en la tabla Sistema
// Último registro insertado con éxito
func AddSistema(m *Sistema) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSistemaById obtiene un registro de la tabla Sistema por su id
// Id no existe
func GetSistemaById(id int) (v *Sistema, err error) {
	o := orm.NewOrm()
	v = &Sistema{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSistema obtiene todos los registros de la tabla Sistema
// No existen registros
func GetAllSistema(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sistema))
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
	var l []Sistema
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

// UpdateSistema actualiza un registro de la tabla Sistema
// El registro a actualizar no existe
func UpdateSistema(m *Sistema) (err error) {
	o := orm.NewOrm()
	v := Sistema{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteSistema elimina un registro de la tabla Sistema
// El registro a eliminar no existe
func DeleteSistema(id int) (err error) {
	o := orm.NewOrm()
	v := Sistema{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Sistema{Id: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
