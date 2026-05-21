package controllers

import (
	"api_crud_tienda/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// CuponesController operations for Cupones
type CuponesController struct {
	beego.Controller
}

// URLMapping ...
func (c *CuponesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Cupones
// @Param	body		body 	models.Cupones	true		"body for Cupones content"
// @Success 201 {int} models.Cupones
// @Failure 403 body is empty
// @router / [post]
func (c *CuponesController) Post() {
	var v models.Cupones
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddCupones(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  201,
				"message": "Cupón creado exitosamente",
				"data":    v,
			}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"message": err.Error(),
				"data":    nil,
			}
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": err.Error(),
			"data":    nil,
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Cupones by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Cupones
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CuponesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El id del cupón debe ser un número entero válido",
			"data":    nil,
		}
		c.ServeJSON()
		return
	}
	v, err := models.GetCuponesById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  404,
			"message": "No se encontró el cupón con el id especificado",
			"data":    nil,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"message": "Consulta exitosa",
			"data":    v,
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Cupones
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Cupones
// @Failure 403
// @router / [get]
func (c *CuponesController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"status":  400,
					"message": errors.New("El filtro de búsqueda tiene un formato inválido, use el formato campo:valor").Error(),
					"data":    nil,
				}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCupones(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": err.Error(),
			"data":    nil,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"message": "Consulta exitosa",
			"data":    l,
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Cupones
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Cupones	true		"body for Cupones content"
// @Success 200 {object} models.Cupones
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CuponesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El id del cupón debe ser un número entero válido",
			"data":    nil,
		}
		c.ServeJSON()
		return
	}
	v := models.Cupones{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateCuponesById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  200,
				"message": "Cupón actualizado exitosamente",
				"data":    v,
			}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"message": "No se pudo actualizar el cupón, verifique que el id exista y que los datos enviados sean correctos",
				"data":    nil,
			}
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El cuerpo de la solicitud no tiene un formato JSON válido",
			"data":    nil,
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Cupones
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CuponesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El id del cupón debe ser un número entero válido",
			"data":    nil,
		}
		c.ServeJSON()
		return
	}
	if err := models.DeleteCupones(id); err == nil {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"message": "Cupón eliminado exitosamente",
			"data":    nil,
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "No se pudo eliminar el cupón, es posible que esté siendo utilizado en pedidos o que el id no exista",
			"data":    nil,
		}
	}
	c.ServeJSON()
}
