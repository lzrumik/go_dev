package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneUserSubscriptionApproveInfo struct {
	Id            int       `orm:"column(approve_id);pk" description:"主键ID"`
	SubId         string    `orm:"column(sub_id);size(36)" description:"认购ID"`
	ApprovalLevel uint8     `orm:"column(approval_level)" description:"审批级别"`
	ApproveAction int8      `orm:"column(approve_action)" description:"审批动作 1通过 2拒绝"`
	RefusalType   int       `orm:"column(refusal_type)" description:"拒绝类型"`
	RefusalCause  string    `orm:"column(refusal_cause);size(100)" description:"拒绝原因"`
	ApproveUserId string    `orm:"column(approve_user_id);size(36)"`
	ApproveTime   time.Time `orm:"column(approve_time);type(timestamp)" description:"创建时间"`
}

func (t *FortuneUserSubscriptionApproveInfo) TableName() string {
	return "fortune_user_subscription_approve_info"
}

func init() {
	orm.RegisterModel(new(FortuneUserSubscriptionApproveInfo))
}

// AddFortuneUserSubscriptionApproveInfo insert a new FortuneUserSubscriptionApproveInfo into database and returns
// last inserted Id on success.
func AddFortuneUserSubscriptionApproveInfo(m *FortuneUserSubscriptionApproveInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneUserSubscriptionApproveInfoById retrieves FortuneUserSubscriptionApproveInfo by Id. Returns error if
// Id doesn't exist
func GetFortuneUserSubscriptionApproveInfoById(id int) (v *FortuneUserSubscriptionApproveInfo, err error) {
	o := orm.NewOrm()
	v = &FortuneUserSubscriptionApproveInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneUserSubscriptionApproveInfo retrieves all FortuneUserSubscriptionApproveInfo matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneUserSubscriptionApproveInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneUserSubscriptionApproveInfo))
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

	var l []FortuneUserSubscriptionApproveInfo
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

// UpdateFortuneUserSubscriptionApproveInfo updates FortuneUserSubscriptionApproveInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneUserSubscriptionApproveInfoById(m *FortuneUserSubscriptionApproveInfo) (err error) {
	o := orm.NewOrm()
	v := FortuneUserSubscriptionApproveInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneUserSubscriptionApproveInfo deletes FortuneUserSubscriptionApproveInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneUserSubscriptionApproveInfo(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneUserSubscriptionApproveInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneUserSubscriptionApproveInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
