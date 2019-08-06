package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerCallHistory20180201 struct {
	Id                      int       `orm:"column(id);pk" description:"沟通历史id"`
	CustomerId              string    `orm:"column(customer_id);size(36)" description:"客户关系id"`
	CallDate                time.Time `orm:"column(call_date);type(datetime)" description:"沟通时间"`
	CallType                string    `orm:"column(call_type);size(30)" description:"沟通方式（电话=phone、微信=wechat、陪访=accompany、开放日=open）"`
	CallOnWorker            string    `orm:"column(call_on_worker);size(200)" description:"陪访人员"`
	CallContent             string    `orm:"column(call_content);size(2000)" description:"沟通内容"`
	CustomerManagerName     string    `orm:"column(customer_manager_name);size(30)" description:"客户经理姓名"`
	CustomerManagerUsername string    `orm:"column(customer_manager_username);size(30)" description:"客户经理账号"`
	DateEntered             time.Time `orm:"column(date_entered);type(datetime)"`
	CreatedUserId           string    `orm:"column(created_user_id);size(36)" description:"创建者"`
	Deleted                 int8      `orm:"column(deleted)" description:"是否删除"`
}

func (t *FortuneCustomerCallHistory20180201) TableName() string {
	return "fortune_customer_call_history_20180201"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerCallHistory20180201))
}

// AddFortuneCustomerCallHistory20180201 insert a new FortuneCustomerCallHistory20180201 into database and returns
// last inserted Id on success.
func AddFortuneCustomerCallHistory20180201(m *FortuneCustomerCallHistory20180201) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerCallHistory20180201ById retrieves FortuneCustomerCallHistory20180201 by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerCallHistory20180201ById(id int) (v *FortuneCustomerCallHistory20180201, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerCallHistory20180201{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerCallHistory20180201 retrieves all FortuneCustomerCallHistory20180201 matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerCallHistory20180201(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerCallHistory20180201))
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

	var l []FortuneCustomerCallHistory20180201
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

// UpdateFortuneCustomerCallHistory20180201 updates FortuneCustomerCallHistory20180201 by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerCallHistory20180201ById(m *FortuneCustomerCallHistory20180201) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerCallHistory20180201{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerCallHistory20180201 deletes FortuneCustomerCallHistory20180201 by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerCallHistory20180201(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerCallHistory20180201{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerCallHistory20180201{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
