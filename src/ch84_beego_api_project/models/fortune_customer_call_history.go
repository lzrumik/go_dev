package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerCallHistory struct {
	Id                      int       `orm:"column(id);pk" description:"沟通历史id"`
	CustomerId              string    `orm:"column(customer_id);size(36)" description:"客户关系id"`
	MarketUnbindWorkorderId string    `orm:"column(market_unbind_workorder_id);size(36)" description:"未绑定意向用户id"`
	UserId                  string    `orm:"column(user_id);size(36)" description:"用户id"`
	OrderSource             int8      `orm:"column(order_source)" description:"工单来源"`
	CallDate                time.Time `orm:"column(call_date);type(datetime)" description:"沟通时间"`
	CallType                string    `orm:"column(call_type);size(30)" description:"沟通方式（电话=phone、微信=wechat、陪访=accompany、开放日=open）"`
	ConnectStatus           uint8     `orm:"column(connect_status)" description:"接通状态 0:历史数据 1:未接通 2:已接通"`
	CallOnWorker            string    `orm:"column(call_on_worker);size(200)" description:"陪访人员"`
	CallContent             string    `orm:"column(call_content);size(2000)" description:"沟通内容"`
	CallPlanId              string    `orm:"column(call_plan_id);size(36)" description:"沟通计划id"`
	PlanCallDate            time.Time `orm:"column(plan_call_date);type(datetime)" description:"预计沟通时间"`
	Expect                  string    `orm:"column(expect);size(2000)" description:"预期"`
	IsIntention             int8      `orm:"column(is_intention)" description:"是否意向 1=是 2=否"`
	IntentionExtent         uint8     `orm:"column(intention_extent)" description:"意向程度 0:历史数据 1:无意向 2:意向一般 3:意向强烈"`
	IntentionDetailDesc     uint16    `orm:"column(intention_detail_desc)" description:"意向详情说明"`
	IntentionProductType    string    `orm:"column(intention_product_type);size(50)" description:"意向产品类型"`
	Remark                  string    `orm:"column(remark);size(500)" description:"备注"`
	NextCommunicationTime   time.Time `orm:"column(next_communication_time);type(datetime)" description:"预计下次沟通时间"`
	CustomerManagerName     string    `orm:"column(customer_manager_name);size(30)" description:"客户经理姓名"`
	CustomerManagerUsername string    `orm:"column(customer_manager_username);size(30)" description:"客户经理账号"`
	DateEntered             time.Time `orm:"column(date_entered);type(datetime)"`
	CreatedUserId           string    `orm:"column(created_user_id);size(36)" description:"创建者"`
	Deleted                 int8      `orm:"column(deleted)" description:"是否删除"`
}

func (t *FortuneCustomerCallHistory) TableName() string {
	return "fortune_customer_call_history"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerCallHistory))
}

// AddFortuneCustomerCallHistory insert a new FortuneCustomerCallHistory into database and returns
// last inserted Id on success.
func AddFortuneCustomerCallHistory(m *FortuneCustomerCallHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerCallHistoryById retrieves FortuneCustomerCallHistory by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerCallHistoryById(id int) (v *FortuneCustomerCallHistory, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerCallHistory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerCallHistory retrieves all FortuneCustomerCallHistory matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerCallHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerCallHistory))
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

	var l []FortuneCustomerCallHistory
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

// UpdateFortuneCustomerCallHistory updates FortuneCustomerCallHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerCallHistoryById(m *FortuneCustomerCallHistory) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerCallHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerCallHistory deletes FortuneCustomerCallHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerCallHistory(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerCallHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerCallHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
