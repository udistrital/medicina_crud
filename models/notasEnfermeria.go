package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type NotasEnfermeria struct {
	Id              	int              `orm:"column(id_notas);pk;auto"`
	HistoriaClinica 	*HistoriaClinica `orm:"column(id_historia_clinica);rel(fk);null"`
	HojaHistoria    	*HojaHistoria    `orm:"column(id_hoja_historia);rel(fk);null"`
	Descripcion     	string           `orm:"column(descripcion);null"`
	Ta                  string       	 `orm:"column(ta);null"`
	Fc                  string       	 `orm:"column(fc);null"`
	Sao2                string       	 `orm:"column(sao2);null"`
	Imc                 string       	 `orm:"column(imc);null"`
	Fr                  string       	 `orm:"column(fr);null"`
	Temperatura         string       	 `orm:"column(temperatura);null"`
	Peso                string       	 `orm:"column(peso);null"`
	Talla               string       	 `orm:"column(talla);null"`
	EstadoGeneral       string       	 `orm:"column(estado_general);null"`
	CabezaYCuello       string       	 `orm:"column(cabeza_cuello);null"`
	Orl                 string       	 `orm:"column(orl);null"`
	Ojos                string       	 `orm:"column(ojos);null"`
	Torax               string       	 `orm:"column(torax);null"`
	RuidosRespiratorios string       	 `orm:"column(ruidos_respiratorios);null"`
	RuidosCardiacos     string       	 `orm:"column(ruidos_cardiacos);null"`
	Abdomen             string       	 `orm:"column(abdomen);null"`
	Neurologico         string       	 `orm:"column(neurologico);null"`
	Genital             string       	 `orm:"column(genital);null"`
	Extremidades        string       	 `orm:"column(extremidades);null"`
	Contacto        	string       	 `orm:"column(contacto);null"`
	FechaCreacion     *time.Time 	     `orm:"column(fecha_creacion);type(timestamp without time zone);null"`
	FechaModificacion *time.Time 		 `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	Activo            bool       		 `orm:"column(activo);null"`
}

func (t *NotasEnfermeria) TableName() string {
	return "notasenfermeria"
}
func init() {
	orm.RegisterModel(new(NotasEnfermeria))
}

// AddNotasEnfermeria inserta un registro en la tabla notasenfermeria
// Último registro insertado con éxito
func AddNotasEnfermeria(m *NotasEnfermeria) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNotasEnfermeriaById obtiene un registro de la tabla notasenfermeria por su id
// Id no existe
func GetNotasEnfermeriaById(id int) (v *NotasEnfermeria, err error) {
	o := orm.NewOrm()
	v = &NotasEnfermeria{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNotasEnfermeria obtiene todos los registros de la tabla notasenfermeria
// No existen registros
func GetAllNotasEnfermeria(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NotasEnfermeria))
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
	var l []NotasEnfermeria
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

// UpdateNotasEnfermeria actualiza un registro de la tabla notasenfermeria
// El registro a actualizar no existe
func UpdateNotasEnfermeria(m *NotasEnfermeria) (err error) {
	o := orm.NewOrm()
	v := NotasEnfermeria{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNotasEnfermeria  elimina un registro de la tabla notasenfermeria
// El registro a eliminar no existe
func DeleteNotasEnfermeria(id int) (err error) {
	o := orm.NewOrm()
	v := NotasEnfermeria{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NotasEnfermeria{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
