package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneOrderDetail struct {
	Id             int       `orm:"column(id);pk"`
	CustomerId     string    `orm:"column(customer_id);size(36)" description:"财富ID"`
	UserId         string    `orm:"column(user_id);size(36)" description:"用户ID"`
	RelevantId     string    `orm:"column(relevant_id);size(36)" description:"关联ID"`
	OrderNo        uint64    `orm:"column(order_no)" description:"订单号"`
	Source         uint8     `orm:"column(source)" description:"订单来源 1:财富 2:私募"`
	Status         uint8     `orm:"column(status)" description:"财富认购状态(1:新建保存 2:提交待审批 3:审批中 4:认购成功 5:认购失败 6:待支付 7:审批拒绝 8:作废) 私募订单状态(21:预约(待完成) 22:确认中 23:确认成功 24:确认失败 25:过期确认失败 26:部分确认成功 27:撤单 28:已失效)"`
	IsDoubleRecord uint8     `orm:"column(is_double_record)" description:"是否双录 0:否 1:是"`
	SubType        uint8     `orm:"column(sub_type)" description:"认购类型 1:线上 2:线下"`
	LoanType       string    `orm:"column(loan_type);size(10)" description:"项目类型 财富(0:人民币固定收益 1:美元固定收益 2:私募股权 3:海外置业 4:家庭保障•财富传承 5:私募证券 6:类固定收益产品) 私募(PE:私募股权 PES:私募证券 LE:地产基金)"`
	CreatedUserId  string    `orm:"column(created_user_id);size(36)"`
	ModifiedUserId string    `orm:"column(modified_user_id);size(36)"`
	DateEntered    time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified   time.Time `orm:"column(date_modified);type(datetime)"`
	Deleted        uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneOrderDetail) TableName() string {
	return "fortune_order_detail"
}

func init() {
	orm.RegisterModel(new(FortuneOrderDetail))
}

// AddFortuneOrderDetail insert a new FortuneOrderDetail into database and returns
// last inserted Id on success.
func AddFortuneOrderDetail(m *FortuneOrderDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneOrderDetailById retrieves FortuneOrderDetail by Id. Returns error if
// Id doesn't exist
func GetFortuneOrderDetailById(id int) (v *FortuneOrderDetail, err error) {
	o := orm.NewOrm()
	v = &FortuneOrderDetail{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneOrderDetail retrieves all FortuneOrderDetail matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneOrderDetail(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneOrderDetail))
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

	var l []FortuneOrderDetail
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

// UpdateFortuneOrderDetail updates FortuneOrderDetail by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneOrderDetailById(m *FortuneOrderDetail) (err error) {
	o := orm.NewOrm()
	v := FortuneOrderDetail{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneOrderDetail deletes FortuneOrderDetail by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneOrderDetail(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneOrderDetail{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneOrderDetail{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
