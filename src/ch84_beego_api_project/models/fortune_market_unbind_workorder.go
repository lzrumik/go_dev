package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneMarketUnbindWorkorder struct {
	Id                    int       `orm:"column(id);pk" description:"id"`
	UserId                string    `orm:"column(user_id);size(36)" description:"用户id"`
	Username              string    `orm:"column(username);size(30)" description:"用户名"`
	FullName              string    `orm:"column(full_name);size(20)" description:"用户姓名"`
	Mobile                string    `orm:"column(mobile);size(15)" description:"手机号"`
	Sex                   string    `orm:"column(sex);size(10)" description:"性别 male=男 female=女"`
	Age                   int8      `orm:"column(age)" description:"年龄"`
	Location              string    `orm:"column(location);size(200)" description:"所在地"`
	WorkorderId           string    `orm:"column(workorder_id);size(36)" description:"工单id"`
	WorkorderCreateTime   time.Time `orm:"column(workorder_create_time);type(datetime)" description:"工单创建时间"`
	WorkorderCalledNumber string    `orm:"column(workorder_called_number);size(50)" description:"拨打坐席工号"`
	WorkorderCalledName   string    `orm:"column(workorder_called_name);size(50)" description:"拨打坐席姓名"`
	WorkorderCalledTime   time.Time `orm:"column(workorder_called_time);type(datetime)" description:"拨打时间"`
	WorkorderCallId       string    `orm:"column(workorder_call_id);size(100)" description:"录音id"`
	WorkorderApproveTime  time.Time `orm:"column(workorder_approve_time);type(datetime)" description:"审批通过时间"`
	LastAllotOperaterId   string    `orm:"column(last_allot_operater_id);size(36)" description:"分配操作人id"`
	LastAllotOperaterName string    `orm:"column(last_allot_operater_name);size(50)" description:"分配操作人姓名"`
	LastAllotTime         time.Time `orm:"column(last_allot_time);type(datetime)" description:"分配时间"`
	LastAllotStatus       int8      `orm:"column(last_allot_status)" description:"状态 0=未分配、1=已分配、2=已回收、3=分配失败"`
	LastAllotInfo         string    `orm:"column(last_allot_info);size(100)" description:"分配信息"`
	BindGroupId           string    `orm:"column(bind_group_id);size(36)" description:"团队组id"`
	BindGroupManagerId    string    `orm:"column(bind_group_manager_id);size(36)" description:"团队组长id"`
	BindGroupManagerName  string    `orm:"column(bind_group_manager_name);size(50)" description:"团队组长姓名"`
	BindCityId            string    `orm:"column(bind_city_id);size(36)" description:"城市id"`
	BindCityManagerId     string    `orm:"column(bind_city_manager_id);size(36)" description:"城市经理id"`
	BindCityManagerName   string    `orm:"column(bind_city_manager_name);size(50)" description:"城市经理姓名"`
	BindSamId             string    `orm:"column(bind_sam_id);size(36)" description:"rm id"`
	BindSamName           string    `orm:"column(bind_sam_name);size(50)" description:"rm 姓名"`
	Remark                string    `orm:"column(remark);size(200)" description:"备注"`
	IsShow                int8      `orm:"column(is_show)" description:"状态 0=不显示、1=显示"`
	DateEntered           time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	LastDateModified      time.Time `orm:"column(last_date_modified);type(datetime)" description:"最后修改时间"`
	Deleted               uint8     `orm:"column(deleted)" description:"是否删除"`
}

func (t *FortuneMarketUnbindWorkorder) TableName() string {
	return "fortune_market_unbind_workorder"
}

func init() {
	orm.RegisterModel(new(FortuneMarketUnbindWorkorder))
}

// AddFortuneMarketUnbindWorkorder insert a new FortuneMarketUnbindWorkorder into database and returns
// last inserted Id on success.
func AddFortuneMarketUnbindWorkorder(m *FortuneMarketUnbindWorkorder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneMarketUnbindWorkorderById retrieves FortuneMarketUnbindWorkorder by Id. Returns error if
// Id doesn't exist
func GetFortuneMarketUnbindWorkorderById(id int) (v *FortuneMarketUnbindWorkorder, err error) {
	o := orm.NewOrm()
	v = &FortuneMarketUnbindWorkorder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneMarketUnbindWorkorder retrieves all FortuneMarketUnbindWorkorder matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneMarketUnbindWorkorder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneMarketUnbindWorkorder))
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

	var l []FortuneMarketUnbindWorkorder
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

// UpdateFortuneMarketUnbindWorkorder updates FortuneMarketUnbindWorkorder by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneMarketUnbindWorkorderById(m *FortuneMarketUnbindWorkorder) (err error) {
	o := orm.NewOrm()
	v := FortuneMarketUnbindWorkorder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneMarketUnbindWorkorder deletes FortuneMarketUnbindWorkorder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneMarketUnbindWorkorder(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneMarketUnbindWorkorder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneMarketUnbindWorkorder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
