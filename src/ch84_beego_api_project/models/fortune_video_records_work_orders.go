package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneVideoRecordsWorkOrders struct {
	Id                   int       `orm:"column(id);auto"`
	WoNo                 string    `orm:"column(wo_no);size(45);null" description:"工单编号"`
	MobileNum            string    `orm:"column(mobile_num);size(20);null" description:"手机号"`
	OrderNo              string    `orm:"column(order_no);size(45);null" description:"订单编号"`
	CustomerUserId       string    `orm:"column(customer_user_id);size(45);null" description:"财管客户id"`
	CustomerUserName     string    `orm:"column(customer_user_name);size(45);null" description:"财管客户姓名"`
	CreateUserId         string    `orm:"column(create_user_id);size(45);null" description:"创建工单用户id"`
	CreateUserName       string    `orm:"column(create_user_name);size(45);null" description:"创建工单用户姓名"`
	RecordBindRmUserId   string    `orm:"column(record_bind_rm_user_id);size(45);null" description:"创建工单时绑定RM用户id"`
	RecordBindRmUserName string    `orm:"column(record_bind_rm_user_name);size(45);null" description:"创建工单时绑定RM用户姓名"`
	QuestionReason       string    `orm:"column(question_reason);size(512);null" description:"问题原因"`
	QuestionDetail       string    `orm:"column(question_detail);size(512);null" description:"问题明细"`
	ContractCustomer     int8      `orm:"column(contract_customer);null" description:"是否需RM联系用户:0 否 1 是"`
	WoDes                string    `orm:"column(wo_des);size(512);null" description:"问题说明"`
	WoStatus             int8      `orm:"column(wo_status);null" description:"状态:1 已完成 2 其他"`
	FileName             string    `orm:"column(file_name);size(128);null" description:"视频文件名称"`
	ProductName          string    `orm:"column(product_name);size(512);null" description:"认购产品名称"`
	CreateTime           time.Time `orm:"column(create_time);type(timestamp);null;auto_now_add" description:"创建时间"`
	Ext1                 string    `orm:"column(ext1);size(20);null" description:"扩展字段1"`
	Ext2                 string    `orm:"column(ext2);size(20);null" description:"扩展字段2"`
}

func (t *FortuneVideoRecordsWorkOrders) TableName() string {
	return "fortune_video_records_work_orders"
}

func init() {
	orm.RegisterModel(new(FortuneVideoRecordsWorkOrders))
}

// AddFortuneVideoRecordsWorkOrders insert a new FortuneVideoRecordsWorkOrders into database and returns
// last inserted Id on success.
func AddFortuneVideoRecordsWorkOrders(m *FortuneVideoRecordsWorkOrders) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneVideoRecordsWorkOrdersById retrieves FortuneVideoRecordsWorkOrders by Id. Returns error if
// Id doesn't exist
func GetFortuneVideoRecordsWorkOrdersById(id int) (v *FortuneVideoRecordsWorkOrders, err error) {
	o := orm.NewOrm()
	v = &FortuneVideoRecordsWorkOrders{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneVideoRecordsWorkOrders retrieves all FortuneVideoRecordsWorkOrders matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneVideoRecordsWorkOrders(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneVideoRecordsWorkOrders))
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

	var l []FortuneVideoRecordsWorkOrders
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

// UpdateFortuneVideoRecordsWorkOrders updates FortuneVideoRecordsWorkOrders by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneVideoRecordsWorkOrdersById(m *FortuneVideoRecordsWorkOrders) (err error) {
	o := orm.NewOrm()
	v := FortuneVideoRecordsWorkOrders{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneVideoRecordsWorkOrders deletes FortuneVideoRecordsWorkOrders by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneVideoRecordsWorkOrders(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneVideoRecordsWorkOrders{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneVideoRecordsWorkOrders{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
