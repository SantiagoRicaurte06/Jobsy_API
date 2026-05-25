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

// TrCarritoItemsController operations for TrCarritoItems
type TrCarritoItemsController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrCarritoItemsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TrCarritoItems
// @Param body body models.TrCarritoItems true "body for TrCarritoItems content"
// @Success 201 {int} models.TrCarritoItems
// @Failure 400 body is empty or invalid
// @router / [post]
func (c *TrCarritoItemsController) Post() {
	var v models.TrCarritoItems

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error("Error al procesar el JSON del cuerpo de la solicitud:", err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El cuerpo de la solicitud no es un JSON válido",
			"data":    nil,
		}
		c.ServeJSON()
		return
	}

	if _, err := models.AddTrCarritoItems(&v); err != nil {
		logs.Error("Error al crear el item del carrito:", err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "Error al crear el item del carrito: " + err.Error(),
			"data":    nil,
		}
	} else {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  201,
			"message": "Item del carrito creado exitosamente",
			"data":    v,
		}
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get TrCarritoItems by id
// @Param id path string true "The key for staticblock"
// @Success 200 {object} models.TrCarritoItems
// @Failure 400 id is invalid
// @Failure 404 TrCarritoItems not found
// @router /:id [get]
func (c *TrCarritoItemsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("ID inválido:", err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El id del item del carrito debe ser un número entero válido",
			"data":    nil,
		}
		c.ServeJSON()
		return
	}

	v, err := models.GetTrCarritoItemsById(id)
	if err != nil {
		logs.Error("Error al obtener item del carrito con id "+idStr+":", err)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  404,
			"message": "No se encontró el item del carrito con el id especificado",
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
// @Description get TrCarritoItems
// @Param query query string false "Filter. e.g. col1:v1,col2:v2 ..."
// @Param fields query string false "Fields returned. e.g. col1,col2 ..."
// @Param sortby query string false "Sorted-by fields. e.g. col1,col2 ..."
// @Param order query string false "Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param limit query string false "Limit the size of result set. Must be an integer"
// @Param offset query string false "Start position of result set. Must be an integer"
// @Success 200 {object} models.TrCarritoItems
// @Failure 400 invalid query parameters
// @router / [get]
func (c *TrCarritoItemsController) GetAll() {
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
				logs.Error("Formato de filtro inválido")

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

	l, err := models.GetAllTrCarritoItems(query, fields, sortby, order, offset, limit)

	if err != nil {
		logs.Error("Error al obtener los items del carrito:", err)

		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "Error al obtener los items del carrito: " + err.Error(),
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
// @Description update the TrCarritoItems
// @Param id path string true "The id you want to update"
// @Param body body models.TrCarritoItems true "body for TrCarritoItems content"
// @Success 200 {object} models.TrCarritoItems
// @Failure 400 id is invalid or body is invalid
// @Failure 404 TrCarritoItems not found
// @router /:id [put]
func (c *TrCarritoItemsController) Put() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("ID inválido:", err)

		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El id del item del carrito debe ser un número entero válido",
			"data":    nil,
		}

		c.ServeJSON()
		return
	}

	v := models.TrCarritoItems{Id: id}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error("Error al procesar el JSON del cuerpo de la solicitud:", err)

		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El cuerpo de la solicitud no es un JSON válido",
			"data":    nil,
		}

		c.ServeJSON()
		return
	}

	if err := models.UpdateTrCarritoItemsById(&v); err != nil {
		logs.Error("Error al actualizar el item del carrito con id "+idStr+":", err)

		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "Error al actualizar el item del carrito: " + err.Error(),
			"data":    nil,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"message": "Item del carrito actualizado exitosamente",
			"data":    v,
		}
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the TrCarritoItems
// @Param id path string true "The id you want to delete"
// @Success 200 {string} delete success
// @Failure 400 id is invalid
// @Failure 404 TrCarritoItems not found
// @router /:id [delete]
func (c *TrCarritoItemsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.Error("ID inválido:", err)

		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "El id del item del carrito debe ser un número entero válido",
			"data":    nil,
		}

		c.ServeJSON()
		return
	}

	if err := models.DeleteTrCarritoItems(id); err != nil {
		logs.Error("Error al eliminar el item del carrito con id "+idStr+":", err)

		c.Data["json"] = map[string]interface{}{
			"success": false,
			"status":  400,
			"message": "Error al eliminar el item del carrito: " + err.Error(),
			"data":    nil,
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"status":  200,
			"message": "Item del carrito eliminado exitosamente",
			"data":    nil,
		}
	}

	c.ServeJSON()
}
