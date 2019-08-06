package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortunePlanCallMsgHistory struct {
	Id           int       `orm:"column(id);pk" description:"主键id"`
	PlanCallDate time.Time `orm:"column(plan_call_date);type(datetime)" description:"沟通计划时间"`
	Mobile       string    `orm:"column(mobile);size(20)" description:"客户经理手机号"`
	SendContent  string    `orm:"column(send_content);size(200)" description:"短信内容"`
	SendStatus   int8      `orm:"column(send_status)" description:"短信发送状态 1=发送成功 2=发送失败"`
	SendReason   string    `orm:"column(send_reason);size(100)" description:"发送状态描述"`
	DateEntered  time.Time `orm:"column(date_entered);type(datetime)" description:"发送时间"`
}

func (t *FortunePlanCallMsgHistory) TableName() string {
	return "fortune_plan_call_msg_history"
}

func init() {
	orm.RegisterModel(new(FortunePlanCallMsgHistory))
}

// AddFortunePlanCallMsgHistory insert a new FortunePlanCallMsgHistory into database and returns
// last inserted Id on success.
func AddFortunePlanCallMsgHistory(m *FortunePlanCallMsgHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortunePlanCallMsgHistoryById retrieves FortunePlanCallMsgHistory by Id. Returns error if
// Id doesn't exist
func GetFortunePlanCallMsgHistoryById(id int) (v *FortunePlanCallMsgHistory, err error) {
	o := orm.NewOrm()
	v = &FortunePlanCallMsgHistory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortunePlanCallMsgHistory retrieves all FortunePlanCallMsgHistory matches certain condition. Returns empty list if
// no records exist
func GetAllFortunePlanCallMsgHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortunePlanCallMsgHistory))
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

	var l []FortunePlanCallMsgHistory
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

// UpdateFortunePlanCallMsgHistory updates FortunePlanCallMsgHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortunePlanCallMsgHistoryById(m *FortunePlanCallMsgHistory) (err error) {
	o := orm.NewOrm()
	v := FortunePlanCallMsgHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortunePlanCallMsgHistory deletes FortunePlanCallMsgHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortunePlanCallMsgHistory(id int) (err error) {
	o := orm.NewOrm()
	v := FortunePlanCallMsgHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortunePlanCallMsgHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
