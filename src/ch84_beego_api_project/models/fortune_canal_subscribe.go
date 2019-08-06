package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCanalSubscribe struct {
	Id                   int       `orm:"column(id);auto"`
	UserId               string    `orm:"column(user_id);size(36)" description:"用户ID"`
	CcAllocateRefId      string    `orm:"column(cc_allocate_ref_id);size(36)" description:"电销客服分配关联id"`
	Canal                string    `orm:"column(canal);size(100)" description:"渠道号"`
	UserFullName         string    `orm:"column(user_full_name);size(30)" description:"用户姓名"`
	UserMobile           string    `orm:"column(user_mobile);size(20)" description:"用户注册手机号"`
	UserMobileLocation   string    `orm:"column(user_mobile_location);size(200)" description:"用户注册手机号归属地"`
	SubscribeLocation    string    `orm:"column(subscribe_location);size(200)" description:"提交预约时所在地"`
	UserRegisterDatetime time.Time `orm:"column(user_register_datetime);type(datetime)" description:"用户注册时间"`
	CommitIsBindSam      string    `orm:"column(commit_is_bind_sam);size(50)" description:"提交预约时是否有对接客户经理"`
	CommitBinSamName     string    `orm:"column(commit_bin_sam_name);size(50)" description:"提交预约时对接的客户经理姓名"`
	BinSamId             string    `orm:"column(bin_sam_id);size(50)" description:"分配的客户经理id"`
	BinSamName           string    `orm:"column(bin_sam_name);size(50)" description:"分配的客户经理姓名"`
	Status               int8      `orm:"column(status)" description:"状态 0=未分配、1=已分配、2=分配失败"`
	AllotDate            time.Time `orm:"column(allot_date);type(datetime)" description:"分配时间"`
	Remark               string    `orm:"column(remark);size(200)" description:"备注"`
	Deleted              int8      `orm:"column(deleted)" description:"是否删除"`
	DateEntered          time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified         time.Time `orm:"column(date_modified);type(datetime)"`
}

func (t *FortuneCanalSubscribe) TableName() string {
	return "fortune_canal_subscribe"
}

func init() {
	orm.RegisterModel(new(FortuneCanalSubscribe))
}

// AddFortuneCanalSubscribe insert a new FortuneCanalSubscribe into database and returns
// last inserted Id on success.
func AddFortuneCanalSubscribe(m *FortuneCanalSubscribe) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCanalSubscribeById retrieves FortuneCanalSubscribe by Id. Returns error if
// Id doesn't exist
func GetFortuneCanalSubscribeById(id int) (v *FortuneCanalSubscribe, err error) {
	o := orm.NewOrm()
	v = &FortuneCanalSubscribe{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCanalSubscribe retrieves all FortuneCanalSubscribe matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCanalSubscribe(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCanalSubscribe))
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

	var l []FortuneCanalSubscribe
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

// UpdateFortuneCanalSubscribe updates FortuneCanalSubscribe by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCanalSubscribeById(m *FortuneCanalSubscribe) (err error) {
	o := orm.NewOrm()
	v := FortuneCanalSubscribe{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCanalSubscribe deletes FortuneCanalSubscribe by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCanalSubscribe(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCanalSubscribe{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCanalSubscribe{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
