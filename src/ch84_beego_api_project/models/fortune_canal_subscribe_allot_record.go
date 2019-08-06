package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCanalSubscribeAllotRecord struct {
	Id                    int       `orm:"column(id);auto"`
	CanalSubscribeId      int       `orm:"column(canal_subscribe_id)" description:"渠道预约id"`
	UserId                string    `orm:"column(user_id);size(36)" description:"用户ID"`
	AllotOperateId        string    `orm:"column(allot_operate_id);size(50)" description:"分配操作者id"`
	AllotOperateName      string    `orm:"column(allot_operate_name);size(50)" description:"分配操作者姓名"`
	BinSamId              string    `orm:"column(bin_sam_id);size(50)" description:"分配的客户经理id"`
	CanalSubscribeVersion string    `orm:"column(canal_subscribe_version);size(2000)" description:"分配信息"`
	AllotStatus           int8      `orm:"column(allot_status)" description:"状态 0=未分配、1=已分配、2=分配失败"`
	AllotInfo             string    `orm:"column(allot_info);size(200)" description:"分配信息"`
	Deleted               int8      `orm:"column(deleted)" description:"是否删除"`
	DateEntered           time.Time `orm:"column(date_entered);type(datetime);null"`
}

func (t *FortuneCanalSubscribeAllotRecord) TableName() string {
	return "fortune_canal_subscribe_allot_record"
}

func init() {
	orm.RegisterModel(new(FortuneCanalSubscribeAllotRecord))
}

// AddFortuneCanalSubscribeAllotRecord insert a new FortuneCanalSubscribeAllotRecord into database and returns
// last inserted Id on success.
func AddFortuneCanalSubscribeAllotRecord(m *FortuneCanalSubscribeAllotRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCanalSubscribeAllotRecordById retrieves FortuneCanalSubscribeAllotRecord by Id. Returns error if
// Id doesn't exist
func GetFortuneCanalSubscribeAllotRecordById(id int) (v *FortuneCanalSubscribeAllotRecord, err error) {
	o := orm.NewOrm()
	v = &FortuneCanalSubscribeAllotRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCanalSubscribeAllotRecord retrieves all FortuneCanalSubscribeAllotRecord matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCanalSubscribeAllotRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCanalSubscribeAllotRecord))
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

	var l []FortuneCanalSubscribeAllotRecord
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

// UpdateFortuneCanalSubscribeAllotRecord updates FortuneCanalSubscribeAllotRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCanalSubscribeAllotRecordById(m *FortuneCanalSubscribeAllotRecord) (err error) {
	o := orm.NewOrm()
	v := FortuneCanalSubscribeAllotRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCanalSubscribeAllotRecord deletes FortuneCanalSubscribeAllotRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCanalSubscribeAllotRecord(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCanalSubscribeAllotRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCanalSubscribeAllotRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
