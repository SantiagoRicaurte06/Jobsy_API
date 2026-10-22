package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// Pagos registra un cobro. Un pago cuelga de un pedido de la tienda O de una
// contratacion de empleo, nunca de ambos ni de ninguno.
type Pagos struct {
	// pk;auto: antes era solo "pk", asi que el ORM enviaba el Id en cada
	// INSERT. Sin asignarlo explicitamente iba 0, y el segundo pago chocaba
	// con el primero por clave primaria duplicada.
	Id int `orm:"column(id_pagos);pk;auto"`

	// Los dos origenes son excluyentes y por tanto ambos anulables. Antes los
	// dos eran obligatorios, lo que hacia imposible registrar cualquiera de
	// los dos casos: un pago de tienda no tiene contratacion, y uno de
	// contratacion no tiene pedido.
	IdPedidos        *Facturacion `orm:"column(id_pedidos);rel(fk);null"`
	IdContrataciones *int         `orm:"column(id_contrataciones);null"`

	IdUsuarios   int         `orm:"column(id_usuarios)"`
	IdMetodoPago *MetodoPago `orm:"column(id_metodo_pago);rel(fk)"`
	IdTipoPago   int         `orm:"column(id_tipo_pago)"`
	IdEstado     int         `orm:"column(id_estado)"`

	// Solo aplica a pagos contra cuenta bancaria; una tarjeta no tiene tipo
	// de cuenta. Los propios inserts de semilla omiten esta columna.
	IdTipoCuenta *int `orm:"column(id_tipo_cuenta);null"`

	Monto          float64   `orm:"column(monto)"`
	ReferenciaExt  *string   `orm:"column(referencia_ext);null"`
	Banco          *string   `orm:"column(banco);null"`
	TitularTarjeta *string   `orm:"column(titular_tarjeta);null"`
	Ultimos4       *string   `orm:"column(ultimos4);null"`
	Activo         bool      `orm:"column(activo)"`
	CreadoEn       time.Time `orm:"column(creado_en);type(timestamp without time zone);auto_now_add"`
	// auto_now, no auto_now_add: con auto_now_add la columna se fijaba al
	// crear y no se tocaba nunca mas, asi que "actualizado_en" mentia.
	ActualizadoEn time.Time `orm:"column(actualizado_en);type(timestamp without time zone);auto_now"`
}

func (t *Pagos) TableName() string {
	return "pagos"
}

func init() {
	orm.RegisterModel(new(Pagos))
}

// AddPagos insert a new Pagos into database and returns
// last inserted Id on success.
func AddPagos(m *Pagos) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPagosById retrieves Pagos by Id. Returns error if
// Id doesn't exist
func GetPagosById(id int) (v *Pagos, err error) {
	o := orm.NewOrm()
	v = &Pagos{Id: id}
	if err = o.Read(v); err == nil {
		// Solo IdPedidos e IdMetodoPago son relaciones. Aqui se llamaba
		// ademas sobre IdUsuarios, IdTipoPago, IdEstado e IdTipoCuenta, que
		// son enteros sueltos: LoadRelated sobre un campo no relacional entra
		// en la rama por defecto de queryRelated y revienta. Tambien estaba
		// IdMetodoPago repetido.
		o.LoadRelated(v, "IdPedidos")
		o.LoadRelated(v, "IdMetodoPago")
		return v, nil
	}
	return nil, err
}

// GetAllPagos retrieves all Pagos matches certain condition. Returns empty list if
// no records exist
func GetAllPagos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Pagos)).RelatedSel()
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

	var l []Pagos
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

// UpdatePagos updates Pagos by Id and returns error if
// the record to be updated doesn't exist
func UpdatePagosById(m *Pagos) (err error) {
	o := orm.NewOrm()
	v := Pagos{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePagos deletes Pagos by Id and returns error if
// the record to be deleted doesn't exist
func DeletePagos(id int) (err error) {
	o := orm.NewOrm()
	v := Pagos{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Pagos{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
