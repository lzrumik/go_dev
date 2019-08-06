package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerTask struct {
	Id                 int       `orm:"column(id);pk" description:"客户分配任务id"`
	AttachName         string    `orm:"column(attach_name);size(200)" description:"附件名称"`
	AttachPath         string    `orm:"column(attach_path);size(300)" description:"附件路径"`
	AttachType         int8      `orm:"column(attach_type)" description:"附件模板类型 1=base  2=number"`
	Total              int       `orm:"column(total)" description:"附件内容总数"`
	CustomerManagerIds string    `orm:"column(customer_manager_ids);size(1000)" description:"客户经理ids"`
	Status             int8      `orm:"column(status)" description:"分配状态 待分配=1、分配中=2、分配完成=3、分配失败=4"`
	RunTaskTime        int       `orm:"column(run_task_time)" description:"分配执行时间"`
	AnomalyMessage     string    `orm:"column(anomaly_message);size(500)" description:"异常原因"`
	TaskMessage        string    `orm:"column(task_message);size(500)" description:"分配信息"`
	CreatedUserId      string    `orm:"column(created_user_id);size(36)" description:"创建者id"`
	CreatedName        string    `orm:"column(created_name);size(30)" description:"创建者姓名"`
	DateEntered        time.Time `orm:"column(date_entered);type(datetime)" description:"创建日期"`
	Deleted            int8      `orm:"column(deleted)" description:"是否删除"`
}

func (t *FortuneCustomerTask) TableName() string {
	return "fortune_customer_task"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerTask))
}

// AddFortuneCustomerTask insert a new FortuneCustomerTask into database and returns
// last inserted Id on success.
func AddFortuneCustomerTask(m *FortuneCustomerTask) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerTaskById retrieves FortuneCustomerTask by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerTaskById(id int) (v *FortuneCustomerTask, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerTask{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerTask retrieves all FortuneCustomerTask matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerTask(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerTask))
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

	var l []FortuneCustomerTask
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

// UpdateFortuneCustomerTask updates FortuneCustomerTask by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerTaskById(m *FortuneCustomerTask) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerTask{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerTask deletes FortuneCustomerTask by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerTask(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerTask{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerTask{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
