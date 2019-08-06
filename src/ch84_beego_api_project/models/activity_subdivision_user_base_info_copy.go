package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ActivitySubdivisionUserBaseInfoCopy struct {
	Id                 int       `orm:"column(uid);pk" description:"int型uid"`
	UserId             string    `orm:"column(user_id);size(36)"`
	UserName           string    `orm:"column(user_name);size(50)" description:"用户名"`
	FullName           string    `orm:"column(full_name);size(20)" description:"姓名"`
	Gender             string    `orm:"column(gender);size(10)" description:"性别 "`
	Level              int8      `orm:"column(level)" description:"层级 1:普通用户 2:白金用户 3:钻石用户"`
	LevelName          string    `orm:"column(level_name);size(200)" description:"层级名称 普通用户 白金用户 钻石用户"`
	TotalPoints        int       `orm:"column(total_points)" description:"总积分"`
	UpgradePoints      int       `orm:"column(upgrade_points);null" description:"升级积分"`
	AvailablePoints    int       `orm:"column(available_points)" description:"可用积分"`
	PhoneMobile        string    `orm:"column(phone_mobile);size(20)" description:"联系电话"`
	BirthDate          time.Time `orm:"column(birth_date);type(date)" description:"出生年月"`
	RegisterDate       time.Time `orm:"column(register_date);type(date)" description:"注册时间"`
	Address            string    `orm:"column(address);size(200)" description:"所在地"`
	FundFirstTime      time.Time `orm:"column(fund_first_time);type(date)" description:"首次投资时间"`
	FundLastestTime    time.Time `orm:"column(fund_lastest_time);type(date)" description:"最近一次投资时间"`
	FundBrushstrokeMax float64   `orm:"column(fund_brushstroke_max);digits(12);decimals(2)" description:"最高单笔投资金额"`
	FundBrushstrokeMin float64   `orm:"column(fund_brushstroke_min);digits(12);decimals(2)" description:"最低单笔投资金额"`
	FundAllAmount      float64   `orm:"column(fund_all_amount);digits(12);decimals(2)" description:"投资总金额"`
	FundAllCount       int       `orm:"column(fund_all_count)" description:"投资次数"`
	ChannelFirst       string    `orm:"column(channel_first);size(200)" description:"第一渠道"`
	ChannelSecond      string    `orm:"column(channel_second);size(200)" description:"第二渠道"`
	Age                int8      `orm:"column(age)" description:"年龄"`
	IsOpenAccount      int8      `orm:"column(is_open_account)" description:"是否开通托管账户 0否 1是"`
	OpenAccountTime    time.Time `orm:"column(open_account_time);type(date)" description:"开通托管账户时间"`
	IsBindBank         int8      `orm:"column(is_bind_bank)" description:"是否绑定银行卡"`
	BindBankTime       time.Time `orm:"column(bind_bank_time);type(date)" description:"绑定银行卡时间"`
	AccountBalance     float64   `orm:"column(account_balance);digits(12);decimals(2)" description:"账户余额"`
	FirstRechargeTime  time.Time `orm:"column(first_recharge_time);type(date)" description:"首次充值时间"`
	FirstWithdrawTime  time.Time `orm:"column(first_withdraw_time);type(date)" description:"首次提现时间"`
	LastWithdrawTime   time.Time `orm:"column(last_withdraw_time);type(date)" description:"最后一次提现时间（取最近一次）"`
	AccRepayAmount     float64   `orm:"column(acc_repay_amount);digits(12);decimals(2)" description:"累计回款金额"`
	AccRechargeAmount  float64   `orm:"column(acc_recharge_amount);digits(12);decimals(2)" description:"累计充值总额"`
	BaseFields         string    `orm:"column(base_fields);size(2000)" description:"json格式的数据 便于动态变更结果字段"`
}

func (t *ActivitySubdivisionUserBaseInfoCopy) TableName() string {
	return "activity_subdivision_user_base_info_copy"
}

func init() {
	orm.RegisterModel(new(ActivitySubdivisionUserBaseInfoCopy))
}

// AddActivitySubdivisionUserBaseInfoCopy insert a new ActivitySubdivisionUserBaseInfoCopy into database and returns
// last inserted Id on success.
func AddActivitySubdivisionUserBaseInfoCopy(m *ActivitySubdivisionUserBaseInfoCopy) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetActivitySubdivisionUserBaseInfoCopyById retrieves ActivitySubdivisionUserBaseInfoCopy by Id. Returns error if
// Id doesn't exist
func GetActivitySubdivisionUserBaseInfoCopyById(id int) (v *ActivitySubdivisionUserBaseInfoCopy, err error) {
	o := orm.NewOrm()
	v = &ActivitySubdivisionUserBaseInfoCopy{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllActivitySubdivisionUserBaseInfoCopy retrieves all ActivitySubdivisionUserBaseInfoCopy matches certain condition. Returns empty list if
// no records exist
func GetAllActivitySubdivisionUserBaseInfoCopy(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ActivitySubdivisionUserBaseInfoCopy))
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

	var l []ActivitySubdivisionUserBaseInfoCopy
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

// UpdateActivitySubdivisionUserBaseInfoCopy updates ActivitySubdivisionUserBaseInfoCopy by Id and returns error if
// the record to be updated doesn't exist
func UpdateActivitySubdivisionUserBaseInfoCopyById(m *ActivitySubdivisionUserBaseInfoCopy) (err error) {
	o := orm.NewOrm()
	v := ActivitySubdivisionUserBaseInfoCopy{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteActivitySubdivisionUserBaseInfoCopy deletes ActivitySubdivisionUserBaseInfoCopy by Id and returns error if
// the record to be deleted doesn't exist
func DeleteActivitySubdivisionUserBaseInfoCopy(id int) (err error) {
	o := orm.NewOrm()
	v := ActivitySubdivisionUserBaseInfoCopy{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ActivitySubdivisionUserBaseInfoCopy{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
