package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneAppropriateApproveHistory struct {
	Id                int       `orm:"column(id);auto"`
	AppropriateInfoId int       `orm:"column(appropriate_info_id)"`
	HistoryId         int       `orm:"column(history_id)" description:"历史记录id"`
	UserId            string    `orm:"column(user_id);size(36)" description:"用户user_id"`
	SubmitTime        time.Time `orm:"column(submit_time);type(datetime)" description:"提交审核的时间"`
	SubmitName        string    `orm:"column(submit_name);size(50)" description:"提交人"`
	AuditName         string    `orm:"column(audit_name);size(50)" description:"审核人"`
	AddTime           time.Time `orm:"column(add_time);type(datetime)"`
	UpdateTime        time.Time `orm:"column(update_time);type(datetime)"`
	Status            int8      `orm:"column(status)" description:"状态 1审核通过,2审核拒绝"`
	Reason            int8      `orm:"column(reason)" description:"拒绝理由"`
	Mark              string    `orm:"column(mark);size(300)" description:"备注"`
	Deleted           int8      `orm:"column(deleted)"`
}

func (t *FortuneAppropriateApproveHistory) TableName() string {
	return "fortune_appropriate_approve_history"
}

func init() {
	orm.RegisterModel(new(FortuneAppropriateApproveHistory))
}

// AddFortuneAppropriateApproveHistory insert a new FortuneAppropriateApproveHistory into database and returns
// last inserted Id on success.
func AddFortuneAppropriateApproveHistory(m *FortuneAppropriateApproveHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneAppropriateApproveHistoryById retrieves FortuneAppropriateApproveHistory by Id. Returns error if
// Id doesn't exist
func GetFortuneAppropriateApproveHistoryById(id int) (v *FortuneAppropriateApproveHistory, err error) {
	o := orm.NewOrm()
	v = &FortuneAppropriateApproveHistory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneAppropriateApproveHistory retrieves all FortuneAppropriateApproveHistory matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneAppropriateApproveHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneAppropriateApproveHistory))
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

	var l []FortuneAppropriateApproveHistory
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

// UpdateFortuneAppropriateApproveHistory updates FortuneAppropriateApproveHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneAppropriateApproveHistoryById(m *FortuneAppropriateApproveHistory) (err error) {
	o := orm.NewOrm()
	v := FortuneAppropriateApproveHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneAppropriateApproveHistory deletes FortuneAppropriateApproveHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneAppropriateApproveHistory(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneAppropriateApproveHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneAppropriateApproveHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
