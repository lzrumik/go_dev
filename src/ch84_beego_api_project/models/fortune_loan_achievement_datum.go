package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneLoanAchievementDatum struct {
	Id              int       `orm:"column(id);auto" description:"自增主健"`
	LoanId          string    `orm:"column(loan_id);size(36)" description:"项目ID"`
	MinInvestAmount float64   `orm:"column(min_invest_amount);digits(12);decimals(4)" description:"最小起投金额"`
	MaxInvestAmount float64   `orm:"column(max_invest_amount);digits(12);decimals(4)" description:"最大起投金额"`
	InterestRate    float64   `orm:"column(interest_rate);digits(12);decimals(2)" description:"利率(业绩基准)"`
	ServiceRate     float64   `orm:"column(service_rate);digits(12);decimals(2)" description:"服务费率"`
	CreatedUserId   string    `orm:"column(created_user_id);size(36)" description:"创建人ID"`
	ModifiedUserId  string    `orm:"column(modified_user_id);size(36)" description:"修改人ID"`
	DateEntered     time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified    time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	Deleted         uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneLoanAchievementDatum) TableName() string {
	return "fortune_loan_achievement_datum"
}

func init() {
	orm.RegisterModel(new(FortuneLoanAchievementDatum))
}

// AddFortuneLoanAchievementDatum insert a new FortuneLoanAchievementDatum into database and returns
// last inserted Id on success.
func AddFortuneLoanAchievementDatum(m *FortuneLoanAchievementDatum) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneLoanAchievementDatumById retrieves FortuneLoanAchievementDatum by Id. Returns error if
// Id doesn't exist
func GetFortuneLoanAchievementDatumById(id int) (v *FortuneLoanAchievementDatum, err error) {
	o := orm.NewOrm()
	v = &FortuneLoanAchievementDatum{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneLoanAchievementDatum retrieves all FortuneLoanAchievementDatum matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneLoanAchievementDatum(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneLoanAchievementDatum))
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

	var l []FortuneLoanAchievementDatum
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

// UpdateFortuneLoanAchievementDatum updates FortuneLoanAchievementDatum by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneLoanAchievementDatumById(m *FortuneLoanAchievementDatum) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanAchievementDatum{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneLoanAchievementDatum deletes FortuneLoanAchievementDatum by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneLoanAchievementDatum(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanAchievementDatum{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneLoanAchievementDatum{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
