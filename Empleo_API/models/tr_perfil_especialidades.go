package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type TrPerfilEspecialidades struct {
	Id                 int               `orm:"column(id_tr_perfil_especialidades);pk;auto"`
	IdPerfilesEmpleado *PerfilesEmpleado `orm:"column(id_perfiles_empleado);rel(fk)"`
	IdEspecialidades   *Especialidades   `orm:"column(id_especialidades);rel(fk)"`
	Activo             bool              `orm:"column(activo)"`
	CreadoEn           time.Time         `orm:"column(creado_en);type(timestamp without time zone);auto_now_add"`
	ActualizadoEn      time.Time         `orm:"column(actualizado_en);type(timestamp without time zone);auto_now"`
}

func (t *TrPerfilEspecialidades) TableName() string {
	return "tr_perfil_especialidades"
}

func init() {
	orm.RegisterModel(new(TrPerfilEspecialidades))
}

// AddTrPerfilEspecialidades insert a new TrPerfilEspecialidades into database and returns
// last inserted Id on success.
func AddTrPerfilEspecialidades(m *TrPerfilEspecialidades) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTrPerfilEspecialidadesById retrieves TrPerfilEspecialidades by Id. Returns error if
// Id doesn't exist
func GetTrPerfilEspecialidadesById(id int) (v *TrPerfilEspecialidades, err error) {
	o := orm.NewOrm()
	v = &TrPerfilEspecialidades{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "IderfilesEmpleado")
		o.LoadRelated(v, "IdEspecialidades")
		return v, nil
	}
	return nil, err
}

// GetAllTrPerfilEspecialidades retrieves all TrPerfilEspecialidades matches certain condition. Returns empty list if
// no records exist
func GetAllTrPerfilEspecialidades(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TrPerfilEspecialidades))
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

	var l []TrPerfilEspecialidades
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

// UpdateTrPerfilEspecialidades updates TrPerfilEspecialidades by Id and returns error if
// the record to be updated doesn't exist
func UpdateTrPerfilEspecialidadesById(m *TrPerfilEspecialidades) (err error) {
	o := orm.NewOrm()
	v := TrPerfilEspecialidades{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTrPerfilEspecialidades deletes TrPerfilEspecialidades by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTrPerfilEspecialidades(id int) (err error) {
	o := orm.NewOrm()
	v := TrPerfilEspecialidades{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TrPerfilEspecialidades{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
