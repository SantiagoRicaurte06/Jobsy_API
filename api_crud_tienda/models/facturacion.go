package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Facturacion struct {
	Id               int       `orm:"column(id_pedidos);pk"`
	IdUsuarios       int       `orm:"column(id_usuarios);rel(fk)"`
	IdCarritos       *Carritos `orm:"column(id_carritos);rel(fk)"`
	IdEstado         int       `orm:"column(id_estado)`
	Subtotal         float64   `orm:"column(subtotal)"`
	DescuentoTotal   float64   `orm:"column(descuento_total)"`
	CuponDescuento   float64   `orm:"column(cupon_descuento)"`
	Envio            float64   `orm:"column(envio)"`
	Total            float64   `orm:"column(total)"`
	FacturaNombre    string    `orm:"column(factura_nombre);null"`
	FacturaDoc       string    `orm:"column(factura_doc);null"`
	FacturaCorreo    string    `orm:"column(factura_correo);null"`
	FacturaTelefono  string    `orm:"column(factura_telefono);null"`
	FacturaDireccion string    `orm:"column(factura_direccion);null"`
	FacturaCiudad    string    `orm:"column(factura_ciudad);null"`
	FacturaDpto      string    `orm:"column(factura_dpto);null"`
	Activo           bool      `orm:"column(activo)"`
	CreadoEn         time.Time `orm:"column(creado_en);type(timestamp without time zone);auto_now_add"`
	ActualizadoEn    time.Time `orm:"column(actualizado_en);type(timestamp without time zone);auto_now_add"`
}

func (t *Facturacion) TableName() string {
	return "facturacion"
}

func init() {
	orm.RegisterModel(new(Facturacion))
}

// AddFacturacion insert a new Facturacion into database and returns
// last inserted Id on success.
func AddFacturacion(m *Facturacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFacturacionById retrieves Facturacion by Id. Returns error if
// Id doesn't exist
func GetFacturacionById(id int) (v *Facturacion, err error) {
	o := orm.NewOrm()
	v = &Facturacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFacturacion retrieves all Facturacion matches certain condition. Returns empty list if
// no records exist
func GetAllFacturacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Facturacion))
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

	var l []Facturacion
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

// UpdateFacturacion updates Facturacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateFacturacionById(m *Facturacion) (err error) {
	o := orm.NewOrm()
	v := Facturacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFacturacion deletes Facturacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFacturacion(id int) (err error) {
	o := orm.NewOrm()
	v := Facturacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Facturacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
