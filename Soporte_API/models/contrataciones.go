package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Contrataciones struct {
	Id                 int                 `orm:"column(id_contrataciones);pk"`
	IdEmpleador        *Usuarios           `orm:"column(id_empleador);rel(fk)"`
	IdEmpleado         *Usuarios           `orm:"column(id_empleado);rel(fk)"`
	IdOfertas          *Ofertas            `orm:"column(id_ofertas);rel(fk)"`
	IdPostulaciones    *Postulaciones      `orm:"column(id_postulaciones);rel(fk)"`
	IdEstado           *Estados            `orm:"column(id_estado);rel(fk)"`
	DireccionServicio  string              `orm:"column(direccion_servicio)"`
	FechaInicio        time.Time           `orm:"column(fecha_inicio);type(date)"`
	HoraInicio         string              `orm:"column(hora_inicio)"`
	DuracionHoras      float64             `orm:"column(duracion_horas);null"`
	IdModalidad        *ModalidadesTrabajo `orm:"column(id_modalidad);rel(fk)"`
	TarifaAcordada     float64             `orm:"column(tarifa_acordada)"`
	TotalEstimado      float64             `orm:"column(total_estimado);null"`
	DescripcionTrabajo string              `orm:"column(descripcion_trabajo);null"`
	NotaAdicional      string              `orm:"column(nota_adicional);null"`
	Activo             bool                `orm:"column(activo)"`
	CreadoEn           time.Time           `orm:"column(creado_en);type(timestamp without time zone);auto_now_add"`
	ActualizadoEn      time.Time           `orm:"column(actualizado_en);type(timestamp without time zone);auto_now_add"`
	FinalizadoEn       time.Time           `orm:"column(finalizado_en);type(timestamp without time zone);null"`
}

func (t *Contrataciones) TableName() string {
	return "contrataciones"
}

func init() {
	orm.RegisterModel(new(Contrataciones))
}

// AddContrataciones insert a new Contrataciones into database and returns
// last inserted Id on success.
func AddContrataciones(m *Contrataciones) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetContratacionesById retrieves Contrataciones by Id. Returns error if
// Id doesn't exist
func GetContratacionesById(id int) (v *Contrataciones, err error) {
	o := orm.NewOrm()
	v = &Contrataciones{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllContrataciones retrieves all Contrataciones matches certain condition. Returns empty list if
// no records exist
func GetAllContrataciones(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Contrataciones))
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

	var l []Contrataciones
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

// UpdateContrataciones updates Contrataciones by Id and returns error if
// the record to be updated doesn't exist
func UpdateContratacionesById(m *Contrataciones) (err error) {
	o := orm.NewOrm()
	v := Contrataciones{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteContrataciones deletes Contrataciones by Id and returns error if
// the record to be deleted doesn't exist
func DeleteContrataciones(id int) (err error) {
	o := orm.NewOrm()
	v := Contrataciones{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Contrataciones{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
