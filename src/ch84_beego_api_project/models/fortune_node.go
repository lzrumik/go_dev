package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneNode struct {
	Id             int       `orm:"column(id);pk"`
	Name           string    `orm:"column(name);size(20)" description:"名称"`
	Url            string    `orm:"column(url);size(100)" description:"节点对应url"`
	Code           string    `orm:"column(code);size(30)" description:"唯一标识码"`
	Status         int8      `orm:"column(status);null" description:"是否可见 0否 1可见"`
	Sort           uint8     `orm:"column(sort);null" description:"排序"`
	Pid            string    `orm:"column(pid);size(36)" description:"父级ID"`
	IsParent       int8      `orm:"column(is_parent);null" description:"是否为父节点 0否 1是"`
	CreatedUserId  string    `orm:"column(created_user_id);size(36);null" description:"创建人ID"`
	ModifiedUserId string    `orm:"column(modified_user_id);size(36);null" description:"修改人ID"`
	DateEntered    time.Time `orm:"column(date_entered);type(datetime);null" description:"创建时间"`
	DateModified   time.Time `orm:"column(date_modified);type(datetime);null" description:"修改时间"`
	Deleted        int8      `orm:"column(deleted)" description:"是否删除 0否 1是"`
}

func (t *FortuneNode) TableName() string {
	return "fortune_node"
}

func init() {
	orm.RegisterModel(new(FortuneNode))
}

// AddFortuneNode insert a new FortuneNode into database and returns
// last inserted Id on success.
func AddFortuneNode(m *FortuneNode) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneNodeById retrieves FortuneNode by Id. Returns error if
// Id doesn't exist
func GetFortuneNodeById(id int) (v *FortuneNode, err error) {
	o := orm.NewOrm()
	v = &FortuneNode{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneNode retrieves all FortuneNode matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneNode(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneNode))
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

	var l []FortuneNode
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

// UpdateFortuneNode updates FortuneNode by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneNodeById(m *FortuneNode) (err error) {
	o := orm.NewOrm()
	v := FortuneNode{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneNode deletes FortuneNode by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneNode(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneNode{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneNode{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
