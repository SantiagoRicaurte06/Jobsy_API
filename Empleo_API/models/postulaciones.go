package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Postulaciones struct {
	Id                 int        `orm:"column(id_postulaciones);pk;auto"`
	IdOfertas          *Ofertas   `orm:"column(id_ofertas);rel(fk)"`
	IdEmpleado         int        `orm:"column(id_empleado)"`
	IdHojasVida        *HojasVida `orm:"column(id_hojas_vida);rel(fk)"`
	IdEstado           int        `orm:"column(id_estado)"`
	DisponibleDesde    time.Time  `orm:"column(disponible_desde);type(date)"`
	IdModalidadDisp    int        `orm:"column(id_modalidad_disp)"`
	IdTurnoPostulacion int        `orm:"column(id_turno_postulacion)"`
	NotaDisponibilidad string     `orm:"column(nota_disponibilidad);null"`
	IdFormaCobro       int        `orm:"column(id_forma_cobro)"`
	TarifaMin          float64    `orm:"column(tarifa_min);null"`
	TarifaMax          float64    `orm:"column(tarifa_max);null"`
	MensajeEmpleador   string     `orm:"column(mensaje_empleador);null"`
	AceptaTerminos     bool       `orm:"column(acepta_terminos)"`
	Activo             bool       `orm:"column(activo)"`
	PostuladoEn        time.Time  `orm:"column(postulado_en);type(timestamp without time zone);auto_now_add"`
	RespondidoEn       time.Time  `orm:"column(respondido_en);type(timestamp without time zone);null"`
	NotaRechazo        string     `orm:"column(nota_rechazo);null"`
	OrigenPostulacion  bool       `orm:"column(origen_postulacion);null"`
	ActualizadoEn      time.Time  `orm:"column(actualizado_en);type(timestamp without time zone);auto_now"`
}

func (t *Postulaciones) TableName() string {
	return "postulaciones"
}

func init() {
	orm.RegisterModel(new(Postulaciones))
}

// AddPostulaciones insert a new Postulaciones into database and returns
// last inserted Id on success.
func AddPostulaciones(m *Postulaciones) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPostulacionesById retrieves Postulaciones by Id. Returns error if
// Id doesn't exist
func GetPostulacionesById(id int) (v *Postulaciones, err error) {
	o := orm.NewOrm()
	v = &Postulaciones{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "IdOfertas")
		o.LoadRelated(v, "IdEmpleado")
		o.LoadRelated(v, "IdHojasVida")
		o.LoadRelated(v, "IdEstado")
		o.LoadRelated(v, "IdModalidadDisp")
		o.LoadRelated(v, "IdTurnoPostulacion")
		o.LoadRelated(v, "IdFormaCobro")
		return v, nil
	}
	return nil, err
}

// GetAllPostulaciones retrieves all Postulaciones matches certain condition. Returns empty list if
// no records exist
func GetAllPostulaciones(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Postulaciones)).RelatedSel()
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

	var l []Postulaciones
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

// UpdatePostulaciones updates Postulaciones by Id and returns error if
// the record to be updated doesn't exist
func UpdatePostulacionesById(m *Postulaciones) (err error) {
	o := orm.NewOrm()
	v := Postulaciones{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePostulaciones deletes Postulaciones by Id and returns error if
// the record to be deleted doesn't exist
func DeletePostulaciones(id int) (err error) {
	o := orm.NewOrm()
	v := Postulaciones{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Postulaciones{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
