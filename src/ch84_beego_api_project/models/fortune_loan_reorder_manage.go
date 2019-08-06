package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneLoanReorderManage struct {
	Id            int       `orm:"column(id);auto"`
	Name          string    `orm:"column(name);size(50)" description:"项目名称"`
	LoanId        string    `orm:"column(loan_id);size(36)" description:"项目ID"`
	RelevancePage uint8     `orm:"column(relevance_page)" description:"针对页面 0:首页"`
	Sort          int8      `orm:"column(sort)" description:"排序"`
	Type          string    `orm:"column(type);size(5)" description:"类型 FRE:人民币固定收益 FCE:类固定收益 LE:地产基金 PE:私募股权 PES:私募证券"`
	DateEntered   time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified  time.Time `orm:"column(date_modified);type(datetime)"`
	CreatedUser   string    `orm:"column(created_user);size(100)" description:"创建人"`
	Deleted       uint8     `orm:"column(deleted)"`
}

func (t *FortuneLoanReorderManage) TableName() string {
	return "fortune_loan_reorder_manage"
}

func init() {
	orm.RegisterModel(new(FortuneLoanReorderManage))
}

// AddFortuneLoanReorderManage insert a new FortuneLoanReorderManage into database and returns
// last inserted Id on success.
func AddFortuneLoanReorderManage(m *FortuneLoanReorderManage) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneLoanReorderManageById retrieves FortuneLoanReorderManage by Id. Returns error if
// Id doesn't exist
func GetFortuneLoanReorderManageById(id int) (v *FortuneLoanReorderManage, err error) {
	o := orm.NewOrm()
	v = &FortuneLoanReorderManage{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneLoanReorderManage retrieves all FortuneLoanReorderManage matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneLoanReorderManage(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneLoanReorderManage))
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

	var l []FortuneLoanReorderManage
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

// UpdateFortuneLoanReorderManage updates FortuneLoanReorderManage by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneLoanReorderManageById(m *FortuneLoanReorderManage) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanReorderManage{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneLoanReorderManage deletes FortuneLoanReorderManage by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneLoanReorderManage(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanReorderManage{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneLoanReorderManage{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
