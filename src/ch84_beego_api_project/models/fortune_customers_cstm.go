package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomersCstm struct {
	Id               int       `orm:"column(customer_id);pk" description:"客户id"`
	MarkName         string    `orm:"column(mark_name);size(100)" description:"客户备注姓名"`
	Birthday         string    `orm:"column(birthday);size(20)" description:"生日"`
	Sex              string    `orm:"column(sex);size(10)" description:"性别 male=男 female=女"`
	NativePlace      string    `orm:"column(native_place);size(100)" description:"籍贯"`
	ContactAddress   string    `orm:"column(contact_address);size(200)" description:"联系地址"`
	WorkingCondition string    `orm:"column(working_condition);size(200)" description:"工作情况"`
	Hobbies          string    `orm:"column(hobbies);size(200)" description:"兴趣爱好"`
	PhoneMobileArea  string    `orm:"column(phone_mobile_area);size(50)" description:"手机归属地"`
	OpenDateBirth    string    `orm:"column(open_date_birth);size(20)" description:"开户出生日期"`
	OpenSex          string    `orm:"column(open_sex);size(10)" description:"开户性别 male=男 female=女"`
	OpenLocation     string    `orm:"column(open_location);size(200)" description:"开户所在地"`
	DateEntered      time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified     time.Time `orm:"column(date_modified);type(datetime)"`
}

func (t *FortuneCustomersCstm) TableName() string {
	return "fortune_customers_cstm"
}

func init() {
	orm.RegisterModel(new(FortuneCustomersCstm))
}

// AddFortuneCustomersCstm insert a new FortuneCustomersCstm into database and returns
// last inserted Id on success.
func AddFortuneCustomersCstm(m *FortuneCustomersCstm) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomersCstmById retrieves FortuneCustomersCstm by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomersCstmById(id int) (v *FortuneCustomersCstm, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomersCstm{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomersCstm retrieves all FortuneCustomersCstm matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomersCstm(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomersCstm))
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

	var l []FortuneCustomersCstm
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

// UpdateFortuneCustomersCstm updates FortuneCustomersCstm by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomersCstmById(m *FortuneCustomersCstm) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomersCstm{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomersCstm deletes FortuneCustomersCstm by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomersCstm(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomersCstm{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomersCstm{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
