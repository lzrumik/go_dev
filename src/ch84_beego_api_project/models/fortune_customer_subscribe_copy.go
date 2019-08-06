package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerSubscribeCopy struct {
	Id                   int       `orm:"column(id);auto"`
	UserId               string    `orm:"column(user_id);size(36)" description:"用户ID"`
	CcAllocateRefId      string    `orm:"column(cc_allocate_ref_id);size(36)" description:"电销客服分配关联id"`
	UserFullName         string    `orm:"column(user_full_name);size(20)" description:"用户姓名"`
	UserMobile           string    `orm:"column(user_mobile);size(20)" description:"用户手机号"`
	UserRegisterDatetime time.Time `orm:"column(user_register_datetime);type(datetime);null" description:"用户注册时间"`
	Location             string    `orm:"column(location);size(100)" description:"所在地"`
	Status               int8      `orm:"column(status)" description:"状态 1=未分配、2=已分配、3=分配失败"`
	BindSamId            string    `orm:"column(bind_sam_id);size(36)" description:"分配的客户经理ID"`
	BindSamName          string    `orm:"column(bind_sam_name);size(50)" description:"分配的客户经理姓名"`
	AllotDate            time.Time `orm:"column(allot_date);type(datetime)" description:"分配时间"`
	Remark               string    `orm:"column(remark);size(200)" description:"备注"`
	ExtendField          string    `orm:"column(extend_field);size(2000)" description:"扩展数据存放"`
	Deleted              int8      `orm:"column(deleted)" description:"是否删除"`
	DateEntered          time.Time `orm:"column(date_entered);type(datetime);null"`
	DateModified         time.Time `orm:"column(date_modified);type(datetime);null"`
}

func (t *FortuneCustomerSubscribeCopy) TableName() string {
	return "fortune_customer_subscribe_copy"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerSubscribeCopy))
}

// AddFortuneCustomerSubscribeCopy insert a new FortuneCustomerSubscribeCopy into database and returns
// last inserted Id on success.
func AddFortuneCustomerSubscribeCopy(m *FortuneCustomerSubscribeCopy) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerSubscribeCopyById retrieves FortuneCustomerSubscribeCopy by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerSubscribeCopyById(id int) (v *FortuneCustomerSubscribeCopy, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerSubscribeCopy{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerSubscribeCopy retrieves all FortuneCustomerSubscribeCopy matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerSubscribeCopy(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerSubscribeCopy))
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

	var l []FortuneCustomerSubscribeCopy
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

// UpdateFortuneCustomerSubscribeCopy updates FortuneCustomerSubscribeCopy by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerSubscribeCopyById(m *FortuneCustomerSubscribeCopy) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerSubscribeCopy{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerSubscribeCopy deletes FortuneCustomerSubscribeCopy by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerSubscribeCopy(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerSubscribeCopy{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerSubscribeCopy{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
