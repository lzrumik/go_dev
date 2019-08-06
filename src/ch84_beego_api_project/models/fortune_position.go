package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortunePosition struct {
	Id             int       `orm:"column(id);auto"`
	Title          string    `orm:"column(title);size(20)" description:"职位名称"`
	Number         string    `orm:"column(number);size(20)" description:"职位编号"`
	DistrictId     uint16    `orm:"column(district_id)" description:"所属区ID"`
	CityId         uint16    `orm:"column(city_id)" description:"所属市ID"`
	Pid            uint      `orm:"column(pid)" description:"父职位ID"`
	Level          uint8     `orm:"column(level)" description:"级别 1城市经理 2团队经理 3客户经理"`
	CreatedUserId  string    `orm:"column(created_user_id);size(36)" description:"创建人ID"`
	ModifiedUserId string    `orm:"column(modified_user_id);size(36)" description:"修改人ID"`
	DateEntered    time.Time `orm:"column(date_entered);type(datetime);auto_now_add" description:"创建时间"`
	DateModified   time.Time `orm:"column(date_modified);type(datetime);auto_now_add" description:"修改时间"`
}

func (t *FortunePosition) TableName() string {
	return "fortune_position"
}

func init() {
	orm.RegisterModel(new(FortunePosition))
}

// AddFortunePosition insert a new FortunePosition into database and returns
// last inserted Id on success.
func AddFortunePosition(m *FortunePosition) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortunePositionById retrieves FortunePosition by Id. Returns error if
// Id doesn't exist
func GetFortunePositionById(id int) (v *FortunePosition, err error) {
	o := orm.NewOrm()
	v = &FortunePosition{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortunePosition retrieves all FortunePosition matches certain condition. Returns empty list if
// no records exist
func GetAllFortunePosition(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortunePosition))
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

	var l []FortunePosition
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

// UpdateFortunePosition updates FortunePosition by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortunePositionById(m *FortunePosition) (err error) {
	o := orm.NewOrm()
	v := FortunePosition{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortunePosition deletes FortunePosition by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortunePosition(id int) (err error) {
	o := orm.NewOrm()
	v := FortunePosition{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortunePosition{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
