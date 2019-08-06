package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneActivitySignup struct {
	Id              int       `orm:"column(id);auto" description:"主键"`
	ActivityNo      string    `orm:"column(activity_no);size(36)" description:"活动编号"`
	UserId          string    `orm:"column(user_id);size(36)" description:"用户ID"`
	Name            string    `orm:"column(name);size(36)" description:"用户姓名"`
	Mobile          string    `orm:"column(mobile);size(36)" description:"用户手机号"`
	AppointmentUser int8      `orm:"column(appointment_user);null" description:"是否为预约用户：0 否；1 是；"`
	Status          int8      `orm:"column(status)" description:"状态：1 待确认；2 已通过；3 已拒绝；4 已取消；"`
	Type            int8      `orm:"column(type)" description:"是否为符合条件用户：0 否；1 是；"`
	OwnCity         string    `orm:"column(own_city);size(20);null" description:"所属城市"`
	OwnCityId       int       `orm:"column(own_city_id);null" description:"所属城市id"`
	OwnGroup        string    `orm:"column(own_group);size(20);null" description:"所属团队"`
	OwnGroupId      int       `orm:"column(own_group_id);null" description:"所属团队id"`
	Creator         int       `orm:"column(creator);null" description:"创建人"`
	CreateTime      time.Time `orm:"column(create_time);type(datetime);auto_now_add" description:"创建时间"`
	CancleTime      time.Time `orm:"column(cancle_time);type(datetime);null" description:"取消报名时间"`
	Approver        int       `orm:"column(approver);null" description:"审批人"`
	ApproveTime     time.Time `orm:"column(approve_time);type(datetime);null" description:"审批时间"`
	Reason          int8      `orm:"column(reason);null" description:"拒绝原因：1 不符合条件；2 人数已满；3 活动暂停；4 其他；"`
	Remark          string    `orm:"column(remark);size(200);null" description:"说明：拒绝时必填"`
	LastUpdateTime  time.Time `orm:"column(last_update_time);type(datetime);auto_now_add" description:"最后更新时间"`
}

func (t *FortuneActivitySignup) TableName() string {
	return "fortune_activity_signup"
}

func init() {
	orm.RegisterModel(new(FortuneActivitySignup))
}

// AddFortuneActivitySignup insert a new FortuneActivitySignup into database and returns
// last inserted Id on success.
func AddFortuneActivitySignup(m *FortuneActivitySignup) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneActivitySignupById retrieves FortuneActivitySignup by Id. Returns error if
// Id doesn't exist
func GetFortuneActivitySignupById(id int) (v *FortuneActivitySignup, err error) {
	o := orm.NewOrm()
	v = &FortuneActivitySignup{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneActivitySignup retrieves all FortuneActivitySignup matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneActivitySignup(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneActivitySignup))
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

	var l []FortuneActivitySignup
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

// UpdateFortuneActivitySignup updates FortuneActivitySignup by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneActivitySignupById(m *FortuneActivitySignup) (err error) {
	o := orm.NewOrm()
	v := FortuneActivitySignup{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneActivitySignup deletes FortuneActivitySignup by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneActivitySignup(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneActivitySignup{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneActivitySignup{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
