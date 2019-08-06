package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type FortuneClassifyTaskRecordInfo201906 struct {
	Id         int    `orm:"column(id);auto"`
	Uid        int    `orm:"column(uid)" description:"int型uid"`
	ClassifyId int    `orm:"column(classify_id)" description:"分类ID"`
	AccFields  string `orm:"column(acc_fields);size(2000)" description:"json格式的数据 便于动态变更结果字段"`
	Deleted    int8   `orm:"column(deleted)" description:"是否删除"`
}

func (t *FortuneClassifyTaskRecordInfo201906) TableName() string {
	return "fortune_classify_task_record_info_201906"
}

func init() {
	orm.RegisterModel(new(FortuneClassifyTaskRecordInfo201906))
}

// AddFortuneClassifyTaskRecordInfo201906 insert a new FortuneClassifyTaskRecordInfo201906 into database and returns
// last inserted Id on success.
func AddFortuneClassifyTaskRecordInfo201906(m *FortuneClassifyTaskRecordInfo201906) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneClassifyTaskRecordInfo201906ById retrieves FortuneClassifyTaskRecordInfo201906 by Id. Returns error if
// Id doesn't exist
func GetFortuneClassifyTaskRecordInfo201906ById(id int) (v *FortuneClassifyTaskRecordInfo201906, err error) {
	o := orm.NewOrm()
	v = &FortuneClassifyTaskRecordInfo201906{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneClassifyTaskRecordInfo201906 retrieves all FortuneClassifyTaskRecordInfo201906 matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneClassifyTaskRecordInfo201906(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneClassifyTaskRecordInfo201906))
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

	var l []FortuneClassifyTaskRecordInfo201906
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

// UpdateFortuneClassifyTaskRecordInfo201906 updates FortuneClassifyTaskRecordInfo201906 by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneClassifyTaskRecordInfo201906ById(m *FortuneClassifyTaskRecordInfo201906) (err error) {
	o := orm.NewOrm()
	v := FortuneClassifyTaskRecordInfo201906{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneClassifyTaskRecordInfo201906 deletes FortuneClassifyTaskRecordInfo201906 by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneClassifyTaskRecordInfo201906(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneClassifyTaskRecordInfo201906{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneClassifyTaskRecordInfo201906{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}