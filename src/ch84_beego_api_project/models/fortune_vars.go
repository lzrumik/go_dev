package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneVars struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(100)" description:"变量名"`
	Topic      string    `orm:"column(topic);size(100)" description:"变量的标识"`
	Content    string    `orm:"column(content);size(1000)" description:"变量内容"`
	SendMethod int8      `orm:"column(send_method)" description:"1短信，2邮件"`
	SendType   int8      `orm:"column(send_type)" description:"1自动，2手动"`
	Status     int8      `orm:"column(status)" description:"1启用，2禁用"`
	AddTime    time.Time `orm:"column(add_time);type(datetime)"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime)"`
	AddUser    string    `orm:"column(add_user);size(50)"`
	UpdateUser string    `orm:"column(update_user);size(50)"`
}

func (t *FortuneVars) TableName() string {
	return "fortune_vars"
}

func init() {
	orm.RegisterModel(new(FortuneVars))
}

// AddFortuneVars insert a new FortuneVars into database and returns
// last inserted Id on success.
func AddFortuneVars(m *FortuneVars) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneVarsById retrieves FortuneVars by Id. Returns error if
// Id doesn't exist
func GetFortuneVarsById(id int) (v *FortuneVars, err error) {
	o := orm.NewOrm()
	v = &FortuneVars{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneVars retrieves all FortuneVars matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneVars(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneVars))
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

	var l []FortuneVars
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

// UpdateFortuneVars updates FortuneVars by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneVarsById(m *FortuneVars) (err error) {
	o := orm.NewOrm()
	v := FortuneVars{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneVars deletes FortuneVars by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneVars(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneVars{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneVars{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
