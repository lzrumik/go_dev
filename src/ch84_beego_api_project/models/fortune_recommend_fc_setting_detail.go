package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneRecommendFcSettingDetail struct {
	Id               int       `orm:"column(id);auto"`
	SetId            int       `orm:"column(set_id)" description:"配置id"`
	RmId             int       `orm:"column(rm_id)" description:"RM id"`
	City             uint8     `orm:"column(city)" description:"北京 上海 深圳"`
	TotalScore       float64   `orm:"column(total_score);digits(10);decimals(5)" description:"得分"`
	SetOrder         uint16    `orm:"column(set_order)" description:"排名顺序"`
	Achievement      uint      `orm:"column(achievement)" description:"业绩数"`
	AchievementOrder uint16    `orm:"column(achievement_order)" description:"业绩数排名"`
	Workorder        uint16    `orm:"column(workorder)" description:"工单数"`
	WorkorderOrder   uint16    `orm:"column(workorder_order)" description:"工单数排名"`
	DateEntered      time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified     time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	Deleted          uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneRecommendFcSettingDetail) TableName() string {
	return "fortune_recommend_fc_setting_detail"
}

func init() {
	orm.RegisterModel(new(FortuneRecommendFcSettingDetail))
}

// AddFortuneRecommendFcSettingDetail insert a new FortuneRecommendFcSettingDetail into database and returns
// last inserted Id on success.
func AddFortuneRecommendFcSettingDetail(m *FortuneRecommendFcSettingDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneRecommendFcSettingDetailById retrieves FortuneRecommendFcSettingDetail by Id. Returns error if
// Id doesn't exist
func GetFortuneRecommendFcSettingDetailById(id int) (v *FortuneRecommendFcSettingDetail, err error) {
	o := orm.NewOrm()
	v = &FortuneRecommendFcSettingDetail{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneRecommendFcSettingDetail retrieves all FortuneRecommendFcSettingDetail matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneRecommendFcSettingDetail(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneRecommendFcSettingDetail))
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

	var l []FortuneRecommendFcSettingDetail
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

// UpdateFortuneRecommendFcSettingDetail updates FortuneRecommendFcSettingDetail by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneRecommendFcSettingDetailById(m *FortuneRecommendFcSettingDetail) (err error) {
	o := orm.NewOrm()
	v := FortuneRecommendFcSettingDetail{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneRecommendFcSettingDetail deletes FortuneRecommendFcSettingDetail by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneRecommendFcSettingDetail(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneRecommendFcSettingDetail{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneRecommendFcSettingDetail{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
