package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerSubscribeAllotHistory struct {
	Id                  int       `orm:"column(id);auto"`
	CustomerSubscribeId int       `orm:"column(customer_subscribe_id)" description:"预约绑定客户id"`
	LastVersion         string    `orm:"column(last_version);size(2000)" description:"最后版本数据"`
	BindSamId           string    `orm:"column(bind_sam_id);size(36)" description:"分配的客户经理ID"`
	AllotInfo           string    `orm:"column(allot_info);size(200)" description:"分配信息"`
	CreateUserId        string    `orm:"column(create_user_id);size(36)" description:"分配人id"`
	CreateUserFullName  string    `orm:"column(create_user_full_name);size(20)" description:"分配人姓名"`
	DateEntered         time.Time `orm:"column(date_entered);type(datetime);null"`
}

func (t *FortuneCustomerSubscribeAllotHistory) TableName() string {
	return "fortune_customer_subscribe_allot_history"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerSubscribeAllotHistory))
}

// AddFortuneCustomerSubscribeAllotHistory insert a new FortuneCustomerSubscribeAllotHistory into database and returns
// last inserted Id on success.
func AddFortuneCustomerSubscribeAllotHistory(m *FortuneCustomerSubscribeAllotHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerSubscribeAllotHistoryById retrieves FortuneCustomerSubscribeAllotHistory by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerSubscribeAllotHistoryById(id int) (v *FortuneCustomerSubscribeAllotHistory, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerSubscribeAllotHistory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerSubscribeAllotHistory retrieves all FortuneCustomerSubscribeAllotHistory matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerSubscribeAllotHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerSubscribeAllotHistory))
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

	var l []FortuneCustomerSubscribeAllotHistory
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

// UpdateFortuneCustomerSubscribeAllotHistory updates FortuneCustomerSubscribeAllotHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerSubscribeAllotHistoryById(m *FortuneCustomerSubscribeAllotHistory) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerSubscribeAllotHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerSubscribeAllotHistory deletes FortuneCustomerSubscribeAllotHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerSubscribeAllotHistory(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerSubscribeAllotHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerSubscribeAllotHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
