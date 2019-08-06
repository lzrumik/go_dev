package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneClassifyTask struct {
	Id           int       `orm:"column(classify_id);auto" description:"自增id"`
	ClassifyName string    `orm:"column(classify_name);size(255)" description:"分类名名"`
	ClassifyInfo string    `orm:"column(classify_info);size(10000)" description:"任务信息json"`
	SetCount     int       `orm:"column(set_count)" description:"集合个数"`
	TaskStatus   uint8     `orm:"column(task_status)" description:"任务状态 1未执行 2执行中  3已执行 4执行失败 "`
	TaskType     uint8     `orm:"column(task_type)" description:"任务状态 1普通筛选 4周期性任务 "`
	CreateTime   time.Time `orm:"column(create_time);type(timestamp)" description:"创建时间"`
	ModifyTime   time.Time `orm:"column(modify_time);type(timestamp)" description:"更新时间"`
	CreateUserId string    `orm:"column(create_user_id);size(36)" description:"创建人id"`
	RecordCount  int       `orm:"column(record_count)" description:"记录行数"`
	ExtraData    string    `orm:"column(extra_data);size(2000)" description:"备注信息"`
	RandKey      string    `orm:"column(rand_key);size(50)" description:"随机串"`
	Memo         string    `orm:"column(memo);size(1000)" description:"备注信息"`
	CostTime     uint      `orm:"column(cost_time)" description:"花费时间"`
	ExcelPath    string    `orm:"column(excel_path);size(200)" description:"生成的excel地址"`
	Deleted      int8      `orm:"column(deleted)" description:"是否被删除 0否 1是"`
}

func (t *FortuneClassifyTask) TableName() string {
	return "fortune_classify_task"
}

func init() {
	orm.RegisterModel(new(FortuneClassifyTask))
}

// AddFortuneClassifyTask insert a new FortuneClassifyTask into database and returns
// last inserted Id on success.
func AddFortuneClassifyTask(m *FortuneClassifyTask) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneClassifyTaskById retrieves FortuneClassifyTask by Id. Returns error if
// Id doesn't exist
func GetFortuneClassifyTaskById(id int) (v *FortuneClassifyTask, err error) {
	o := orm.NewOrm()
	v = &FortuneClassifyTask{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneClassifyTask retrieves all FortuneClassifyTask matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneClassifyTask(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneClassifyTask))
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

	var l []FortuneClassifyTask
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

// UpdateFortuneClassifyTask updates FortuneClassifyTask by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneClassifyTaskById(m *FortuneClassifyTask) (err error) {
	o := orm.NewOrm()
	v := FortuneClassifyTask{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneClassifyTask deletes FortuneClassifyTask by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneClassifyTask(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneClassifyTask{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneClassifyTask{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
