package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TestTb2 struct {
	Id     int       `orm:"column(id);auto"`
	Str1   string    `orm:"column(str_1);size(20)"`
	Str2   string    `orm:"column(str_2);size(20)"`
	Str3   string    `orm:"column(str_3);size(36)"`
	Str4   string    `orm:"column(str_4);size(36)"`
	Int1   int       `orm:"column(int_1)"`
	Int2   int       `orm:"column(int_2)"`
	Int3   int64     `orm:"column(int_3)"`
	Int4   int64     `orm:"column(int_4)"`
	Float1 float32   `orm:"column(float_1)"`
	Float2 float64   `orm:"column(float_2)"`
	Float3 float64   `orm:"column(float_3);digits(10);decimals(2)"`
	Float4 float64   `orm:"column(float_4);digits(10);decimals(2)"`
	Time1  time.Time `orm:"column(time_1);type(datetime)"`
	Time2  time.Time `orm:"column(time_2);type(timestamp);auto_now"`
}

func (t *TestTb2) TableName() string {
	return "test_tb2"
}

func init() {
	orm.RegisterModel(new(TestTb2))
}

// AddTestTb2 insert a new TestTb2 into database and returns
// last inserted Id on success.
func AddTestTb2(m *TestTb2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTestTb2ById retrieves TestTb2 by Id. Returns error if
// Id doesn't exist
func GetTestTb2ById(id int) (v *TestTb2, err error) {
	o := orm.NewOrm()
	v = &TestTb2{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTestTb2 retrieves all TestTb2 matches certain condition. Returns empty list if
// no records exist
func GetAllTestTb2(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TestTb2))
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

	var l []TestTb2
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

// UpdateTestTb2 updates TestTb2 by Id and returns error if
// the record to be updated doesn't exist
func UpdateTestTb2ById(m *TestTb2) (err error) {
	o := orm.NewOrm()
	v := TestTb2{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTestTb2 deletes TestTb2 by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTestTb2(id int) (err error) {
	o := orm.NewOrm()
	v := TestTb2{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TestTb2{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
