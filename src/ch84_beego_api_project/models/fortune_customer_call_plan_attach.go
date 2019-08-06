package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerCallPlanAttach struct {
	Id             int       `orm:"column(id);pk" description:"ID"`
	PlanId         string    `orm:"column(plan_id);size(36)" description:"订单ID"`
	Type           uint8     `orm:"column(type)" description:"附件类型 1:日程附件 2:日程处理情况附件"`
	FilePath       string    `orm:"column(file_path);size(100)" description:"文件路径"`
	FileName       string    `orm:"column(file_name);size(100)" description:"文件名"`
	CreatedUserId  string    `orm:"column(created_user_id);size(36)"`
	ModifiedUserId string    `orm:"column(modified_user_id);size(36)"`
	DateEntered    time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified   time.Time `orm:"column(date_modified);type(datetime)"`
	Deleted        uint8     `orm:"column(deleted)" description:"0:正常 1:删除"`
}

func (t *FortuneCustomerCallPlanAttach) TableName() string {
	return "fortune_customer_call_plan_attach"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerCallPlanAttach))
}

// AddFortuneCustomerCallPlanAttach insert a new FortuneCustomerCallPlanAttach into database and returns
// last inserted Id on success.
func AddFortuneCustomerCallPlanAttach(m *FortuneCustomerCallPlanAttach) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerCallPlanAttachById retrieves FortuneCustomerCallPlanAttach by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerCallPlanAttachById(id int) (v *FortuneCustomerCallPlanAttach, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerCallPlanAttach{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerCallPlanAttach retrieves all FortuneCustomerCallPlanAttach matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerCallPlanAttach(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerCallPlanAttach))
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

	var l []FortuneCustomerCallPlanAttach
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

// UpdateFortuneCustomerCallPlanAttach updates FortuneCustomerCallPlanAttach by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerCallPlanAttachById(m *FortuneCustomerCallPlanAttach) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerCallPlanAttach{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerCallPlanAttach deletes FortuneCustomerCallPlanAttach by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerCallPlanAttach(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerCallPlanAttach{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerCallPlanAttach{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
