package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneMsgSendrecord struct {
	Id         int       `orm:"column(id);auto"`
	CustomerId string    `orm:"column(customer_id);size(36)" description:"customer_id"`
	Type       int8      `orm:"column(type)" description:"类型，1短信，2邮件"`
	SendType   int8      `orm:"column(send_type)" description:"发送方式,1自动,2手动"`
	SendTo     uint8     `orm:"column(send_to)" description:"1用户，2RM"`
	Contacts   string    `orm:"column(contacts);size(100)" description:"联系方式，手机或邮件"`
	Name       string    `orm:"column(name);size(30)" description:"发送人姓名"`
	Title      string    `orm:"column(title);size(50)" description:"标题"`
	Content    string    `orm:"column(content);size(255)" description:"短信内容"`
	Platform   int8      `orm:"column(platform)" description:"平台,1财富，2私募，3callcenter"`
	Scene      int       `orm:"column(scene)" description:"发送场景"`
	Status     int8      `orm:"column(status)" description:"发送状态，1成功,0失败"`
	AddTime    time.Time `orm:"column(add_time);type(datetime)"`
	ResendTime time.Time `orm:"column(resend_time);type(datetime)"`
	ResendId   uint      `orm:"column(resend_id)" description:"重新发送的id"`
}

func (t *FortuneMsgSendrecord) TableName() string {
	return "fortune_msg_sendrecord"
}

func init() {
	orm.RegisterModel(new(FortuneMsgSendrecord))
}

// AddFortuneMsgSendrecord insert a new FortuneMsgSendrecord into database and returns
// last inserted Id on success.
func AddFortuneMsgSendrecord(m *FortuneMsgSendrecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneMsgSendrecordById retrieves FortuneMsgSendrecord by Id. Returns error if
// Id doesn't exist
func GetFortuneMsgSendrecordById(id int) (v *FortuneMsgSendrecord, err error) {
	o := orm.NewOrm()
	v = &FortuneMsgSendrecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneMsgSendrecord retrieves all FortuneMsgSendrecord matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneMsgSendrecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneMsgSendrecord))
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

	var l []FortuneMsgSendrecord
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

// UpdateFortuneMsgSendrecord updates FortuneMsgSendrecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneMsgSendrecordById(m *FortuneMsgSendrecord) (err error) {
	o := orm.NewOrm()
	v := FortuneMsgSendrecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneMsgSendrecord deletes FortuneMsgSendrecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneMsgSendrecord(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneMsgSendrecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneMsgSendrecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
