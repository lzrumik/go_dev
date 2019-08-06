package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortunePayReportAttachmentDownloadTask struct {
	Id           int       `orm:"column(id);auto"`
	PayReportId  string    `orm:"column(pay_report_id);null" description:"fortune_pay_report表id"`
	OperatorName string    `orm:"column(operator_name);size(36);null" description:"操作人"`
	ProductType  int8      `orm:"column(product_type);null" description:"产品类型：1 私募类；2 非私募类"`
	CreateUserId int       `orm:"column(create_user_id)" description:"创建人/操作人"`
	CreateTime   time.Time `orm:"column(create_time);type(datetime)" description:"创建时间/操作时间"`
	State        int8      `orm:"column(state);null" description:"状态：1 进行中；2 完成"`
}

func (t *FortunePayReportAttachmentDownloadTask) TableName() string {
	return "fortune_pay_report_attachment_download_task"
}

func init() {
	orm.RegisterModel(new(FortunePayReportAttachmentDownloadTask))
}

// AddFortunePayReportAttachmentDownloadTask insert a new FortunePayReportAttachmentDownloadTask into database and returns
// last inserted Id on success.
func AddFortunePayReportAttachmentDownloadTask(m *FortunePayReportAttachmentDownloadTask) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortunePayReportAttachmentDownloadTaskById retrieves FortunePayReportAttachmentDownloadTask by Id. Returns error if
// Id doesn't exist
func GetFortunePayReportAttachmentDownloadTaskById(id int) (v *FortunePayReportAttachmentDownloadTask, err error) {
	o := orm.NewOrm()
	v = &FortunePayReportAttachmentDownloadTask{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortunePayReportAttachmentDownloadTask retrieves all FortunePayReportAttachmentDownloadTask matches certain condition. Returns empty list if
// no records exist
func GetAllFortunePayReportAttachmentDownloadTask(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortunePayReportAttachmentDownloadTask))
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

	var l []FortunePayReportAttachmentDownloadTask
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

// UpdateFortunePayReportAttachmentDownloadTask updates FortunePayReportAttachmentDownloadTask by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortunePayReportAttachmentDownloadTaskById(m *FortunePayReportAttachmentDownloadTask) (err error) {
	o := orm.NewOrm()
	v := FortunePayReportAttachmentDownloadTask{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortunePayReportAttachmentDownloadTask deletes FortunePayReportAttachmentDownloadTask by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortunePayReportAttachmentDownloadTask(id int) (err error) {
	o := orm.NewOrm()
	v := FortunePayReportAttachmentDownloadTask{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortunePayReportAttachmentDownloadTask{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
