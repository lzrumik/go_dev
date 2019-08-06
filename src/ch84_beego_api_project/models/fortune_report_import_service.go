package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneReportImportService struct {
	Id                  int       `orm:"column(id);auto"`
	LoanId              string    `orm:"column(loan_id);size(36)" description:"项目id"`
	LoanPeriodId        string    `orm:"column(loan_period_id);size(36)" description:"起息日id，按逗号分隔"`
	LoanPeriodIdNum     int       `orm:"column(loan_period_id_num)" description:"起息日个数"`
	ValueDate           string    `orm:"column(value_date);size(200)" description:"起息日列表"`
	ServiceDiscountReal float64   `orm:"column(service_discount_real);digits(10);decimals(2)" description:"实收服务费(元)"`
	AccountingDate      time.Time `orm:"column(accounting_date);type(date)" description:"到账时间"`
	AttrPath            string    `orm:"column(attr_path);size(200)" description:"附件地址"`
	AttrName            string    `orm:"column(attr_name);size(200)" description:"附件名"`
	DateEntered         time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified        time.Time `orm:"column(date_modified);type(datetime);auto_now_add" description:"修改时间"`
	Deleted             int8      `orm:"column(deleted)" description:"是否删除 0否 1是"`
}

func (t *FortuneReportImportService) TableName() string {
	return "fortune_report_import_service"
}

func init() {
	orm.RegisterModel(new(FortuneReportImportService))
}

// AddFortuneReportImportService insert a new FortuneReportImportService into database and returns
// last inserted Id on success.
func AddFortuneReportImportService(m *FortuneReportImportService) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneReportImportServiceById retrieves FortuneReportImportService by Id. Returns error if
// Id doesn't exist
func GetFortuneReportImportServiceById(id int) (v *FortuneReportImportService, err error) {
	o := orm.NewOrm()
	v = &FortuneReportImportService{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneReportImportService retrieves all FortuneReportImportService matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneReportImportService(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneReportImportService))
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

	var l []FortuneReportImportService
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

// UpdateFortuneReportImportService updates FortuneReportImportService by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneReportImportServiceById(m *FortuneReportImportService) (err error) {
	o := orm.NewOrm()
	v := FortuneReportImportService{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneReportImportService deletes FortuneReportImportService by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneReportImportService(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneReportImportService{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneReportImportService{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
