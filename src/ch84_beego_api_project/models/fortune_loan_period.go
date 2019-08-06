package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneLoanPeriod struct {
	Id             int       `orm:"column(id);auto" description:"自增主健"`
	LoanId         string    `orm:"column(loan_id);size(36)" description:"项目ID"`
	PeriodNumber   uint      `orm:"column(period_number)" description:"期数(当前第几期)"`
	ValueDate      time.Time `orm:"column(value_date);type(date)" description:"产品起息日"`
	ExpireDate     time.Time `orm:"column(expire_date);type(date)" description:"产品到期日"`
	CreatedUserId  string    `orm:"column(created_user_id);size(36)" description:"创建人ID"`
	ModifiedUserId string    `orm:"column(modified_user_id);size(36)" description:"修改人ID"`
	DateEntered    time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified   time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	Deleted        uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneLoanPeriod) TableName() string {
	return "fortune_loan_period"
}

func init() {
	orm.RegisterModel(new(FortuneLoanPeriod))
}

// AddFortuneLoanPeriod insert a new FortuneLoanPeriod into database and returns
// last inserted Id on success.
func AddFortuneLoanPeriod(m *FortuneLoanPeriod) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneLoanPeriodById retrieves FortuneLoanPeriod by Id. Returns error if
// Id doesn't exist
func GetFortuneLoanPeriodById(id int) (v *FortuneLoanPeriod, err error) {
	o := orm.NewOrm()
	v = &FortuneLoanPeriod{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneLoanPeriod retrieves all FortuneLoanPeriod matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneLoanPeriod(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneLoanPeriod))
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

	var l []FortuneLoanPeriod
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

// UpdateFortuneLoanPeriod updates FortuneLoanPeriod by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneLoanPeriodById(m *FortuneLoanPeriod) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanPeriod{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneLoanPeriod deletes FortuneLoanPeriod by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneLoanPeriod(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanPeriod{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneLoanPeriod{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
