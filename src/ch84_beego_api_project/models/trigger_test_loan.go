package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TriggerTestLoan struct {
	Id   int  `orm:"column(id);pk"`
	Type int8 `orm:"column(type)" description:"项目类型:0人民币,1美元,2私募股权,3海外置业,4家庭保障、财富传承,5私募证券,6类固收"`
}

func (t *TriggerTestLoan) TableName() string {
	return "trigger_test_loan"
}

func init() {
	orm.RegisterModel(new(TriggerTestLoan))
}

// AddTriggerTestLoan insert a new TriggerTestLoan into database and returns
// last inserted Id on success.
func AddTriggerTestLoan(m *TriggerTestLoan) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTriggerTestLoanById retrieves TriggerTestLoan by Id. Returns error if
// Id doesn't exist
func GetTriggerTestLoanById(id int) (v *TriggerTestLoan, err error) {
	o := orm.NewOrm()
	v = &TriggerTestLoan{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTriggerTestLoan retrieves all TriggerTestLoan matches certain condition. Returns empty list if
// no records exist
func GetAllTriggerTestLoan(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TriggerTestLoan))
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

	var l []TriggerTestLoan
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

// UpdateTriggerTestLoan updates TriggerTestLoan by Id and returns error if
// the record to be updated doesn't exist
func UpdateTriggerTestLoanById(m *TriggerTestLoan) (err error) {
	o := orm.NewOrm()
	v := TriggerTestLoan{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTriggerTestLoan deletes TriggerTestLoan by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTriggerTestLoan(id int) (err error) {
	o := orm.NewOrm()
	v := TriggerTestLoan{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TriggerTestLoan{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
