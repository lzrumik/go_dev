package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCustomersCopy struct {
	Id                           int       `orm:"column(id);pk"`
	FullName                     string    `orm:"column(full_name);size(100)" description:"姓名"`
	PhonePrefix                  string    `orm:"column(phone_prefix);size(10)" description:"手机号前缀 国家码"`
	PhoneMobile                  string    `orm:"column(phone_mobile);size(50)" description:"手机号"`
	Channel                      uint8     `orm:"column(channel)" description:"渠道 1:线上 2:线下"`
	LockDate                     time.Time `orm:"column(lock_date);type(datetime)" description:"绑定日期"`
	IsBindSam                    uint8     `orm:"column(is_bind_sam)" description:"是否绑定客户经理 1:是 2:否"`
	BindSamId                    string    `orm:"column(bind_sam_id);size(36)" description:"客户经理ID"`
	BindGroupManagerId           string    `orm:"column(bind_group_manager_id);size(36)" description:"团队经理"`
	BindCityManagerId            string    `orm:"column(bind_city_manager_id);size(36)" description:"城市经理"`
	InvestAmount                 float64   `orm:"column(invest_amount);digits(12);decimals(2)" description:"累计投资金额"`
	DollarInvestAmount           float64   `orm:"column(dollar_invest_amount);digits(12);decimals(2)" description:"美元累计投资金额"`
	CustomerTaskId               string    `orm:"column(customer_task_id);size(36)" description:"客户分配任务ID"`
	Status                       uint8     `orm:"column(status)" description:"状态(区分我的客户和我的锁定客户) 0:无 1:锁定 2:正式 3:解绑"`
	RegisterDatetime             time.Time `orm:"column(register_datetime);type(datetime)" description:"国内注册时间"`
	UserId                       string    `orm:"column(user_id);size(36)" description:"国内注册用户ID"`
	OverseasRegsiterDatetime     time.Time `orm:"column(overseas_regsiter_datetime);type(datetime)" description:"海外注册时间"`
	OverseasUserId               string    `orm:"column(overseas_user_id);size(36)" description:"海外注册用户ID"`
	OverseasIdentity             uint8     `orm:"column(overseas_identity)" description:"海外身份 0:无 1:国际 2:美国"`
	AccountStatus                uint8     `orm:"column(account_status)" description:"国内开户状态 0:否 1:是"`
	InternationalAccountStatus   uint8     `orm:"column(international_account_status)" description:"国际开户状态 0:否 1:是"`
	AmericaAccountStatus         uint8     `orm:"column(america_account_status)" description:"美国开户状态 0:否 1:是"`
	AccountDatetime              time.Time `orm:"column(account_datetime);type(datetime)" description:"国内开户时间"`
	InternationalAccountDatetime time.Time `orm:"column(international_account_datetime);type(datetime)" description:"国际开户时间"`
	AmericaAccountDatetime       time.Time `orm:"column(america_account_datetime);type(datetime)" description:"美国开户时间"`
	DateEntered                  time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	DateModified                 time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	Deleted                      uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneCustomersCopy) TableName() string {
	return "fortune_customers_copy"
}

func init() {
	orm.RegisterModel(new(FortuneCustomersCopy))
}

// AddFortuneCustomersCopy insert a new FortuneCustomersCopy into database and returns
// last inserted Id on success.
func AddFortuneCustomersCopy(m *FortuneCustomersCopy) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCustomersCopyById retrieves FortuneCustomersCopy by Id. Returns error if
// Id doesn't exist
func GetFortuneCustomersCopyById(id int) (v *FortuneCustomersCopy, err error) {
	o := orm.NewOrm()
	v = &FortuneCustomersCopy{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCustomersCopy retrieves all FortuneCustomersCopy matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCustomersCopy(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCustomersCopy))
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

	var l []FortuneCustomersCopy
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

// UpdateFortuneCustomersCopy updates FortuneCustomersCopy by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCustomersCopyById(m *FortuneCustomersCopy) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomersCopy{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCustomersCopy deletes FortuneCustomersCopy by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCustomersCopy(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCustomersCopy{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCustomersCopy{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
