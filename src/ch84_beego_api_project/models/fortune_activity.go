package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneActivity struct {
	Id                    int       `orm:"column(id);auto" description:"主键"`
	ActivityName          string    `orm:"column(activity_name);size(50)" description:"活动名称"`
	ActivityNo            string    `orm:"column(activity_no);size(36)" description:"活动编号"`
	Type                  int8      `orm:"column(type)" description:"活动类型：1 客户转化；2 获取新客；3 品牌；"`
	Status                int8      `orm:"column(status)" description:"活动状态：1 新建；2 预热；3 活动中；4 完成；5 取消；6 暂停；"`
	Province              string    `orm:"column(province);size(20)" description:"活动所在省份"`
	ProvinceId            int       `orm:"column(province_id);null" description:"活动所在省份id"`
	City                  string    `orm:"column(city);size(20)" description:"活动所在城市"`
	CityId                int       `orm:"column(city_id);null" description:"活动所在城市id"`
	Address               string    `orm:"column(address);size(200)" description:"活动地点"`
	HoldingDay            time.Time `orm:"column(holding_day);type(datetime)" description:"活动举办日"`
	PlanActivityStart     time.Time `orm:"column(plan_activity_start);type(datetime)" description:"计划活动周期开始"`
	PlanActivityEnd       time.Time `orm:"column(plan_activity_end);type(datetime)" description:"计划活动周期结束"`
	ActuallyActivityStart time.Time `orm:"column(actually_activity_start);type(datetime);null" description:"实际活动周期开始"`
	ActuallyActivityEnd   time.Time `orm:"column(actually_activity_end);type(datetime);null" description:"实际活动周期结束"`
	Topic                 string    `orm:"column(topic);size(200)" description:"活动主题"`
	Schedule              string    `orm:"column(schedule);size(2000)" description:"活动日程"`
	UserGroupCondition    int8      `orm:"column(user_group_condition)" description:"用户条件是否存在：0 否；1 是；"`
	TaskId                int       `orm:"column(task_id);null" description:"用户组编号"`
	TaskName              string    `orm:"column(task_name);size(255);null" description:"用户组名称"`
	UserCondition         string    `orm:"column(user_condition);size(255);null" description:"用户条件"`
	ExpectTotalCount      int       `orm:"column(expect_total_count)" description:"预计参与总人数"`
	ActuallyTotalCount    int       `orm:"column(actually_total_count)" description:"实际参与总人数（签到成功总人数）"`
	SurplusQuota          int       `orm:"column(surplus_quota)" description:"活动总剩余名额（预计参与总人数 - 已经报名成功的人数（审批通过））"`
	TargetPerformance     string    `orm:"column(target_performance);size(255)" description:"目标业绩"`
	Creator               int       `orm:"column(creator)" description:"创建人"`
	CreateTime            time.Time `orm:"column(create_time);type(datetime);auto_now_add" description:"创建时间"`
	LastUpdateTime        time.Time `orm:"column(last_update_time);type(datetime);auto_now_add" description:"最后更新时间"`
}

func (t *FortuneActivity) TableName() string {
	return "fortune_activity"
}

func init() {
	orm.RegisterModel(new(FortuneActivity))
}

// AddFortuneActivity insert a new FortuneActivity into database and returns
// last inserted Id on success.
func AddFortuneActivity(m *FortuneActivity) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneActivityById retrieves FortuneActivity by Id. Returns error if
// Id doesn't exist
func GetFortuneActivityById(id int) (v *FortuneActivity, err error) {
	o := orm.NewOrm()
	v = &FortuneActivity{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneActivity retrieves all FortuneActivity matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneActivity(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneActivity))
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

	var l []FortuneActivity
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

// UpdateFortuneActivity updates FortuneActivity by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneActivityById(m *FortuneActivity) (err error) {
	o := orm.NewOrm()
	v := FortuneActivity{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneActivity deletes FortuneActivity by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneActivity(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneActivity{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneActivity{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
