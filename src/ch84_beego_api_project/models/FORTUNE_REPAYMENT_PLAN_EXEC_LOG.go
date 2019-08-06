package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FORTUNEREPAYMENTPLANEXECLOG struct {
	Id             int       `orm:"column(ID);auto" description:"ID"`
	LOANID         string    `orm:"column(LOAN_ID);size(40)" description:"标的ID"`
	PERIODID       uint64    `orm:"column(PERIOD_ID)" description:"批次号ID"`
	TOTALPERIOD    uint      `orm:"column(TOTAL_PERIOD)" description:"总期数"`
	CURRENTPERIOD  int       `orm:"column(CURRENT_PERIOD)" description:"当前期数"`
	SUBTYPE        int       `orm:"column(SUB_TYPE)" description:"类型  1:批量导入回息 2:修改回息"`
	LOGS           string    `orm:"column(LOGS)" description:"日志信息"`
	CREATEUSERNAME string    `orm:"column(CREATE_USER_NAME);size(30)" description:"创建人"`
	CREATEUSERID   string    `orm:"column(CREATE_USER_ID);size(36)" description:"创建人"`
	DATEENTERED    time.Time `orm:"column(DATE_ENTERED);type(datetime)" description:"添加时间"`
}

func (t *FORTUNEREPAYMENTPLANEXECLOG) TableName() string {
	return "FORTUNE_REPAYMENT_PLAN_EXEC_LOG"
}

func init() {
	orm.RegisterModel(new(FORTUNEREPAYMENTPLANEXECLOG))
}

// AddFORTUNEREPAYMENTPLANEXECLOG insert a new FORTUNEREPAYMENTPLANEXECLOG into database and returns
// last inserted Id on success.
func AddFORTUNEREPAYMENTPLANEXECLOG(m *FORTUNEREPAYMENTPLANEXECLOG) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFORTUNEREPAYMENTPLANEXECLOGById retrieves FORTUNEREPAYMENTPLANEXECLOG by Id. Returns error if
// Id doesn't exist
func GetFORTUNEREPAYMENTPLANEXECLOGById(id int) (v *FORTUNEREPAYMENTPLANEXECLOG, err error) {
	o := orm.NewOrm()
	v = &FORTUNEREPAYMENTPLANEXECLOG{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFORTUNEREPAYMENTPLANEXECLOG retrieves all FORTUNEREPAYMENTPLANEXECLOG matches certain condition. Returns empty list if
// no records exist
func GetAllFORTUNEREPAYMENTPLANEXECLOG(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FORTUNEREPAYMENTPLANEXECLOG))
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

	var l []FORTUNEREPAYMENTPLANEXECLOG
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

// UpdateFORTUNEREPAYMENTPLANEXECLOG updates FORTUNEREPAYMENTPLANEXECLOG by Id and returns error if
// the record to be updated doesn't exist
func UpdateFORTUNEREPAYMENTPLANEXECLOGById(m *FORTUNEREPAYMENTPLANEXECLOG) (err error) {
	o := orm.NewOrm()
	v := FORTUNEREPAYMENTPLANEXECLOG{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFORTUNEREPAYMENTPLANEXECLOG deletes FORTUNEREPAYMENTPLANEXECLOG by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFORTUNEREPAYMENTPLANEXECLOG(id int) (err error) {
	o := orm.NewOrm()
	v := FORTUNEREPAYMENTPLANEXECLOG{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FORTUNEREPAYMENTPLANEXECLOG{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
