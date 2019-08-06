package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneMarketUnbindWorkorderAllotRecord struct {
	Id                      int       `orm:"column(id);auto"`
	MarketUnbindWorkorderId string    `orm:"column(market_unbind_workorder_id);size(36)" description:"未绑定意向用户id"`
	AllotTime               time.Time `orm:"column(allot_time);type(datetime)" description:"分配时间"`
	AllotStatus             int8      `orm:"column(allot_status)" description:"状态 0=未分配、1=已分配、2=已回收、3=分配失败"`
	AllotInfo               string    `orm:"column(allot_info);size(300)" description:"分配信息"`
	BindSamId               string    `orm:"column(bind_sam_id);size(36)" description:"分配 rm id"`
	BindSamName             string    `orm:"column(bind_sam_name);size(50)" description:"分配 rm 姓名"`
	RetrieveTime            time.Time `orm:"column(retrieve_time);type(datetime)" description:"回收时间"`
	AllotOperaterId         string    `orm:"column(allot_operater_id);size(36)" description:"分配操作人id"`
	AllotOperaterName       string    `orm:"column(allot_operater_name);size(50)" description:"分配操作人姓名"`
	DateEntered             time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	LastDateModified        time.Time `orm:"column(last_date_modified);type(datetime)" description:"最后修改时间"`
	Deleted                 uint8     `orm:"column(deleted)" description:"是否删除"`
}

func (t *FortuneMarketUnbindWorkorderAllotRecord) TableName() string {
	return "fortune_market_unbind_workorder_allot_record"
}

func init() {
	orm.RegisterModel(new(FortuneMarketUnbindWorkorderAllotRecord))
}

// AddFortuneMarketUnbindWorkorderAllotRecord insert a new FortuneMarketUnbindWorkorderAllotRecord into database and returns
// last inserted Id on success.
func AddFortuneMarketUnbindWorkorderAllotRecord(m *FortuneMarketUnbindWorkorderAllotRecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneMarketUnbindWorkorderAllotRecordById retrieves FortuneMarketUnbindWorkorderAllotRecord by Id. Returns error if
// Id doesn't exist
func GetFortuneMarketUnbindWorkorderAllotRecordById(id int) (v *FortuneMarketUnbindWorkorderAllotRecord, err error) {
	o := orm.NewOrm()
	v = &FortuneMarketUnbindWorkorderAllotRecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneMarketUnbindWorkorderAllotRecord retrieves all FortuneMarketUnbindWorkorderAllotRecord matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneMarketUnbindWorkorderAllotRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneMarketUnbindWorkorderAllotRecord))
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

	var l []FortuneMarketUnbindWorkorderAllotRecord
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

// UpdateFortuneMarketUnbindWorkorderAllotRecord updates FortuneMarketUnbindWorkorderAllotRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneMarketUnbindWorkorderAllotRecordById(m *FortuneMarketUnbindWorkorderAllotRecord) (err error) {
	o := orm.NewOrm()
	v := FortuneMarketUnbindWorkorderAllotRecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneMarketUnbindWorkorderAllotRecord deletes FortuneMarketUnbindWorkorderAllotRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneMarketUnbindWorkorderAllotRecord(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneMarketUnbindWorkorderAllotRecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneMarketUnbindWorkorderAllotRecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
