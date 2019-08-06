package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerBookHistory struct {
	Id      int       `orm:"column(id);auto"`
	BookId  string    `orm:"column(book_id);size(36)" description:"用户的ID"`
	Status  int8      `orm:"column(status)" description:"0未处理,1跟进中,2已完成"`
	OpUser  string    `orm:"column(op_user);size(20)" description:"操作人"`
	AddTime time.Time `orm:"column(add_time);type(datetime)"`
}

func (t *FortuneCustomerBookHistory) TableName() string {
	return "fortune_customer_book_history"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerBookHistory))
}

// AddFortuneCustomerBookHistory insert a new FortuneCustomerBookHistory into database and returns
// last inserted Id on success.
func AddFortuneCustomerBookHistory(m *FortuneCustomerBookHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerBookHistoryById retrieves FortuneCustomerBookHistory by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerBookHistoryById(id int) (v *FortuneCustomerBookHistory, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerBookHistory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerBookHistory retrieves all FortuneCustomerBookHistory matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerBookHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerBookHistory))
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

	var l []FortuneCustomerBookHistory
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

// UpdateFortuneCustomerBookHistory updates FortuneCustomerBookHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerBookHistoryById(m *FortuneCustomerBookHistory) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerBookHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerBookHistory deletes FortuneCustomerBookHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerBookHistory(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerBookHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerBookHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
