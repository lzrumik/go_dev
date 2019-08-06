package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type FortuneLoanCstm struct {
	Id              int    `orm:"column(id);pk"`
	PublisherDesc   string `orm:"column(publisher_desc)" description:"发行人介绍"`
	UseDesc         string `orm:"column(use_desc)" description:"募集资金用途"`
	BrightenedDot   string `orm:"column(brightened_dot)" description:"产品亮点"`
	FundManagerDesc string `orm:"column(fund_manager_desc)" description:"基金管理人介绍"`
	FitInvestor     string `orm:"column(fit_investor)" description:"适合投资人"`
	BlockoutPeriod  string `orm:"column(blockout_period)" description:"封闭期"`
	FundCost        string `orm:"column(fund_cost)" description:"基金费用"`
}

func (t *FortuneLoanCstm) TableName() string {
	return "fortune_loan_cstm"
}

func init() {
	orm.RegisterModel(new(FortuneLoanCstm))
}

// AddFortuneLoanCstm insert a new FortuneLoanCstm into database and returns
// last inserted Id on success.
func AddFortuneLoanCstm(m *FortuneLoanCstm) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneLoanCstmById retrieves FortuneLoanCstm by Id. Returns error if
// Id doesn't exist
func GetFortuneLoanCstmById(id int) (v *FortuneLoanCstm, err error) {
	o := orm.NewOrm()
	v = &FortuneLoanCstm{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneLoanCstm retrieves all FortuneLoanCstm matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneLoanCstm(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneLoanCstm))
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

	var l []FortuneLoanCstm
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

// UpdateFortuneLoanCstm updates FortuneLoanCstm by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneLoanCstmById(m *FortuneLoanCstm) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanCstm{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneLoanCstm deletes FortuneLoanCstm by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneLoanCstm(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanCstm{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneLoanCstm{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}