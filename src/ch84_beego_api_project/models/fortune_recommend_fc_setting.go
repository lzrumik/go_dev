package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneRecommendFcSetting struct {
	Id                 int       `orm:"column(set_id);auto"`
	SettingType        uint8     `orm:"column(setting_type)" description:"配置规则 1手动 2自动"`
	SettingStatus      uint8     `orm:"column(setting_status)" description:"配置状态 1未生效 2生效中 3已过期"`
	AchievementPercent uint8     `orm:"column(achievement_percent)" description:"业绩配比"`
	WorkorderPercent   uint8     `orm:"column(workorder_percent)" description:"业绩配比"`
	MostSetting        string    `orm:"column(most_setting);size(255)" description:"最多展示数量设置"`
	AchievementSetting string    `orm:"column(achievement_setting);size(255)" description:"自动配置-业绩"`
	WorkorderSetting   string    `orm:"column(workorder_setting);size(255)" description:"自动配置-工单"`
	DateBegin          time.Time `orm:"column(date_begin);type(datetime)" description:"自动配置生效时间"`
	DateEntered        time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified       time.Time `orm:"column(date_modified);type(datetime)" description:"更新时间"`
	Deleted            uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneRecommendFcSetting) TableName() string {
	return "fortune_recommend_fc_setting"
}

func init() {
	orm.RegisterModel(new(FortuneRecommendFcSetting))
}

// AddFortuneRecommendFcSetting insert a new FortuneRecommendFcSetting into database and returns
// last inserted Id on success.
func AddFortuneRecommendFcSetting(m *FortuneRecommendFcSetting) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneRecommendFcSettingById retrieves FortuneRecommendFcSetting by Id. Returns error if
// Id doesn't exist
func GetFortuneRecommendFcSettingById(id int) (v *FortuneRecommendFcSetting, err error) {
	o := orm.NewOrm()
	v = &FortuneRecommendFcSetting{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneRecommendFcSetting retrieves all FortuneRecommendFcSetting matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneRecommendFcSetting(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneRecommendFcSetting))
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

	var l []FortuneRecommendFcSetting
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

// UpdateFortuneRecommendFcSetting updates FortuneRecommendFcSetting by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneRecommendFcSettingById(m *FortuneRecommendFcSetting) (err error) {
	o := orm.NewOrm()
	v := FortuneRecommendFcSetting{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneRecommendFcSetting deletes FortuneRecommendFcSetting by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneRecommendFcSetting(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneRecommendFcSetting{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneRecommendFcSetting{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
