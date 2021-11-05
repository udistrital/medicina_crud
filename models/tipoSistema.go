package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TipoSistema struct {
	Id                int        `orm:"column(id_tipo_sistema);pk;auto"`
	Nombre            string     `orm:"column(nombre);null"`
	Activo            bool       `orm:"column(activo);null"`
	FechaModificacion *time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
}

func (p *TipoSistema) TableName() string {
	return "TipoSistema"
}
func init() {
	orm.RegisterModel(new(TipoSistema))
}

// AddTipoSistema inserta un registro en la tabla TipoSistema
// Último registro insertado con éxito
func AddTipoSistema(m *TipoSistema) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTipoSistemaById obtiene un registro de la tabla TipoSistema por su id
// Id no existe
func GetTipoSistemaById(id int) (v *TipoSistema, err error) {
	o := orm.NewOrm()
	v = &TipoSistema{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTipoSistema obtiene todos los registros de la tabla TipoSistema
// No existen registros
func GetAllTipoSistema(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TipoSistema))
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
	var l []TipoSistema
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

// UpdateTipoSistema actualiza un registro de la tabla TipoSistema
// El registro a actualizar no existe
func UpdateTipoSistema(m *TipoSistema) (err error) {
	o := orm.NewOrm()
	v := TipoSistema{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteTipoSistema elimina un registro de la tabla TipoSistema
// El registro a eliminar no existe
func DeleteTipoSistema(id int) (err error) {
	o := orm.NewOrm()
	v := TipoSistema{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TipoSistema{Id: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
