package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerChangeRelationHistoryCopy struct {
	Id                      int       `orm:"column(id);pk" description:"变更关系历史id"`
	CustomerId              string    `orm:"column(customer_id);size(36)" description:"客户关系id"`
	BatchNumber             int       `orm:"column(batch_number)" description:"批次号"`
	CustomerManagerName     string    `orm:"column(customer_manager_name);size(30)" description:"客户经理姓名"`
	CustomerManagerUsername string    `orm:"column(customer_manager_username);size(50)" description:"客户经理登录账号"`
	RelationStartDate       time.Time `orm:"column(relation_start_date);type(datetime)" description:"归属关系产生时间"`
	RelationEndDate         time.Time `orm:"column(relation_end_date);type(datetime)" description:"归属关系结束时间"`
	CreateOperate           string    `orm:"column(create_operate);size(30)" description:"创建操作者（系统 ? system : 经理姓名）"`
	ChangeOperate           string    `orm:"column(change_operate);size(30)" description:"变更操作者（系统 ? system : 经理姓名）"`
	DateEntered             time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified            time.Time `orm:"column(date_modified);type(datetime)"`
	LastVersionContent      string    `orm:"column(last_version_content);size(2000)" description:"变更前当前客户的所有信息和变更前的所有变更历史记录"`
	Deleted                 int8      `orm:"column(deleted)" description:"是否删除"`
}

func (t *FortuneCustomerChangeRelationHistoryCopy) TableName() string {
	return "fortune_customer_change_relation_history_copy"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerChangeRelationHistoryCopy))
}

// AddFortuneCustomerChangeRelationHistoryCopy insert a new FortuneCustomerChangeRelationHistoryCopy into database and returns
// last inserted Id on success.
func AddFortuneCustomerChangeRelationHistoryCopy(m *FortuneCustomerChangeRelationHistoryCopy) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerChangeRelationHistoryCopyById retrieves FortuneCustomerChangeRelationHistoryCopy by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerChangeRelationHistoryCopyById(id int) (v *FortuneCustomerChangeRelationHistoryCopy, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerChangeRelationHistoryCopy{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerChangeRelationHistoryCopy retrieves all FortuneCustomerChangeRelationHistoryCopy matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerChangeRelationHistoryCopy(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerChangeRelationHistoryCopy))
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

	var l []FortuneCustomerChangeRelationHistoryCopy
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

// UpdateFortuneCustomerChangeRelationHistoryCopy updates FortuneCustomerChangeRelationHistoryCopy by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerChangeRelationHistoryCopyById(m *FortuneCustomerChangeRelationHistoryCopy) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerChangeRelationHistoryCopy{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerChangeRelationHistoryCopy deletes FortuneCustomerChangeRelationHistoryCopy by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerChangeRelationHistoryCopy(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerChangeRelationHistoryCopy{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerChangeRelationHistoryCopy{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
