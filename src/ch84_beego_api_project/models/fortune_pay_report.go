package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortunePayReport struct {
	Id                  int       `orm:"column(id);auto"`
	ProdType            int8      `orm:"column(prod_type)" description:"产品类型：1 私募股权；2 私募证券；3 地产基金；4类固定收益；5 信托"`
	ProdId              string    `orm:"column(prod_id);size(36)" description:"产品Id"`
	ProdName            string    `orm:"column(prod_name);size(255)" description:"产品名称"`
	Type                int8      `orm:"column(type)" description:"类型：1 私募；2 非私募"`
	CustomerId          string    `orm:"column(customer_id);size(36)" description:"客户ID"`
	ContractCode        string    `orm:"column(contract_code);size(30);null" description:"合同编号"`
	IdCard              string    `orm:"column(id_card);size(18);null" description:"身份证号"`
	RiskLevel           string    `orm:"column(risk_level);size(15);null" description:"风险等级（取财管  C1/C2/C3/C4/C5）"`
	BankId              int       `orm:"column(bank_id);null" description:"银行id"`
	BankName            string    `orm:"column(bank_name);size(255);null" description:"银行名称"`
	BankCardNo          string    `orm:"column(bank_card_no);size(36);null" description:"银行卡号/打款帐号"`
	BankBranchName      string    `orm:"column(bank_branch_name);size(20);null" description:"支行名称"`
	BankProvinceCode    string    `orm:"column(bank_province_code);size(36);null" description:"开户行所在省"`
	BankProvinceName    string    `orm:"column(bank_province_name);size(36);null" description:"开户行所在省"`
	BankCityCode        string    `orm:"column(bank_city_code);size(36);null" description:"开户行所在市"`
	BankCityName        string    `orm:"column(bank_city_name);size(36);null" description:"开户行所在市"`
	PayAmount           float64   `orm:"column(pay_amount);null" description:"打款金额"`
	PayTime             time.Time `orm:"column(pay_time);type(datetime);null" description:"打款日期"`
	ContracProvinceCode string    `orm:"column(contrac_province_code);size(36);null" description:"签约省"`
	ContracProvinceName string    `orm:"column(contrac_province_name);size(36);null"`
	ContracCityCode     string    `orm:"column(contrac_city_code);size(36);null" description:"签约市"`
	ContracCityName     string    `orm:"column(contrac_city_name);size(36);null"`
	Profit              string    `orm:"column(profit);size(10);null" description:"利率（该客户打款金额，对应的产品利率）"`
	VisitPhone          string    `orm:"column(visit_phone);size(11);null" description:"回访电话"`
	VisitStartTime      time.Time `orm:"column(visit_start_time);type(datetime);null" description:"回访开始时间"`
	VisitEndTime        time.Time `orm:"column(visit_end_time);type(datetime);null" description:"回访结束时间"`
	OwnCityId           int       `orm:"column(own_city_id);null" description:"创建人所属城市"`
	OwnGroupId          int       `orm:"column(own_group_id);null" description:"创建人所属组别"`
	Remark              string    `orm:"column(remark);size(255);null" description:"备注"`
	CreateTime          time.Time `orm:"column(create_time);type(datetime)"`
	UpdateTime          time.Time `orm:"column(update_time);type(datetime)"`
	State               int       `orm:"column(state);null" description:"1,正常"`
	CreateUserName      string    `orm:"column(create_user_name);size(255);null"`
	CreateUserId        int       `orm:"column(create_user_id)" description:"创建人"`
	Ext1                string    `orm:"column(ext1);size(255);null"`
}

func (t *FortunePayReport) TableName() string {
	return "fortune_pay_report"
}

func init() {
	orm.RegisterModel(new(FortunePayReport))
}

// AddFortunePayReport insert a new FortunePayReport into database and returns
// last inserted Id on success.
func AddFortunePayReport(m *FortunePayReport) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortunePayReportById retrieves FortunePayReport by Id. Returns error if
// Id doesn't exist
func GetFortunePayReportById(id int) (v *FortunePayReport, err error) {
	o := orm.NewOrm()
	v = &FortunePayReport{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortunePayReport retrieves all FortunePayReport matches certain condition. Returns empty list if
// no records exist
func GetAllFortunePayReport(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortunePayReport))
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

	var l []FortunePayReport
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

// UpdateFortunePayReport updates FortunePayReport by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortunePayReportById(m *FortunePayReport) (err error) {
	o := orm.NewOrm()
	v := FortunePayReport{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortunePayReport deletes FortunePayReport by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortunePayReport(id int) (err error) {
	o := orm.NewOrm()
	v := FortunePayReport{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortunePayReport{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
