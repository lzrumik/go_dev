package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneAppropriateInfo struct {
	Id                         int       `orm:"column(id);auto"`
	UserId                     string    `orm:"column(user_id);size(36)" description:"用户userId"`
	AppropriateInfo            string    `orm:"column(appropriate_info)" description:"适当性信息,json格式"`
	AddTime                    time.Time `orm:"column(add_time);type(datetime)"`
	UpdateTime                 time.Time `orm:"column(update_time);type(datetime)"`
	Source                     int8      `orm:"column(source)" description:"来源1客户端,2rm后端,3pad端"`
	IsEffective                int8      `orm:"column(is_effective)" description:"是否有效，0是，1否"`
	Status                     int8      `orm:"column(status)" description:"状态0成功,1待审核,2失败"`
	City                       int8      `orm:"column(city)" description:"提交时的城市"`
	BindSamId                  string    `orm:"column(bind_sam_id);size(36)" description:"该记录提交时的id"`
	CreateUser                 string    `orm:"column(create_user);size(255)" description:"创建人姓名"`
	CurrentBindSamId           string    `orm:"column(current_bind_sam_id);size(36)" description:"当前绑定客户经理ID"`
	CurrentBindSamCityId       string    `orm:"column(current_bind_sam_city_id);size(36)" description:"当前绑定客户经理所属城市ID"`
	PreCustomerSubscribeStatus int8      `orm:"column(pre_customer_subscribe_status)" description:"理财师分配预处理标识默认0,1未分配,2已分配"`
	ApproveFlowId              int       `orm:"column(approve_flow_id)" description:"审批流程ID"`
	ApproveTaskId              int       `orm:"column(approve_task_id)" description:"审批任务ID"`
	NowApproveProcessId        int       `orm:"column(now_approve_process_id)" description:"当前审批节点ID"`
	NowApproveHistoryId        int       `orm:"column(now_approve_history_id)" description:"当前任务审批节点ID"`
	Deleted                    int8      `orm:"column(deleted)"`
}

func (t *FortuneAppropriateInfo) TableName() string {
	return "fortune_appropriate_info"
}

func init() {
	orm.RegisterModel(new(FortuneAppropriateInfo))
}

// AddFortuneAppropriateInfo insert a new FortuneAppropriateInfo into database and returns
// last inserted Id on success.
func AddFortuneAppropriateInfo(m *FortuneAppropriateInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneAppropriateInfoById retrieves FortuneAppropriateInfo by Id. Returns error if
// Id doesn't exist
func GetFortuneAppropriateInfoById(id int) (v *FortuneAppropriateInfo, err error) {
	o := orm.NewOrm()
	v = &FortuneAppropriateInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneAppropriateInfo retrieves all FortuneAppropriateInfo matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneAppropriateInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneAppropriateInfo))
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

	var l []FortuneAppropriateInfo
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

// UpdateFortuneAppropriateInfo updates FortuneAppropriateInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneAppropriateInfoById(m *FortuneAppropriateInfo) (err error) {
	o := orm.NewOrm()
	v := FortuneAppropriateInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneAppropriateInfo deletes FortuneAppropriateInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneAppropriateInfo(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneAppropriateInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneAppropriateInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
