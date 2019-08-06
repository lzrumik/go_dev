package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCompany struct {
	Id                     int       `orm:"column(id);auto"`
	CompanyName            string    `orm:"column(company_name);size(100)" description:"企业名称"`
	CompanyNo              string    `orm:"column(company_no);size(50)" description:"企业编号"`
	CorporateName          string    `orm:"column(corporate_name);size(50)" description:"法人代表"`
	SignatureNo            string    `orm:"column(signature_no);size(50)" description:"电子签章码"`
	Keyword                string    `orm:"column(keyword);size(50)" description:"关键字"`
	DateExpired            time.Time `orm:"column(date_expired);type(datetime)" description:"到期时间"`
	CorporateSignatureName string    `orm:"column(corporate_signature_name);size(100)" description:"法人签章名"`
	CorporateSignatureNo   string    `orm:"column(corporate_signature_no);size(50)" description:"法人电子签章码"`
	CorporateKeyword       string    `orm:"column(corporate_keyword);size(50)" description:"法人关键字"`
	CorporateDateExpired   time.Time `orm:"column(corporate_date_expired);type(datetime)" description:"法人信息过期时间"`
	CreatedUserId          string    `orm:"column(created_user_id);size(36)"`
	DateEntered            time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified           time.Time `orm:"column(date_modified);type(datetime)"`
	Deleted                uint8     `orm:"column(deleted)"`
}

func (t *FortuneCompany) TableName() string {
	return "fortune_company"
}

func init() {
	orm.RegisterModel(new(FortuneCompany))
}

// AddFortuneCompany insert a new FortuneCompany into database and returns
// last inserted Id on success.
func AddFortuneCompany(m *FortuneCompany) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCompanyById retrieves FortuneCompany by Id. Returns error if
// Id doesn't exist
func GetFortuneCompanyById(id int) (v *FortuneCompany, err error) {
	o := orm.NewOrm()
	v = &FortuneCompany{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCompany retrieves all FortuneCompany matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCompany(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCompany))
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

	var l []FortuneCompany
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

// UpdateFortuneCompany updates FortuneCompany by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCompanyById(m *FortuneCompany) (err error) {
	o := orm.NewOrm()
	v := FortuneCompany{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCompany deletes FortuneCompany by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCompany(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCompany{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCompany{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
