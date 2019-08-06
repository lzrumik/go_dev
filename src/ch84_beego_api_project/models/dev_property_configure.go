package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type DevPropertyConfigure struct {
	Id             int       `orm:"column(id);auto" description:"自增主健"`
	ExternType     string    `orm:"column(extern_type);size(255)" description:"外部类型"`
	Type           string    `orm:"column(type);size(255)" description:"类型"`
	Name           string    `orm:"column(name);size(255)" description:"名称"`
	Value          string    `orm:"column(value)" description:"值"`
	Description    string    `orm:"column(description)" description:"描述"`
	Sort           int       `orm:"column(sort)" description:"排序"`
	Status         uint      `orm:"column(status)" description:"状态 1:正常 2:废弃"`
	CreatedUserId  string    `orm:"column(created_user_id);size(36)" description:"创建人ID"`
	ModifiedUserId string    `orm:"column(modified_user_id);size(36)" description:"修改人ID"`
	DateEntered    time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified   time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	Deleted        uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *DevPropertyConfigure) TableName() string {
	return "dev_property_configure"
}

func init() {
	orm.RegisterModel(new(DevPropertyConfigure))
}

// AddDevPropertyConfigure insert a new DevPropertyConfigure into database and returns
// last inserted Id on success.
func AddDevPropertyConfigure(m *DevPropertyConfigure) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDevPropertyConfigureById retrieves DevPropertyConfigure by Id. Returns error if
// Id doesn't exist
func GetDevPropertyConfigureById(id int) (v *DevPropertyConfigure, err error) {
	o := orm.NewOrm()
	v = &DevPropertyConfigure{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDevPropertyConfigure retrieves all DevPropertyConfigure matches certain condition. Returns empty list if
// no records exist
func GetAllDevPropertyConfigure(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DevPropertyConfigure))
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

	var l []DevPropertyConfigure
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

// UpdateDevPropertyConfigure updates DevPropertyConfigure by Id and returns error if
// the record to be updated doesn't exist
func UpdateDevPropertyConfigureById(m *DevPropertyConfigure) (err error) {
	o := orm.NewOrm()
	v := DevPropertyConfigure{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDevPropertyConfigure deletes DevPropertyConfigure by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDevPropertyConfigure(id int) (err error) {
	o := orm.NewOrm()
	v := DevPropertyConfigure{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DevPropertyConfigure{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
