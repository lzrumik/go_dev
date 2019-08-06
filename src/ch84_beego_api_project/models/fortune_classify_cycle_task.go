package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneClassifyCycleTask struct {
	Id           int       `orm:"column(task_id);auto" description:"周期性任务id"`
	TaskName     string    `orm:"column(task_name);size(255)" description:"分类名名"`
	ClassifyInfo string    `orm:"column(classify_info);size(10000)" description:"任务信息json"`
	TaskStatus   uint8     `orm:"column(task_status)" description:"任务状态 1未执行 2执行中  3已执行 4执行失败 "`
	CreateTime   time.Time `orm:"column(create_time);type(timestamp)" description:"创建时间"`
	StartDate    time.Time `orm:"column(start_date);type(date)" description:"开始日期"`
	EndDate      time.Time `orm:"column(end_date);type(date)" description:"结束日期"`
	CreateUserId string    `orm:"column(create_user_id);size(36)" description:"创建人id"`
	ExtraData    string    `orm:"column(extra_data);size(2000)" description:"备注信息"`
	IsRepetitive int8      `orm:"column(is_repetitive)" description:"是否排重  0 排重 1不排重"`
	Notify       string    `orm:"column(notify);size(30)" description:"提醒类型"`
	NotifyDate   time.Time `orm:"column(notify_date);type(date)" description:"结束日期"`
	Memo         string    `orm:"column(memo);size(1000)" description:"备注信息"`
	Deleted      int8      `orm:"column(deleted)" description:"是否被删除 0否 1是"`
}

func (t *FortuneClassifyCycleTask) TableName() string {
	return "fortune_classify_cycle_task"
}

func init() {
	orm.RegisterModel(new(FortuneClassifyCycleTask))
}

// AddFortuneClassifyCycleTask insert a new FortuneClassifyCycleTask into database and returns
// last inserted Id on success.
func AddFortuneClassifyCycleTask(m *FortuneClassifyCycleTask) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneClassifyCycleTaskById retrieves FortuneClassifyCycleTask by Id. Returns error if
// Id doesn't exist
func GetFortuneClassifyCycleTaskById(id int) (v *FortuneClassifyCycleTask, err error) {
	o := orm.NewOrm()
	v = &FortuneClassifyCycleTask{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneClassifyCycleTask retrieves all FortuneClassifyCycleTask matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneClassifyCycleTask(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneClassifyCycleTask))
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

	var l []FortuneClassifyCycleTask
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

// UpdateFortuneClassifyCycleTask updates FortuneClassifyCycleTask by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneClassifyCycleTaskById(m *FortuneClassifyCycleTask) (err error) {
	o := orm.NewOrm()
	v := FortuneClassifyCycleTask{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneClassifyCycleTask deletes FortuneClassifyCycleTask by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneClassifyCycleTask(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneClassifyCycleTask{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneClassifyCycleTask{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
