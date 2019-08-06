package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortunePayReportZipArchive struct {
	Id           int       `orm:"column(id);auto"`
	PayReportId  int       `orm:"column(pay_report_id)" description:"报单ID"`
	Type         int8      `orm:"column(type)" description:"类型 1:非私募(财富) 2:私募"`
	Status       int8      `orm:"column(status)" description:"状态 1:初始 2:处理中 3:已处理(打包完成)"`
	DateEntered  time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	Deleted      int8      `orm:"column(deleted)"`
}

func (t *FortunePayReportZipArchive) TableName() string {
	return "fortune_pay_report_zip_archive"
}

func init() {
	orm.RegisterModel(new(FortunePayReportZipArchive))
}

// AddFortunePayReportZipArchive insert a new FortunePayReportZipArchive into database and returns
// last inserted Id on success.
func AddFortunePayReportZipArchive(m *FortunePayReportZipArchive) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortunePayReportZipArchiveById retrieves FortunePayReportZipArchive by Id. Returns error if
// Id doesn't exist
func GetFortunePayReportZipArchiveById(id int) (v *FortunePayReportZipArchive, err error) {
	o := orm.NewOrm()
	v = &FortunePayReportZipArchive{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortunePayReportZipArchive retrieves all FortunePayReportZipArchive matches certain condition. Returns empty list if
// no records exist
func GetAllFortunePayReportZipArchive(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortunePayReportZipArchive))
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

	var l []FortunePayReportZipArchive
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

// UpdateFortunePayReportZipArchive updates FortunePayReportZipArchive by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortunePayReportZipArchiveById(m *FortunePayReportZipArchive) (err error) {
	o := orm.NewOrm()
	v := FortunePayReportZipArchive{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortunePayReportZipArchive deletes FortunePayReportZipArchive by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortunePayReportZipArchive(id int) (err error) {
	o := orm.NewOrm()
	v := FortunePayReportZipArchive{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortunePayReportZipArchive{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
