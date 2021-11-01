package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"Medicina/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type HojaHistoriaController struct {
	beego.Controller
}

//URLMapping ...
func (c *HojaHistoriaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description agregar un registro en la tabla HojaHistoria
// @Param	body		body 	models.HojaHistoria	true		"Cuerpo para el contenido de ConsultaFisioterapia"
// @Success 201 {int} models.HojaHistoria
// @Failure 403 Cuerpo Vacío
// @router / [post]
func (c *HojaHistoriaController) Post() {
	var v models.HojaHistoria
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddHojaHistoria(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description consultar un registro de la tabla HojaHistoria por su id
// @Param	id		path 	string	true		"Id a consultar"
// @Success 200 {object} models.HojaHistoria
// @Failure 403 :id está vacío
// @router /:id [get]
func (c *HojaHistoriaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetHojaHistoriaById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description consulta todos los registros de la tabla HojaHistoria
// @Param   query   consulta    string  false   "Filtro. Por ejemplo, col1: v1, col2: v2 ..."
// @Param   fields  consulta    string  false   "Campos devueltos. Por ejemplo, col1, col2 ..."
// @Param   sortby  consulta    string  false   "Campos ordenados por. Por ejemplo, Col1, col2 ..."
// @Param   order   consulta    string  false   "El orden correspondiente a cada campo de clasificación, si es un valor único, se aplica a todos los campos de clasificación. Por ejemplo, desc, asc ..."
// @Param   limit   consulta    string  false   "Limite el tamaño del conjunto de resultados. Debe ser un número entero"
// @Param   offset  consulta    string  false   "Posición inicial del conjunto de resultados. Debe ser un número entero"
// @Success 200 {object} models.HojaHistoria
// @Failure 403
// @router / [get]
func (c *HojaHistoriaController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllHojaHistoria(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

//Put ...
// @Title Put
// @Description actualizar un registro de la tabla HojaHistoria
// @Param	id		path 	string	true		"Id del registro a actualizar"
// @Param	body		body 	models.HojaHistoria	true		"Cuerpo para el contenido de ConsultaFisioterapia"
// @Success 200 {object} models.HojaHistoria
// @Failure 403 :id no es entero
// @router /:id [put]
func (c *HojaHistoriaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.HojaHistoria{IdHojaHistoria: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateHojaHistoria(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

//Delete ...
// @Title Delete
// @Description elimina un registro de la tabla HojaHistoria
// @Param	id		path 	string	true		"Id del registro a eliminar"
// @Success 200 {string} borrado exitoso!
// @Failure 403 Id vacío
// @router /:id [delete]
func (c *HojaHistoriaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteHojaHistoria(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
