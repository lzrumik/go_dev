package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneLoanAttachmentInfo struct {
	Id           int       `orm:"column(id);pk"`
	LoanId       string    `orm:"column(loan_id);size(36)" description:"对应fortune_loan表的id"`
	Type         uint8     `orm:"column(type)" description:"类型 0:其他 1:理财指南 2:产品亮点图片 3:资料包 4:产品合同 5:风险提示函 6:电子签名告知书"`
	FilePath     string    `orm:"column(file_path);size(200)" description:"文件路径"`
	FileName     string    `orm:"column(file_name);size(200)" description:"文件名"`
	TemplateId   string    `orm:"column(template_id);size(36);null" description:"合同服务-模板ID"`
	CreateUserId string    `orm:"column(create_user_id);size(36)"`
	DateEntered  time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified time.Time `orm:"column(date_modified);type(datetime)" description:"修改时间"`
	Deleted      uint8     `orm:"column(deleted)" description:"是否删除 0:否 1:是"`
}

func (t *FortuneLoanAttachmentInfo) TableName() string {
	return "fortune_loan_attachment_info"
}

func init() {
	orm.RegisterModel(new(FortuneLoanAttachmentInfo))
}

// AddFortuneLoanAttachmentInfo insert a new FortuneLoanAttachmentInfo into database and returns
// last inserted Id on success.
func AddFortuneLoanAttachmentInfo(m *FortuneLoanAttachmentInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneLoanAttachmentInfoById retrieves FortuneLoanAttachmentInfo by Id. Returns error if
// Id doesn't exist
func GetFortuneLoanAttachmentInfoById(id int) (v *FortuneLoanAttachmentInfo, err error) {
	o := orm.NewOrm()
	v = &FortuneLoanAttachmentInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneLoanAttachmentInfo retrieves all FortuneLoanAttachmentInfo matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneLoanAttachmentInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneLoanAttachmentInfo))
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

	var l []FortuneLoanAttachmentInfo
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

// UpdateFortuneLoanAttachmentInfo updates FortuneLoanAttachmentInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneLoanAttachmentInfoById(m *FortuneLoanAttachmentInfo) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanAttachmentInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneLoanAttachmentInfo deletes FortuneLoanAttachmentInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneLoanAttachmentInfo(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanAttachmentInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneLoanAttachmentInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
