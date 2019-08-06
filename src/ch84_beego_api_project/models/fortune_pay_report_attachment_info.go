package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortunePayReportAttachmentInfo struct {
	Id           int       `orm:"column(id);auto"`
	PayReportId  int       `orm:"column(pay_report_id)" description:"对应表的fortune_pay_report_id"`
	Type         uint8     `orm:"column(type)" description:"类型 1.非私募  2，私募"`
	FilePath     string    `orm:"column(file_path);size(200)" description:"文件路径"`
	FileName     string    `orm:"column(file_name);size(200)" description:"文件名"`
	CreateUserId string    `orm:"column(create_user_id);size(36)"`
	DateEntered  time.Time `orm:"column(date_entered);type(datetime);null"`
	DateModified time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	Deleted      uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortunePayReportAttachmentInfo) TableName() string {
	return "fortune_pay_report_attachment_info"
}

func init() {
	orm.RegisterModel(new(FortunePayReportAttachmentInfo))
}

// AddFortunePayReportAttachmentInfo insert a new FortunePayReportAttachmentInfo into database and returns
// last inserted Id on success.
func AddFortunePayReportAttachmentInfo(m *FortunePayReportAttachmentInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortunePayReportAttachmentInfoById retrieves FortunePayReportAttachmentInfo by Id. Returns error if
// Id doesn't exist
func GetFortunePayReportAttachmentInfoById(id int) (v *FortunePayReportAttachmentInfo, err error) {
	o := orm.NewOrm()
	v = &FortunePayReportAttachmentInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortunePayReportAttachmentInfo retrieves all FortunePayReportAttachmentInfo matches certain condition. Returns empty list if
// no records exist
func GetAllFortunePayReportAttachmentInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortunePayReportAttachmentInfo))
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

	var l []FortunePayReportAttachmentInfo
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

// UpdateFortunePayReportAttachmentInfo updates FortunePayReportAttachmentInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortunePayReportAttachmentInfoById(m *FortunePayReportAttachmentInfo) (err error) {
	o := orm.NewOrm()
	v := FortunePayReportAttachmentInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortunePayReportAttachmentInfo deletes FortunePayReportAttachmentInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortunePayReportAttachmentInfo(id int) (err error) {
	o := orm.NewOrm()
	v := FortunePayReportAttachmentInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortunePayReportAttachmentInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
