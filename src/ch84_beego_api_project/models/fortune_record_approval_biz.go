package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneRecordApprovalBiz struct {
	Id           int       `orm:"column(id);auto"`
	ProductType  uint8     `orm:"column(product_type)" description:"产品线(产品类型)"`
	SubmitType   uint8     `orm:"column(submit_type)" description:"认购方式"`
	BelongCity   int8      `orm:"column(belong_city)" description:"所属城市"`
	CreateUserId string    `orm:"column(create_user_id);size(36)" description:"创建人ID"`
	DateEntered  time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	Deleted      uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneRecordApprovalBiz) TableName() string {
	return "fortune_record_approval_biz"
}

func init() {
	orm.RegisterModel(new(FortuneRecordApprovalBiz))
}

// AddFortuneRecordApprovalBiz insert a new FortuneRecordApprovalBiz into database and returns
// last inserted Id on success.
func AddFortuneRecordApprovalBiz(m *FortuneRecordApprovalBiz) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneRecordApprovalBizById retrieves FortuneRecordApprovalBiz by Id. Returns error if
// Id doesn't exist
func GetFortuneRecordApprovalBizById(id int) (v *FortuneRecordApprovalBiz, err error) {
	o := orm.NewOrm()
	v = &FortuneRecordApprovalBiz{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneRecordApprovalBiz retrieves all FortuneRecordApprovalBiz matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneRecordApprovalBiz(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneRecordApprovalBiz))
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

	var l []FortuneRecordApprovalBiz
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

// UpdateFortuneRecordApprovalBiz updates FortuneRecordApprovalBiz by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneRecordApprovalBizById(m *FortuneRecordApprovalBiz) (err error) {
	o := orm.NewOrm()
	v := FortuneRecordApprovalBiz{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneRecordApprovalBiz deletes FortuneRecordApprovalBiz by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneRecordApprovalBiz(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneRecordApprovalBiz{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneRecordApprovalBiz{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
