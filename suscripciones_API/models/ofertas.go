package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Ofertas struct {
	Id                 int                 `orm:"column(id_ofertas);pk"`
	IdEmpleador        *Usuarios           `orm:"column(id_empleador);rel(fk)"`
	IdDiasTrabajados   *DiasTrabajados     `orm:"column(id_dias_trabajados);rel(fk)"`
	Titulo             string              `orm:"column(titulo)"`
	Descripcion        string              `orm:"column(descripcion);null"`
	IdTipoInmueble     *TipoInmueble       `orm:"column(id_tipo_inmueble);rel(fk)"`
	IdTamanoInmueble   *TamanoInmueble     `orm:"column(id_tamano_inmueble);rel(fk)"`
	NumHabitaciones    int                 `orm:"column(num_habitaciones);null"`
	NumBanos           int                 `orm:"column(num_banos);null"`
	Barrio             string              `orm:"column(barrio);null"`
	Direccion          string              `orm:"column(direccion)"`
	Ciudad             string              `orm:"column(ciudad)"`
	TarifaHora         float64             `orm:"column(tarifa_hora)"`
	TarifaJornada      float64             `orm:"column(tarifa_jornada);null"`
	HorarioDescripcion string              `orm:"column(horario_descripcion);null"`
	IdModalidadTrabajo *ModalidadesTrabajo `orm:"column(id_modalidad_trabajo);rel(fk)"`
	InsumosCliente     bool                `orm:"column(insumos_cliente)"`
	Vigilancia         bool                `orm:"column(vigilancia)"`
	IdTipoMascota      *Tipos              `orm:"column(id_tipo_mascota);rel(fk)"`
	CondicionesExtra   string              `orm:"column(condiciones_extra);null"`
	IdEstado           *Estados            `orm:"column(id_estado);rel(fk)"`
	Destacada          bool                `orm:"column(destacada)"`
	Nueva              bool                `orm:"column(nueva)"`
	Vistas             int                 `orm:"column(vistas)"`
	Activo             bool                `orm:"column(activo)"`
	CreadoEn           time.Time           `orm:"column(creado_en);type(timestamp without time zone);auto_now_add"`
	ActualizadoEn      time.Time           `orm:"column(actualizado_en);type(timestamp without time zone);auto_now_add"`
	ExpiraEn           time.Time           `orm:"column(expira_en);type(timestamp without time zone)"`
}

func (t *Ofertas) TableName() string {
	return "ofertas"
}

func init() {
	orm.RegisterModel(new(Ofertas))
}

// AddOfertas insert a new Ofertas into database and returns
// last inserted Id on success.
func AddOfertas(m *Ofertas) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOfertasById retrieves Ofertas by Id. Returns error if
// Id doesn't exist
func GetOfertasById(id int) (v *Ofertas, err error) {
	o := orm.NewOrm()
	v = &Ofertas{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOfertas retrieves all Ofertas matches certain condition. Returns empty list if
// no records exist
func GetAllOfertas(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Ofertas))
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

	var l []Ofertas
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

// UpdateOfertas updates Ofertas by Id and returns error if
// the record to be updated doesn't exist
func UpdateOfertasById(m *Ofertas) (err error) {
	o := orm.NewOrm()
	v := Ofertas{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOfertas deletes Ofertas by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOfertas(id int) (err error) {
	o := orm.NewOrm()
	v := Ofertas{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Ofertas{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
