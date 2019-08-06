package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerCallPlan struct {
	Id                      int       `orm:"column(id);pk" description:"沟通计划id"`
	CustomerId              string    `orm:"column(customer_id);size(36)" description:"客户关系id"`
	IsAllDay                uint8     `orm:"column(is_all_day)" description:"是否全天日程 0:否 1:是"`
	PlanCallDate            time.Time `orm:"column(plan_call_date);type(datetime)" description:"开始时间  原预计沟通时间"`
	EndTime                 time.Time `orm:"column(end_time);type(datetime)" description:"结束时间"`
	CallType                string    `orm:"column(call_type);size(30)" description:"沟通方式 phone:电话 wechat:微信 accompany:陪访 open:开放日 visit:拜访 eat:吃饭 anniversary:纪念日 cash_back_day:回款日 birthday:生日 meeting:会议 other:其他"`
	Expect                  string    `orm:"column(expect);null" description:"预期"`
	CustomerManagerName     string    `orm:"column(customer_manager_name);size(30)" description:"客户经理姓名"`
	CustomerManagerUsername string    `orm:"column(customer_manager_username);size(30)" description:"客户经理账号"`
	HandleSituation         string    `orm:"column(handle_situation);null" description:"处理情况"`
	Status                  uint8     `orm:"column(status)" description:"状态 0:无 1:未完成 2:已完成 3:已取消"`
	DateEntered             time.Time `orm:"column(date_entered);type(datetime)"`
	CreatedUserId           string    `orm:"column(created_user_id);size(36)" description:"创建者"`
	DateModified            time.Time `orm:"column(date_modified);type(datetime);null"`
	ModifiedUserId          string    `orm:"column(modified_user_id);size(36)" description:"最后更新者"`
	Deleted                 int8      `orm:"column(deleted)" description:"是否删除"`
}

func (t *FortuneCustomerCallPlan) TableName() string {
	return "fortune_customer_call_plan"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerCallPlan))
}

// AddFortuneCustomerCallPlan insert a new FortuneCustomerCallPlan into database and returns
// last inserted Id on success.
func AddFortuneCustomerCallPlan(m *FortuneCustomerCallPlan) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerCallPlanById retrieves FortuneCustomerCallPlan by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerCallPlanById(id int) (v *FortuneCustomerCallPlan, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerCallPlan{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerCallPlan retrieves all FortuneCustomerCallPlan matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerCallPlan(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerCallPlan))
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

	var l []FortuneCustomerCallPlan
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

// UpdateFortuneCustomerCallPlan updates FortuneCustomerCallPlan by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerCallPlanById(m *FortuneCustomerCallPlan) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerCallPlan{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerCallPlan deletes FortuneCustomerCallPlan by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerCallPlan(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerCallPlan{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerCallPlan{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
