package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneUsaAccountInfo struct {
	Id              int       `orm:"column(uid);pk"`
	FirstName       string    `orm:"column(first_name);size(50)" description:"姓"`
	LastName        string    `orm:"column(last_name);size(50)" description:"名"`
	Ssn             string    `orm:"column(ssn);size(9)" description:"美国社会安全码"`
	KycStatus       uint8     `orm:"column(kyc_status)" description:"KYC状态 0:未审核 1:审核中 2:审核失败 3:审核通过"`
	KycTime         time.Time `orm:"column(kyc_time);type(datetime);null" description:"KYC状态变更时间"`
	UsaVerifyStatus uint8     `orm:"column(usa_verify_status)" description:"美国合格投资人认证审核状态 0:未审核 1:审核中 2:审核失败 3:审核通过"`
	UsaVerifyTime   time.Time `orm:"column(usa_verify_time);type(datetime);null" description:"状态变更时间"`
	DateEntered     time.Time `orm:"column(date_entered);type(datetime);null"`
	DateModified    time.Time `orm:"column(date_modified);type(datetime);null"`
	Deleted         uint8     `orm:"column(deleted)"`
}

func (t *FortuneUsaAccountInfo) TableName() string {
	return "fortune_usa_account_info"
}

func init() {
	orm.RegisterModel(new(FortuneUsaAccountInfo))
}

// AddFortuneUsaAccountInfo insert a new FortuneUsaAccountInfo into database and returns
// last inserted Id on success.
func AddFortuneUsaAccountInfo(m *FortuneUsaAccountInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneUsaAccountInfoById retrieves FortuneUsaAccountInfo by Id. Returns error if
// Id doesn't exist
func GetFortuneUsaAccountInfoById(id int) (v *FortuneUsaAccountInfo, err error) {
	o := orm.NewOrm()
	v = &FortuneUsaAccountInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneUsaAccountInfo retrieves all FortuneUsaAccountInfo matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneUsaAccountInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneUsaAccountInfo))
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

	var l []FortuneUsaAccountInfo
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

// UpdateFortuneUsaAccountInfo updates FortuneUsaAccountInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneUsaAccountInfoById(m *FortuneUsaAccountInfo) (err error) {
	o := orm.NewOrm()
	v := FortuneUsaAccountInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneUsaAccountInfo deletes FortuneUsaAccountInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneUsaAccountInfo(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneUsaAccountInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneUsaAccountInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
