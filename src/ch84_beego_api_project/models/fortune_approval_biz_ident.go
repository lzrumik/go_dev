package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneApprovalBizIdent struct {
	Id                   int       `orm:"column(id);auto"`
	BizName              string    `orm:"column(biz_name);size(100)" description:"业务名称"`
	UniqueIdent          string    `orm:"column(unique_ident);size(50)" description:"唯一标识"`
	IsRelevanceCustomBiz uint8     `orm:"column(is_relevance_custom_biz)" description:"是关联自定义业务 0:否 1:是"`
	CustomBizTableName   string    `orm:"column(custom_biz_table_name);size(100)" description:"自定义配置表名"`
	CustomBizUrl         string    `orm:"column(custom_biz_url);size(200)" description:"自定义业务配置地址"`
	CreateUserId         string    `orm:"column(create_user_id);size(36)" description:"创建人ID"`
	DateEntered          time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified         time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	Deleted              uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneApprovalBizIdent) TableName() string {
	return "fortune_approval_biz_ident"
}

func init() {
	orm.RegisterModel(new(FortuneApprovalBizIdent))
}

// AddFortuneApprovalBizIdent insert a new FortuneApprovalBizIdent into database and returns
// last inserted Id on success.
func AddFortuneApprovalBizIdent(m *FortuneApprovalBizIdent) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneApprovalBizIdentById retrieves FortuneApprovalBizIdent by Id. Returns error if
// Id doesn't exist
func GetFortuneApprovalBizIdentById(id int) (v *FortuneApprovalBizIdent, err error) {
	o := orm.NewOrm()
	v = &FortuneApprovalBizIdent{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneApprovalBizIdent retrieves all FortuneApprovalBizIdent matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneApprovalBizIdent(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneApprovalBizIdent))
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

	var l []FortuneApprovalBizIdent
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

// UpdateFortuneApprovalBizIdent updates FortuneApprovalBizIdent by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneApprovalBizIdentById(m *FortuneApprovalBizIdent) (err error) {
	o := orm.NewOrm()
	v := FortuneApprovalBizIdent{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneApprovalBizIdent deletes FortuneApprovalBizIdent by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneApprovalBizIdent(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneApprovalBizIdent{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneApprovalBizIdent{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
