package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneNodeButton struct {
	Id           int       `orm:"column(id);pk"`
	NodeId       string    `orm:"column(node_id);size(36)" description:"所属节点ID"`
	Name         string    `orm:"column(name);size(20)" description:"名称"`
	Status       int8      `orm:"column(status);null" description:"是否可见 0否 1可见"`
	Sort         uint8     `orm:"column(sort);null" description:"排序"`
	DateEntered  time.Time `orm:"column(date_entered);type(datetime);null" description:"创建时间"`
	DateModified time.Time `orm:"column(date_modified);type(datetime);null" description:"修改时间"`
	Deleted      int8      `orm:"column(deleted)" description:"是否删除 0否 1是"`
	IsUrl        int8      `orm:"column(is_url);null" description:"是否为url 0否 1是"`
	Url          string    `orm:"column(url);size(100)" description:"节点对应url"`
	Bid          string    `orm:"column(bid);size(36)" description:"所属按钮ID"`
}

func (t *FortuneNodeButton) TableName() string {
	return "fortune_node_button"
}

func init() {
	orm.RegisterModel(new(FortuneNodeButton))
}

// AddFortuneNodeButton insert a new FortuneNodeButton into database and returns
// last inserted Id on success.
func AddFortuneNodeButton(m *FortuneNodeButton) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneNodeButtonById retrieves FortuneNodeButton by Id. Returns error if
// Id doesn't exist
func GetFortuneNodeButtonById(id int) (v *FortuneNodeButton, err error) {
	o := orm.NewOrm()
	v = &FortuneNodeButton{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneNodeButton retrieves all FortuneNodeButton matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneNodeButton(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneNodeButton))
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

	var l []FortuneNodeButton
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

// UpdateFortuneNodeButton updates FortuneNodeButton by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneNodeButtonById(m *FortuneNodeButton) (err error) {
	o := orm.NewOrm()
	v := FortuneNodeButton{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneNodeButton deletes FortuneNodeButton by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneNodeButton(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneNodeButton{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneNodeButton{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
