package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneMsg struct {
	Id          int       `orm:"column(id);auto"`
	SendType    uint8     `orm:"column(send_type)" description:"发送方式"`
	Scene       int       `orm:"column(scene)" description:"发送场景"`
	Content     string    `orm:"column(content);size(500)" description:"短信内容"`
	Status      int8      `orm:"column(status)" description:"默认有效，1有效，2无效"`
	CreateUser  string    `orm:"column(create_user);size(50)" description:"创建人"`
	ModifyUser  string    `orm:"column(modify_user);size(50)" description:"修改人"`
	Mark        string    `orm:"column(mark);size(100)" description:"备注"`
	DateEntered time.Time `orm:"column(date_entered);type(datetime)"`
	DateModify  time.Time `orm:"column(date_modify);type(datetime)"`
}

func (t *FortuneMsg) TableName() string {
	return "fortune_msg"
}

func init() {
	orm.RegisterModel(new(FortuneMsg))
}

// AddFortuneMsg insert a new FortuneMsg into database and returns
// last inserted Id on success.
func AddFortuneMsg(m *FortuneMsg) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneMsgById retrieves FortuneMsg by Id. Returns error if
// Id doesn't exist
func GetFortuneMsgById(id int) (v *FortuneMsg, err error) {
	o := orm.NewOrm()
	v = &FortuneMsg{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneMsg retrieves all FortuneMsg matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneMsg(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneMsg))
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

	var l []FortuneMsg
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

// UpdateFortuneMsg updates FortuneMsg by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneMsgById(m *FortuneMsg) (err error) {
	o := orm.NewOrm()
	v := FortuneMsg{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneMsg deletes FortuneMsg by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneMsg(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneMsg{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneMsg{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
