package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FORTUNEREPAYMENTPLAN struct {
	Id                     int       `orm:"column(ID);pk" description:"ID"`
	SUBSCRIPTIONID         string    `orm:"column(SUBSCRIPTION_ID);size(36)" description:"认购ID"`
	USERID                 string    `orm:"column(USER_ID);size(36);null" description:"Customer ID"`
	ONLINEUSERID           string    `orm:"column(ONLINE_USER_ID);size(36)" description:"线上用户ID"`
	LOANID                 string    `orm:"column(LOAN_ID);size(40)" description:"标的ID"`
	PERIODID               uint64    `orm:"column(PERIOD_ID)" description:"批次号ID"`
	CURRENTPERIOD          int       `orm:"column(CURRENT_PERIOD)" description:"当前期数"`
	PREVIOUSREPAYMENTDATE  time.Time `orm:"column(PREVIOUS_REPAYMENT_DATE);type(datetime);null" description:"上一个还款日"`
	CURRENTREPAYMENTDATE   time.Time `orm:"column(CURRENT_REPAYMENT_DATE);type(datetime);null" description:"本期还款日"`
	REALAMOUNTPRINCIPAL    float64   `orm:"column(REAL_AMOUNTPRINCIPAL);null;digits(15);decimals(2)" description:"实际本金"`
	REALAMOUNTINTEREST     float64   `orm:"column(REAL_AMOUNTINTEREST);null;digits(15);decimals(2)" description:"实际利息"`
	PREDICTAMOUNTPRINCIPAL float64   `orm:"column(PREDICT_AMOUNTPRINCIPAL);digits(15);decimals(2)" description:"预计本金"`
	PREDICTAMOUNTINTEREST  float64   `orm:"column(PREDICT_AMOUNTINTEREST);digits(15);decimals(2)" description:"预计利息"`
	REALREPAYMENTDATE      time.Time `orm:"column(REAL_REPAYMENT_DATE);type(datetime);null" description:"实际还款日期"`
	ORDERNO                uint64    `orm:"column(ORDER_NO)" description:"订单编号"`
	OPENACCOUNTNAME        string    `orm:"column(OPEN_ACCOUNT_NAME);size(100)" description:"开户名"`
	CREATETIME             time.Time `orm:"column(CREATE_TIME);type(datetime);null;auto_now_add" description:"创建时间"`
	UPDATETIME             time.Time `orm:"column(UPDATE_TIME);type(datetime);null;auto_now_add" description:"更新时间"`
	TOTALPERIOD            int       `orm:"column(TOTAL_PERIOD)" description:"总期数"`
}

func (t *FORTUNEREPAYMENTPLAN) TableName() string {
	return "FORTUNE_REPAYMENT_PLAN"
}

func init() {
	orm.RegisterModel(new(FORTUNEREPAYMENTPLAN))
}

// AddFORTUNEREPAYMENTPLAN insert a new FORTUNEREPAYMENTPLAN into database and returns
// last inserted Id on success.
func AddFORTUNEREPAYMENTPLAN(m *FORTUNEREPAYMENTPLAN) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFORTUNEREPAYMENTPLANById retrieves FORTUNEREPAYMENTPLAN by Id. Returns error if
// Id doesn't exist
func GetFORTUNEREPAYMENTPLANById(id int) (v *FORTUNEREPAYMENTPLAN, err error) {
	o := orm.NewOrm()
	v = &FORTUNEREPAYMENTPLAN{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFORTUNEREPAYMENTPLAN retrieves all FORTUNEREPAYMENTPLAN matches certain condition. Returns empty list if
// no records exist
func GetAllFORTUNEREPAYMENTPLAN(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FORTUNEREPAYMENTPLAN))
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

	var l []FORTUNEREPAYMENTPLAN
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

// UpdateFORTUNEREPAYMENTPLAN updates FORTUNEREPAYMENTPLAN by Id and returns error if
// the record to be updated doesn't exist
func UpdateFORTUNEREPAYMENTPLANById(m *FORTUNEREPAYMENTPLAN) (err error) {
	o := orm.NewOrm()
	v := FORTUNEREPAYMENTPLAN{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFORTUNEREPAYMENTPLAN deletes FORTUNEREPAYMENTPLAN by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFORTUNEREPAYMENTPLAN(id int) (err error) {
	o := orm.NewOrm()
	v := FORTUNEREPAYMENTPLAN{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FORTUNEREPAYMENTPLAN{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
