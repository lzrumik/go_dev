package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomizeSubscribe struct {
	Id              int       `orm:"column(id);pk"`
	Name            string    `orm:"column(name);size(200)" description:"客户姓名"`
	IpAddress       string    `orm:"column(ip_address);size(100)" description:"IP地址"`
	Location        string    `orm:"column(location);size(100)" description:"预约时所在地"`
	UserId          string    `orm:"column(user_id);size(36)" description:"用户ID"`
	CcAllocateRefId string    `orm:"column(cc_allocate_ref_id);size(36)" description:"电销客服分配关联id"`
	MobilePrefix    string    `orm:"column(mobile_prefix);size(10)" description:"手机号前缀"`
	PhoneMobile     string    `orm:"column(phone_mobile);size(50)" description:"手机号"`
	MobileLocation  string    `orm:"column(mobile_location);size(100)" description:"手机号所在地"`
	IsBindSam       uint8     `orm:"column(is_bind_sam)" description:"是否绑定客户经理"`
	BindSamId       string    `orm:"column(bind_sam_id);size(36)" description:"客户经理ID"`
	BindSamName     string    `orm:"column(bind_sam_name);size(50)" description:"分配的客户经理姓名"`
	AllotStatus     int8      `orm:"column(allot_status)" description:"状态 0=未分配、1=已分配、2=分配失败"`
	AllotDate       time.Time `orm:"column(allot_date);type(datetime)" description:"分配时间"`
	Remark          string    `orm:"column(remark);size(200)" description:"备注"`
	DateEntered     time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified    time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	OpUser          string    `orm:"column(op_user);size(36)" description:"修改人"`
	Deleted         uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneCustomizeSubscribe) TableName() string {
	return "fortune_customize_subscribe"
}

func init() {
	orm.RegisterModel(new(FortuneCustomizeSubscribe))
}

// AddFortuneCustomizeSubscribe insert a new FortuneCustomizeSubscribe into database and returns
// last inserted Id on success.
func AddFortuneCustomizeSubscribe(m *FortuneCustomizeSubscribe) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomizeSubscribeById retrieves FortuneCustomizeSubscribe by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomizeSubscribeById(id int) (v *FortuneCustomizeSubscribe, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomizeSubscribe{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomizeSubscribe retrieves all FortuneCustomizeSubscribe matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomizeSubscribe(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomizeSubscribe))
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

	var l []FortuneCustomizeSubscribe
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

// UpdateFortuneCustomizeSubscribe updates FortuneCustomizeSubscribe by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomizeSubscribeById(m *FortuneCustomizeSubscribe) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomizeSubscribe{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomizeSubscribe deletes FortuneCustomizeSubscribe by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomizeSubscribe(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomizeSubscribe{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomizeSubscribe{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
