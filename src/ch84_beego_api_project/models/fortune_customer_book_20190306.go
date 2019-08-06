package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomerBook20190306 struct {
	Id                         int       `orm:"column(id);pk" description:"预约ID"`
	UserId                     string    `orm:"column(user_id);size(36)" description:"用户ID"`
	CcAllocateRefId            string    `orm:"column(cc_allocate_ref_id);size(36)" description:"电销客服分配关联id"`
	UserMobilePrefix           string    `orm:"column(user_mobile_prefix);size(10)" description:"用户手机号前缀"`
	UserMobile                 string    `orm:"column(user_mobile);size(20)" description:"用户手机号"`
	UserFullName               string    `orm:"column(user_full_name);size(20)" description:"用户开户姓名"`
	Name                       string    `orm:"column(name);size(30)" description:"用户填写的姓名"`
	PrefixMobile               string    `orm:"column(prefix_mobile);size(10)" description:"用户填写的手机号前缀"`
	Mobile                     string    `orm:"column(mobile);size(50)" description:"用户填写的手机号"`
	CfAccountName              string    `orm:"column(cf_account_name);size(100)" description:"财富开户姓名"`
	LoanId                     string    `orm:"column(loan_id);size(50)" description:"标的名称"`
	OutInvest                  int8      `orm:"column(out_invest)" description:"1是,2否"`
	InvestType                 string    `orm:"column(invest_type);size(20)" description:"投资类型"`
	InvestMoney                int8      `orm:"column(invest_money)" description:"可投资产"`
	BindSamId                  string    `orm:"column(bind_sam_id);size(36)" description:"客户经理ID"`
	ManagerName                string    `orm:"column(manager_name);size(30)" description:"客户经理"`
	Status                     int8      `orm:"column(status)" description:"0未处理,1跟进中,2已完成"`
	CommitIsBindSam            string    `orm:"column(commit_is_bind_sam);size(50);null" description:"预约时是否有对接客户经理"`
	CommitBinSamName           string    `orm:"column(commit_bin_sam_name);size(50);null" description:"预约时对接客户经理姓名"`
	UserMobileLocation         string    `orm:"column(user_mobile_location);size(200)" description:"注册手机号归属地"`
	LoanName                   string    `orm:"column(loan_name);size(300)" description:"标的名称"`
	BindSamName                string    `orm:"column(bind_sam_name);size(100)" description:"绑定rm姓名"`
	AllotDate                  time.Time `orm:"column(allot_date);type(datetime);null" description:"分配时间"`
	AllotStatus                int8      `orm:"column(allot_status)" description:"分配状态 0=未分配（历史数据）、1=已分配、2=分配失败"`
	IpArea                     string    `orm:"column(ip_area);size(200)" description:"ip区域"`
	SubscribeQuestionSnapshoot string    `orm:"column(subscribe_question_snapshoot)" description:"预约题快照"`
	SubscribeEvaluationResult  string    `orm:"column(subscribe_evaluation_result);size(2000)" description:"预约测评结果"`
	Remark                     string    `orm:"column(remark);size(100)" description:"备注"`
	RepairRemark               string    `orm:"column(repair_remark);size(100)" description:"备注2"`
	DateEntered                time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified               time.Time `orm:"column(date_modified);type(datetime)"`
	OpUser                     string    `orm:"column(op_user);size(20)" description:"操作人"`
	Deleted                    int8      `orm:"column(deleted)" description:"删除 0 未删除,1删除"`
}

func (t *FortuneCustomerBook20190306) TableName() string {
	return "fortune_customer_book_20190306"
}

func init() {
	orm.RegisterModel(new(FortuneCustomerBook20190306))
}

// AddFortuneCustomerBook20190306 insert a new FortuneCustomerBook20190306 into database and returns
// last inserted Id on success.
func AddFortuneCustomerBook20190306(m *FortuneCustomerBook20190306) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomerBook20190306ById retrieves FortuneCustomerBook20190306 by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomerBook20190306ById(id int) (v *FortuneCustomerBook20190306, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomerBook20190306{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomerBook20190306 retrieves all FortuneCustomerBook20190306 matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomerBook20190306(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomerBook20190306))
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

	var l []FortuneCustomerBook20190306
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

// UpdateFortuneCustomerBook20190306 updates FortuneCustomerBook20190306 by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomerBook20190306ById(m *FortuneCustomerBook20190306) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerBook20190306{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomerBook20190306 deletes FortuneCustomerBook20190306 by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomerBook20190306(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomerBook20190306{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomerBook20190306{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
