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

// TipoTransaccionController operations for TipoTransaccion
type TipoTransaccionController struct {
	beego.Controller
}

// URLMapping ...
func (c *TipoTransaccionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TipoTransaccion
// @Param	body		body 	models.TipoTransaccion	true		"body for TipoTransaccion content"
// @Success 201 {int} models.TipoTransaccion
// @Failure 403 body is empty
// @router / [post]
func (c *TipoTransaccionController) Post() {
	var v models.TipoTransaccion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddTipoTransaccion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  201,
				"message": "Tipo de transacción creado exitosamente",
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
// @Description get TipoTransaccion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TipoTransaccion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TipoTransaccionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El id del tipo de transacción debe ser un número entero válido",
			"data":    nil,
		}
		c.ServeJSON()
		return
	}
	v, err := models.GetTipoTransaccionById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  404,
			"message": "No se encontró el tipo de transacción con el id especificado",
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
// @Description get TipoTransaccion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TipoTransaccion
// @Failure 403
// @router / [get]
func (c *TipoTransaccionController) GetAll() {
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

	l, err := models.GetAllTipoTransaccion(query, fields, sortby, order, offset, limit)
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
// @Description update the TipoTransaccion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TipoTransaccion	true		"body for TipoTransaccion content"
// @Success 200 {object} models.TipoTransaccion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TipoTransaccionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El id del tipo de transacción debe ser un número entero válido",
			"data":    nil,
		}
		c.ServeJSON()
		return
	}
	v := models.TipoTransaccion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTipoTransaccionById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{
				"success": true,
				"status":  200,
				"message": "Tipo de transacción actualizado exitosamente",
				"data":    v,
			}
		} else {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"status":  400,
				"message": "No se pudo actualizar el tipo de transacción, verifique que el id exista y que los datos enviados sean correctos",
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
// @Description delete the TipoTransaccion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TipoTransaccionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El id del tipo de transacción debe ser un número entero válido",
			"data":    nil,
		}
		c.ServeJSON()
		return
	}
	if err := models.DeleteTipoTransaccion(id); err == nil {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"message": "Tipo de transacción eliminado exitosamente",
			"data":    nil,
		}
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "No se pudo eliminar el tipo de transacción, es posible que esté siendo utilizado en transacciones existentes o que el id no exista",
			"data":    nil,
		}
	}
	c.ServeJSON()
}
