package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneActivityExtend struct {
	Id             int       `orm:"column(id);auto" description:"主键"`
	ActivityNo     string    `orm:"column(activity_no);size(36)" description:"活动编号"`
	SponsorCityId  int       `orm:"column(sponsor_city_id);null" description:"活动主办城市id"`
	SponsorCity    string    `orm:"column(sponsor_city);size(255);null" description:"活动主办城市"`
	SponsorGroupId int       `orm:"column(sponsor_group_id);null" description:"活动主办组id"`
	SponsorGroup   string    `orm:"column(sponsor_group);size(255);null" description:"活动主办组"`
	ExpectMinCount int       `orm:"column(expect_min_count)" description:"最小预计参与人数"`
	ExpectMaxCount int       `orm:"column(expect_max_count)" description:"最大预计参与人数"`
	SurplusQuota   int       `orm:"column(surplus_quota)" description:"剩余名额：（预计参与总人数的最大值 - 本组已经审批通过的参加的用户人数）"`
	Creator        int       `orm:"column(creator);null" description:"创建人"`
	CreateTime     time.Time `orm:"column(create_time);type(datetime);auto_now_add" description:"创建时间"`
	LastUpdateTime time.Time `orm:"column(last_update_time);type(datetime);auto_now_add" description:"最后更新时间"`
}

func (t *FortuneActivityExtend) TableName() string {
	return "fortune_activity_extend"
}

func init() {
	orm.RegisterModel(new(FortuneActivityExtend))
}

// AddFortuneActivityExtend insert a new FortuneActivityExtend into database and returns
// last inserted Id on success.
func AddFortuneActivityExtend(m *FortuneActivityExtend) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneActivityExtendById retrieves FortuneActivityExtend by Id. Returns error if
// Id doesn't exist
func GetFortuneActivityExtendById(id int) (v *FortuneActivityExtend, err error) {
	o := orm.NewOrm()
	v = &FortuneActivityExtend{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneActivityExtend retrieves all FortuneActivityExtend matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneActivityExtend(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneActivityExtend))
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

	var l []FortuneActivityExtend
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

// UpdateFortuneActivityExtend updates FortuneActivityExtend by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneActivityExtendById(m *FortuneActivityExtend) (err error) {
	o := orm.NewOrm()
	v := FortuneActivityExtend{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneActivityExtend deletes FortuneActivityExtend by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneActivityExtend(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneActivityExtend{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneActivityExtend{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
