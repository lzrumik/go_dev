package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneBankAuthInfo很危险的 struct {
	Id         int       `orm:"column(id);auto"`
	UserId     string    `orm:"column(user_id);size(36)" description:"用户userId"`
	Name       string    `orm:"column(name);size(30)" description:"真实姓名"`
	IdNum      string    `orm:"column(id_num);size(20)" description:"身份证号"`
	BankNum    string    `orm:"column(bank_num);size(30)" description:"银行卡号"`
	Mobile     string    `orm:"column(mobile);size(18)" description:"银行预留手机号"`
	BankName   string    `orm:"column(bank_name);size(255)" description:"开户行"`
	BankId     int       `orm:"column(bank_id)" description:"银行对应的id"`
	AddTime    time.Time `orm:"column(add_time);type(datetime)"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime)"`
	Status     int8      `orm:"column(status)" description:"状态位,默认0认证完成,1更换中,2更换失败"`
	Reason     int8      `orm:"column(reason)" description:"拒绝原因"`
	Mark       string    `orm:"column(mark);size(500)" description:"说明"`
	Deleted    int8      `orm:"column(deleted)"`
}

func (t *FortuneBankAuthInfo很危险的) TableName() string {
	return "fortune_bank_auth_info_很危险的"
}

func init() {
	orm.RegisterModel(new(FortuneBankAuthInfo很危险的))
}

// AddFortuneBankAuthInfo很危险的 insert a new FortuneBankAuthInfo很危险的 into database and returns
// last inserted Id on success.
func AddFortuneBankAuthInfo很危险的(m *FortuneBankAuthInfo很危险的) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneBankAuthInfo很危险的ById retrieves FortuneBankAuthInfo很危险的 by Id. Returns error if
// Id doesn't exist
func GetFortuneBankAuthInfo很危险的ById(id int) (v *FortuneBankAuthInfo很危险的, err error) {
	o := orm.NewOrm()
	v = &FortuneBankAuthInfo很危险的{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneBankAuthInfo很危险的 retrieves all FortuneBankAuthInfo很危险的 matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneBankAuthInfo很危险的(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneBankAuthInfo很危险的))
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

	var l []FortuneBankAuthInfo很危险的
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

// UpdateFortuneBankAuthInfo很危险的 updates FortuneBankAuthInfo很危险的 by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneBankAuthInfo很危险的ById(m *FortuneBankAuthInfo很危险的) (err error) {
	o := orm.NewOrm()
	v := FortuneBankAuthInfo很危险的{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneBankAuthInfo很危险的 deletes FortuneBankAuthInfo很危险的 by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneBankAuthInfo很危险的(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneBankAuthInfo很危险的{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneBankAuthInfo很危险的{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
