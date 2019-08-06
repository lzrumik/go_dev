package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneUserSubscriptionAttachmentInfo struct {
	Id             int       `orm:"column(attachment_id);pk" description:"主键ID"`
	SubscriptionId string    `orm:"column(subscription_id);size(36)" description:"认购ID"`
	AttachmentType int8      `orm:"column(attachment_type)" description:"1身份证 2协议 3风险测评 4见证声明 5回款凭证 6银行卡上传 "`
	FilePath       string    `orm:"column(file_path);size(100)" description:"文件路径"`
	FileName       string    `orm:"column(file_name);size(100)" description:"文件名"`
	Deleted        int8      `orm:"column(deleted)" description:"状态  0正常  1删除"`
	CreateUserId   string    `orm:"column(create_user_id);size(36)"`
	DateEntered    time.Time `orm:"column(date_entered);type(timestamp)" description:"创建时间"`
}

func (t *FortuneUserSubscriptionAttachmentInfo) TableName() string {
	return "fortune_user_subscription_attachment_info"
}

func init() {
	orm.RegisterModel(new(FortuneUserSubscriptionAttachmentInfo))
}

// AddFortuneUserSubscriptionAttachmentInfo insert a new FortuneUserSubscriptionAttachmentInfo into database and returns
// last inserted Id on success.
func AddFortuneUserSubscriptionAttachmentInfo(m *FortuneUserSubscriptionAttachmentInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneUserSubscriptionAttachmentInfoById retrieves FortuneUserSubscriptionAttachmentInfo by Id. Returns error if
// Id doesn't exist
func GetFortuneUserSubscriptionAttachmentInfoById(id int) (v *FortuneUserSubscriptionAttachmentInfo, err error) {
	o := orm.NewOrm()
	v = &FortuneUserSubscriptionAttachmentInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneUserSubscriptionAttachmentInfo retrieves all FortuneUserSubscriptionAttachmentInfo matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneUserSubscriptionAttachmentInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneUserSubscriptionAttachmentInfo))
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

	var l []FortuneUserSubscriptionAttachmentInfo
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

// UpdateFortuneUserSubscriptionAttachmentInfo updates FortuneUserSubscriptionAttachmentInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneUserSubscriptionAttachmentInfoById(m *FortuneUserSubscriptionAttachmentInfo) (err error) {
	o := orm.NewOrm()
	v := FortuneUserSubscriptionAttachmentInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneUserSubscriptionAttachmentInfo deletes FortuneUserSubscriptionAttachmentInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneUserSubscriptionAttachmentInfo(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneUserSubscriptionAttachmentInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneUserSubscriptionAttachmentInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
