package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneVideoRecordsHistory struct {
	Id            int       `orm:"column(id);auto"`
	VideoRecordId uint      `orm:"column(video_record_id)" description:"双录id"`
	HistoryId     int       `orm:"column(history_id)" description:"历史记录id"`
	AuditName     string    `orm:"column(audit_name);size(50)" description:"审核人姓名"`
	AuditTime     time.Time `orm:"column(audit_time);type(datetime)" description:"审批时间"`
	AuditStatus   int8      `orm:"column(audit_status)" description:"审批结果，0默认，1通过，2拒绝"`
	AuditReason   int8      `orm:"column(audit_reason)" description:"拒绝原因"`
	AuditMark     string    `orm:"column(audit_mark);size(255)" description:"备注"`
}

func (t *FortuneVideoRecordsHistory) TableName() string {
	return "fortune_video_records_history"
}

func init() {
	orm.RegisterModel(new(FortuneVideoRecordsHistory))
}

// AddFortuneVideoRecordsHistory insert a new FortuneVideoRecordsHistory into database and returns
// last inserted Id on success.
func AddFortuneVideoRecordsHistory(m *FortuneVideoRecordsHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneVideoRecordsHistoryById retrieves FortuneVideoRecordsHistory by Id. Returns error if
// Id doesn't exist
func GetFortuneVideoRecordsHistoryById(id int) (v *FortuneVideoRecordsHistory, err error) {
	o := orm.NewOrm()
	v = &FortuneVideoRecordsHistory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneVideoRecordsHistory retrieves all FortuneVideoRecordsHistory matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneVideoRecordsHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneVideoRecordsHistory))
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

	var l []FortuneVideoRecordsHistory
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

// UpdateFortuneVideoRecordsHistory updates FortuneVideoRecordsHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneVideoRecordsHistoryById(m *FortuneVideoRecordsHistory) (err error) {
	o := orm.NewOrm()
	v := FortuneVideoRecordsHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneVideoRecordsHistory deletes FortuneVideoRecordsHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneVideoRecordsHistory(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneVideoRecordsHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneVideoRecordsHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
