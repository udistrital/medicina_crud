package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MedicinaSistemas struct {
	IdSistema      int    `orm:"column(id_sistema);pk;auto"`
	NombreSistema  string `orm:"column(nombre_sistema);null"`
	Observacion    string `orm:"column(observacion);null"`
	IdHojaHistoria *int   `orm:"column(id_hoja_historia);rel(fk);null"`
}

func (p *MedicinaSistemas) TableName() string {
	return "sistemas"
}
func init() {
	orm.RegisterModel(new(MedicinaSistemas))
}

// AddMedicinaSistemas inserta un registro en la tabla sistemas
// Último registro insertado con éxito
func AddMedicinaSistemas(m *MedicinaSistemas) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMedicinaSistemasById obtiene un registro de la tabla sistemas por su id
// Id no existe
func GetMedicinaSistemasById(id int) (v *MedicinaSistemas, err error) {
	o := orm.NewOrm()
	v = &MedicinaSistemas{IdSistema: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMedicinaSistemas obtiene todos los registros de la tabla sistemas
// No existen registros
func GetAllMedicinaSistemas(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MedicinaSistemas))
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
	var l []MedicinaSistemas
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

// UpdateMedicinaSistemas actualiza un registro de la tabla sistemas
// El registro a actualizar no existe
func UpdateMedicinaSistemas(m *MedicinaSistemas) (err error) {
	o := orm.NewOrm()
	v := MedicinaSistemas{IdSistema: m.IdSistema}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteMedicinaSistemas elimina un registro de la tabla sistemas
// El registro a eliminar no existe
func DeleteMedicinaSistemas(id int) (err error) {
	o := orm.NewOrm()
	v := MedicinaSistemas{IdSistema: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MedicinaSistemas{IdSistema: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
