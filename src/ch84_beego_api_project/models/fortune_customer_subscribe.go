package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerSubscribe struct {
	Id                    int       `orm:"column(id);auto"`
	UserId                string    `orm:"column(user_id);size(36)" description:"用户ID"`
	CcAllocateRefId       string    `orm:"column(cc_allocate_ref_id);size(36)" description:"电销客服分配关联id"`
	UserFullName          string    `orm:"column(user_full_name);size(20)" description:"用户姓名"`
	UserMobile            string    `orm:"column(user_mobile);size(20)" description:"用户手机号"`
	UserRegisterDatetime  time.Time `orm:"column(user_register_datetime);type(datetime);null" description:"用户注册时间"`
	Location              string    `orm:"column(location);size(100)" description:"所在地"`
	Status                int8      `orm:"column(status)" description:"当前进度 1=未分配、2=已分配、3=分配失败"`
	AppointmentChannel    int8      `orm:"column(appointment_channel)" description:"预约渠道 1推荐理财师 2已绑定用户预约 3理财师申请"`
	AppointmentStatus     int8      `orm:"column(appointment_status)" description:"沟通状态 0 待沟通 1 已沟通 2已失效"`
	AppointmentDeadline   time.Time `orm:"column(appointment_deadline);type(datetime);null" description:"预约失效时间 下个工作日"`
	AppointmentReason     int8      `orm:"column(appointment_reason)" description:"用户预约原因"`
	AppointmentReasonDesc string    `orm:"column(appointment_reason_desc);size(255)" description:"用户预约原因详情"`
	AppointmentFlag       int8      `orm:"column(appointment_flag)" description:"是否未绑定预约 0 否 1是"`
	BindSamId             string    `orm:"column(bind_sam_id);size(36)" description:"分配的客户经理ID"`
	BindSamName           string    `orm:"column(bind_sam_name);size(50)" description:"分配的客户经理姓名"`
	AllotDate             time.Time `orm:"column(allot_date);type(datetime)" description:"分配时间"`
	Remark                string    `orm:"column(remark);size(200)" description:"备注"`
	ExtendField           string    `orm:"column(extend_field);size(2000)" description:"扩展数据存放"`
	Deleted               int8      `orm:"column(deleted)" description:"是否删除"`
	DateEntered           time.Time `orm:"column(date_entered);type(datetime);null"`
	DateModified          time.Time `orm:"column(date_modified);type(datetime);null"`
}

func (t *FortuneCustomerSubscribe) TableName() string {
	return "fortune_customer_subscribe"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerSubscribe))
}

// AddFortuneCustomerSubscribe insert a new FortuneCustomerSubscribe into database and returns
// last inserted Id on success.
func AddFortuneCustomerSubscribe(m *FortuneCustomerSubscribe) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerSubscribeById retrieves FortuneCustomerSubscribe by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerSubscribeById(id int) (v *FortuneCustomerSubscribe, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerSubscribe{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerSubscribe retrieves all FortuneCustomerSubscribe matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerSubscribe(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerSubscribe))
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

	var l []FortuneCustomerSubscribe
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

// UpdateFortuneCustomerSubscribe updates FortuneCustomerSubscribe by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerSubscribeById(m *FortuneCustomerSubscribe) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerSubscribe{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerSubscribe deletes FortuneCustomerSubscribe by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerSubscribe(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerSubscribe{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerSubscribe{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
