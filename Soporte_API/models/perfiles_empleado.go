package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type PerfilesEmpleado struct {
	Id                 int                 `orm:"column(id_perfiles_empleado);pk"`
	IdUsuarios         *Usuarios           `orm:"column(id_usuarios);rel(fk)"`
	Bio                string              `orm:"column(bio);null"`
	IdNivelExperiencia *NivelesExperiencia `orm:"column(id_nivel_experiencia);rel(fk)"`
	AniosExperiencia   int                 `orm:"column(anios_experiencia);null"`
	TarifaHora         float64             `orm:"column(tarifa_hora);null"`
	TarifaNegociable   bool                `orm:"column(tarifa_negociable)"`
	IdModalidadTrabajo *ModalidadesTrabajo `orm:"column(id_modalidad_trabajo);rel(fk)"`
	IdTurnoPreferido   *Turnos             `orm:"column(id_turno_preferido);rel(fk)"`
	DistanciaMaxKm     int                 `orm:"column(distancia_max_km);null"`
	BarrioZona         string              `orm:"column(barrio_zona);null"`
	HvVisible          bool                `orm:"column(hv_visible)"`
	AparecerBusquedas  bool                `orm:"column(aparecer_busquedas)"`
	RatingPromedio     float64             `orm:"column(rating_promedio)"`
	TotalResenas       int                 `orm:"column(total_resenas)"`
	TotalTrabajos      int                 `orm:"column(total_trabajos)"`
	Verificado         bool                `orm:"column(verificado)"`
	Activo             bool                `orm:"column(activo)"`
	CreadoEn           time.Time           `orm:"column(creado_en);type(timestamp without time zone);auto_now_add"`
	ActualizadoEn      time.Time           `orm:"column(actualizado_en);type(timestamp without time zone);auto_now_add"`
}

func (t *PerfilesEmpleado) TableName() string {
	return "perfiles_empleado"
}

func init() {
	orm.RegisterModel(new(PerfilesEmpleado))
}

// AddPerfilesEmpleado insert a new PerfilesEmpleado into database and returns
// last inserted Id on success.
func AddPerfilesEmpleado(m *PerfilesEmpleado) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPerfilesEmpleadoById retrieves PerfilesEmpleado by Id. Returns error if
// Id doesn't exist
func GetPerfilesEmpleadoById(id int) (v *PerfilesEmpleado, err error) {
	o := orm.NewOrm()
	v = &PerfilesEmpleado{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPerfilesEmpleado retrieves all PerfilesEmpleado matches certain condition. Returns empty list if
// no records exist
func GetAllPerfilesEmpleado(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PerfilesEmpleado))
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

	var l []PerfilesEmpleado
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

// UpdatePerfilesEmpleado updates PerfilesEmpleado by Id and returns error if
// the record to be updated doesn't exist
func UpdatePerfilesEmpleadoById(m *PerfilesEmpleado) (err error) {
	o := orm.NewOrm()
	v := PerfilesEmpleado{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePerfilesEmpleado deletes PerfilesEmpleado by Id and returns error if
// the record to be deleted doesn't exist
func DeletePerfilesEmpleado(id int) (err error) {
	o := orm.NewOrm()
	v := PerfilesEmpleado{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PerfilesEmpleado{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
