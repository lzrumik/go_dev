package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerWealthPlanne struct {
	Id                     int       `orm:"column(id);auto"`
	UserId                 string    `orm:"column(user_id);size(36)" description:"用户ID"`
	Username               string    `orm:"column(username);size(50)" description:"用户名"`
	UserLevel              int8      `orm:"column(user_level)" description:"用户等级 1=高端用户 2=普通用户 3=未知"`
	UserLevelText          string    `orm:"column(user_level_text);size(50)" description:"用户等级 1=高端用户 2=普通用户 3=未知"`
	FirstCommitTime        time.Time `orm:"column(first_commit_time);type(datetime);null" description:"首次提交时间"`
	FirstCommitBindSamId   string    `orm:"column(first_commit_bind_sam_id);size(36)" description:"首次提交绑定客户经理ID"`
	FirstCommitBindSamName string    `orm:"column(first_commit_bind_sam_name);size(50)" description:"首次提交绑定客户经理姓名"`
	LastUpdateTime         time.Time `orm:"column(last_update_time);type(datetime);null" description:"最近一次更新时间"`
	ReportLastTime         time.Time `orm:"column(report_last_time);type(datetime);null" description:"最新一次生成报告时间"`
	BindSamId              string    `orm:"column(bind_sam_id);size(36)" description:"分配的客户经理ID"`
	BindSamName            string    `orm:"column(bind_sam_name);size(50)" description:"分配的客户经理姓名"`
	AllotDate              time.Time `orm:"column(allot_date);type(datetime);null" description:"分配时间"`
	AllotStatus            int8      `orm:"column(allot_status)" description:"分配状态 0=未分配（历史数据）、1=已分配、2=分配失败"`
	DateEntered            time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified           time.Time `orm:"column(date_modified);type(datetime)"`
	Deleted                uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneCustomerWealthPlanne) TableName() string {
	return "fortune_customer_wealth_planne"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerWealthPlanne))
}

// AddFortuneCustomerWealthPlanne insert a new FortuneCustomerWealthPlanne into database and returns
// last inserted Id on success.
func AddFortuneCustomerWealthPlanne(m *FortuneCustomerWealthPlanne) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerWealthPlanneById retrieves FortuneCustomerWealthPlanne by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerWealthPlanneById(id int) (v *FortuneCustomerWealthPlanne, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerWealthPlanne{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerWealthPlanne retrieves all FortuneCustomerWealthPlanne matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerWealthPlanne(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerWealthPlanne))
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

	var l []FortuneCustomerWealthPlanne
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

// UpdateFortuneCustomerWealthPlanne updates FortuneCustomerWealthPlanne by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerWealthPlanneById(m *FortuneCustomerWealthPlanne) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerWealthPlanne{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerWealthPlanne deletes FortuneCustomerWealthPlanne by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerWealthPlanne(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerWealthPlanne{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerWealthPlanne{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
