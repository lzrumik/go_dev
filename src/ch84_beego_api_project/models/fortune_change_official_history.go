package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneChangeOfficialHistory struct {
	Id           int       `orm:"column(id);pk"`
	CustomerId   string    `orm:"column(customer_id);size(36)" description:"客户ID"`
	SamId        string    `orm:"column(sam_id);size(36)" description:"客户经理ID"`
	Status       uint8     `orm:"column(status)" description:"状态 0:无 1:锁定 2:正式 3:解绑"`
	DateEntered  time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified time.Time `orm:"column(date_modified);type(datetime);null" description:"修改时间"`
	Deleted      uint8     `orm:"column(deleted)"`
}

func (t *FortuneChangeOfficialHistory) TableName() string {
	return "fortune_change_official_history"
}

func init() {
	orm.RegisterModel(new(FortuneChangeOfficialHistory))
}

// AddFortuneChangeOfficialHistory insert a new FortuneChangeOfficialHistory into database and returns
// last inserted Id on success.
func AddFortuneChangeOfficialHistory(m *FortuneChangeOfficialHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneChangeOfficialHistoryById retrieves FortuneChangeOfficialHistory by Id. Returns error if
// Id doesn't exist
func GetFortuneChangeOfficialHistoryById(id int) (v *FortuneChangeOfficialHistory, err error) {
	o := orm.NewOrm()
	v = &FortuneChangeOfficialHistory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneChangeOfficialHistory retrieves all FortuneChangeOfficialHistory matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneChangeOfficialHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneChangeOfficialHistory))
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

	var l []FortuneChangeOfficialHistory
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

// UpdateFortuneChangeOfficialHistory updates FortuneChangeOfficialHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneChangeOfficialHistoryById(m *FortuneChangeOfficialHistory) (err error) {
	o := orm.NewOrm()
	v := FortuneChangeOfficialHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneChangeOfficialHistory deletes FortuneChangeOfficialHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneChangeOfficialHistory(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneChangeOfficialHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneChangeOfficialHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
