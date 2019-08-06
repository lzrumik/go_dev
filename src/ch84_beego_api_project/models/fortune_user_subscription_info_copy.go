package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneUserSubscriptionInfoCopy struct {
	Id                 int       `orm:"column(subscription_id);pk" description:"主键ID"`
	UserId             string    `orm:"column(user_id);size(36)"`
	ContractNo         string    `orm:"column(contract_no);size(100)" description:"合同编号"`
	SubscriptionDate   time.Time `orm:"column(subscription_date);type(date)" description:"认购日期"`
	LoanId             string    `orm:"column(loan_id);size(36)"`
	SubscriptionAmount float64   `orm:"column(subscription_amount);digits(12);decimals(2)" description:"认购金额"`
	SubscriptionFee    float64   `orm:"column(subscription_fee);digits(4);decimals(2)" description:"认购费"`
	AllAmount          float64   `orm:"column(all_amount);digits(12);decimals(2)" description:"总金额 =认购金额+认购金额*认购费"`
	UserRemitBank      string    `orm:"column(user_remit_bank);size(100)" description:"客户汇款银行"`
	BranchBank         string    `orm:"column(branch_bank);size(100)" description:"分行银行"`
	BankAccount        string    `orm:"column(bank_account);size(30)" description:"银行账户"`
	SwiftCode          string    `orm:"column(swift_code);size(100)" description:"SWIFT CODE"`
	Address            string    `orm:"column(address);size(255)" description:"地址"`
	InfoStatus         int8      `orm:"column(info_status)" description:"状态  1新建保存 2提交待审批 3审批中 4认购成功 5认购失败"`
	CreateUserId       string    `orm:"column(create_user_id);size(36)"`
	DateEntered        time.Time `orm:"column(date_entered);type(timestamp)" description:"创建时间"`
	DateModified       time.Time `orm:"column(date_modified);type(timestamp)" description:"修改时间"`
	DateSubmited       time.Time `orm:"column(date_submited);type(timestamp)" description:"提交时间"`
	Remark             string    `orm:"column(remark);size(255)" description:"备注"`
}

func (t *FortuneUserSubscriptionInfoCopy) TableName() string {
	return "fortune_user_subscription_info_copy"
}

func init() {
	orm.RegisterModel(new(FortuneUserSubscriptionInfoCopy))
}

// AddFortuneUserSubscriptionInfoCopy insert a new FortuneUserSubscriptionInfoCopy into database and returns
// last inserted Id on success.
func AddFortuneUserSubscriptionInfoCopy(m *FortuneUserSubscriptionInfoCopy) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneUserSubscriptionInfoCopyById retrieves FortuneUserSubscriptionInfoCopy by Id. Returns error if
// Id doesn't exist
func GetFortuneUserSubscriptionInfoCopyById(id int) (v *FortuneUserSubscriptionInfoCopy, err error) {
	o := orm.NewOrm()
	v = &FortuneUserSubscriptionInfoCopy{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneUserSubscriptionInfoCopy retrieves all FortuneUserSubscriptionInfoCopy matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneUserSubscriptionInfoCopy(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneUserSubscriptionInfoCopy))
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

	var l []FortuneUserSubscriptionInfoCopy
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

// UpdateFortuneUserSubscriptionInfoCopy updates FortuneUserSubscriptionInfoCopy by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneUserSubscriptionInfoCopyById(m *FortuneUserSubscriptionInfoCopy) (err error) {
	o := orm.NewOrm()
	v := FortuneUserSubscriptionInfoCopy{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneUserSubscriptionInfoCopy deletes FortuneUserSubscriptionInfoCopy by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneUserSubscriptionInfoCopy(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneUserSubscriptionInfoCopy{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneUserSubscriptionInfoCopy{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
