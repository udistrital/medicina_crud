package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TipoAntecedente struct {
	Id                int        `orm:"column(id_tipo_antecedente);pk;auto"`
	Nombre            string     `orm:"column(nombre);null"`
	Descripcion       string     `orm:"column(descripcion);null"`
	Activo            bool       `orm:"column(activo);null"`
	FechaCreacion     *time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);null"`
	FechaModificacion *time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
}

func (p *TipoAntecedente) TableName() string {
	return "tipoantecedente"
}
func init() {
	orm.RegisterModel(new(TipoAntecedente))
}

// AddTipoAntecedente inserta un registro en la tabla tipoantecedente
// Último registro insertado con éxito
func AddTipoAntecedente(m *TipoAntecedente) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTipoAntecedenteById obtiene un registro de la tabla tipoantecedente por su id
// Id no existe
func GetTipoAntecedenteById(id int) (v *TipoAntecedente, err error) {
	o := orm.NewOrm()
	v = &TipoAntecedente{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTipoAntecedente obtiene todos los registros de la tabla tipoantecedente
// No existen registros
func GetAllTipoAntecedente(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TipoAntecedente))
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
	var l []TipoAntecedente
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

// UpdateTipoAntecedente actualiza un registro de la tabla tipoantecedente
// El registro a actualizar no existe
func UpdateTipoAntecedente(m *TipoAntecedente) (err error) {
	o := orm.NewOrm()
	v := TipoAntecedente{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteTipoAntecedente elimina un registro de la tabla tipoantecedente
// El registro a eliminar no existe
func DeleteTipoAntecedente(id int) (err error) {
	o := orm.NewOrm()
	v := TipoAntecedente{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TipoAntecedente{Id: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
