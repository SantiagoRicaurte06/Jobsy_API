package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type ReporteMensajes struct {
	Id            int       `orm:"column(id_reporte_mensajes);pk;auto"`
	IdReportes    *Reportes `orm:"column(id_reportes);rel(fk)"`
	IdRemitente   int       `orm:"column(id_remitente)"`
	Mensaje       string    `orm:"column(mensaje)"`
	ArchivoUrl    string    `orm:"column(archivo_url);null"`
	Activo        bool      `orm:"column(activo)"`
	CreadoEn      time.Time `orm:"column(creado_en);type(timestamp without time zone);auto_now_add"`
	ActualizadoEn time.Time `orm:"column(actualizado_en);type(timestamp without time zone);auto_now"`
}

func (t *ReporteMensajes) TableName() string {
	return "reporte_mensajes"
}

func init() {
	orm.RegisterModel(new(ReporteMensajes))
}

// AddReporteMensajes insert a new ReporteMensajes into database and returns
// last inserted Id on success.
func AddReporteMensajes(m *ReporteMensajes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReporteMensajesById retrieves ReporteMensajes by Id. Returns error if
// Id doesn't exist
func GetReporteMensajesById(id int) (v *ReporteMensajes, err error) {
	o := orm.NewOrm()
	v = &ReporteMensajes{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "IdReportes")
		o.LoadRelated(v, "IdRemitente")
		return v, nil
	}
	return nil, err
}

// GetAllReporteMensajes retrieves all ReporteMensajes matches certain condition. Returns empty list if
// no records exist
func GetAllReporteMensajes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ReporteMensajes)).RelatedSel()
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

	var l []ReporteMensajes
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

// UpdateReporteMensajes updates ReporteMensajes by Id and returns error if
// the record to be updated doesn't exist
func UpdateReporteMensajesById(m *ReporteMensajes) (err error) {
	o := orm.NewOrm()
	v := ReporteMensajes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReporteMensajes deletes ReporteMensajes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReporteMensajes(id int) (err error) {
	o := orm.NewOrm()
	v := ReporteMensajes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ReporteMensajes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
