package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneUserSubscriptionInfo struct {
	Id                   int       `orm:"column(subscription_id);pk" description:"主键ID"`
	UserId               string    `orm:"column(user_id);size(36)" description:"Customer ID"`
	OnlineUserId         string    `orm:"column(online_user_id);size(36)" description:"线上用户ID"`
	ContractNo           string    `orm:"column(contract_no);size(100)" description:"合同编号"`
	OrderNo              uint64    `orm:"column(order_no)" description:"订单编号"`
	SubscriptionDate     time.Time `orm:"column(subscription_date);type(date)" description:"认购日期"`
	ReceivedDate         time.Time `orm:"column(received_date);type(date)" description:"到账日期"`
	LoanId               string    `orm:"column(loan_id);size(36)"`
	SubscriptionAmount   float64   `orm:"column(subscription_amount);digits(12);decimals(2)" description:"认购金额"`
	SubscriptionFee      float64   `orm:"column(subscription_fee);digits(12);decimals(2)" description:"认购费"`
	LoanStartTime        time.Time `orm:"column(loan_start_time);type(date)" description:"标的气息日"`
	LoanEndTime          time.Time `orm:"column(loan_end_time);type(date)" description:"标的到期日"`
	AllAmount            float64   `orm:"column(all_amount);digits(12);decimals(2)" description:"总金额 =认购金额+认购金额*认购费"`
	UserRemitBank        string    `orm:"column(user_remit_bank);size(100)" description:"客户汇款银行"`
	BranchBank           string    `orm:"column(branch_bank);size(100)" description:"分行银行"`
	OpenAccountName      string    `orm:"column(open_account_name);size(100)" description:"开户名"`
	BankAccount          string    `orm:"column(bank_account);size(30)" description:"银行账户"`
	SwiftCode            string    `orm:"column(swift_code);size(100)" description:"SWIFT CODE"`
	Address              string    `orm:"column(address);size(255)" description:"地址"`
	InfoStatus           int8      `orm:"column(info_status)" description:"状态  1新建保存 2提交待审批 3审批中 4认购成功 5认购失败 6待支付"`
	CreateUserId         string    `orm:"column(create_user_id);size(36)" description:"创建用户ID APP认购理财师"`
	SubType              int8      `orm:"column(sub_type)" description:"类型  1:线上 2:线下"`
	InterestRate         float64   `orm:"column(interest_rate);digits(12);decimals(2)" description:"利率(业绩基准)"`
	ServiceRate          float64   `orm:"column(service_rate);digits(12);decimals(2)" description:"服务费率"`
	JobMarketAddress     int8      `orm:"column(job_market_address)" description:"职场所在地 "`
	AscriptionCategory   int8      `orm:"column(ascription_category)" description:"销售渠道"`
	DateEntered          time.Time `orm:"column(date_entered);type(timestamp)" description:"创建时间"`
	DateModified         time.Time `orm:"column(date_modified);type(timestamp)" description:"修改时间"`
	DateSubmited         time.Time `orm:"column(date_submited);type(timestamp)" description:"提交时间"`
	DateExcepted         time.Time `orm:"column(date_excepted);type(date)" description:"预计支付时间"`
	Remark               string    `orm:"column(remark);size(255)" description:"备注"`
	IsDoubleRecord       uint8     `orm:"column(is_double_record)" description:"是否双录 0:否 1:是"`
	RecordStatus         uint8     `orm:"column(record_status)" description:"双录审核状态"`
	CurrentApprovalLevel uint8     `orm:"column(current_approval_level)" description:"当前审批级别"`
	LoanPeriodId         uint64    `orm:"column(loan_period_id)" description:"项目期数ID"`
	Deleted              uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
	PredictIncome        float64   `orm:"column(predict_income);digits(12);decimals(2)" description:"预期总收益"`
}

func (t *FortuneUserSubscriptionInfo) TableName() string {
	return "fortune_user_subscription_info"
}

func init() {
	orm.RegisterModel(new(FortuneUserSubscriptionInfo))
}

// AddFortuneUserSubscriptionInfo insert a new FortuneUserSubscriptionInfo into database and returns
// last inserted Id on success.
func AddFortuneUserSubscriptionInfo(m *FortuneUserSubscriptionInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneUserSubscriptionInfoById retrieves FortuneUserSubscriptionInfo by Id. Returns error if
// Id doesn't exist
func GetFortuneUserSubscriptionInfoById(id int) (v *FortuneUserSubscriptionInfo, err error) {
	o := orm.NewOrm()
	v = &FortuneUserSubscriptionInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneUserSubscriptionInfo retrieves all FortuneUserSubscriptionInfo matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneUserSubscriptionInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneUserSubscriptionInfo))
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

	var l []FortuneUserSubscriptionInfo
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

// UpdateFortuneUserSubscriptionInfo updates FortuneUserSubscriptionInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneUserSubscriptionInfoById(m *FortuneUserSubscriptionInfo) (err error) {
	o := orm.NewOrm()
	v := FortuneUserSubscriptionInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneUserSubscriptionInfo deletes FortuneUserSubscriptionInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneUserSubscriptionInfo(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneUserSubscriptionInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneUserSubscriptionInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
