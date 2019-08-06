package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneCompanyInfo struct {
	Id             int       `orm:"column(id);pk"`
	RemarkName     string    `orm:"column(remark_name);size(100)" description:"企业名称备注"`
	CompanyNo      string    `orm:"column(company_no);size(50)" description:"企业编号"`
	SignatureNo    string    `orm:"column(signature_no);size(50)" description:"电子签章码"`
	Keyword        string    `orm:"column(keyword);size(50)" description:"关键字"`
	Format         uint8     `orm:"column(format)" description:"格式 0:上传印章图 1:系统生成"`
	CompanyName    string    `orm:"column(company_name);size(100)" description:"系统生成使用的企业名称"`
	StampFilePath  string    `orm:"column(stamp_file_path);size(200)" description:"印章图文件路径"`
	StampFileName  string    `orm:"column(stamp_file_name);size(100)" description:"印章图文件名"`
	DateExpired    time.Time `orm:"column(date_expired);type(datetime)" description:"到期时间"`
	CreatedUserId  string    `orm:"column(created_user_id);size(36)"`
	ModifiedUserId string    `orm:"column(modified_user_id);size(36)"`
	DateEntered    time.Time `orm:"column(date_entered);type(datetime)"`
	DateModified   time.Time `orm:"column(date_modified);type(datetime)"`
	Deleted        uint8     `orm:"column(deleted)"`
}

func (t *FortuneCompanyInfo) TableName() string {
	return "fortune_company_info"
}

func init() {
	orm.RegisterModel(new(FortuneCompanyInfo))
}

// AddFortuneCompanyInfo insert a new FortuneCompanyInfo into database and returns
// last inserted Id on success.
func AddFortuneCompanyInfo(m *FortuneCompanyInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneCompanyInfoById retrieves FortuneCompanyInfo by Id. Returns error if
// Id doesn't exist
func GetFortuneCompanyInfoById(id int) (v *FortuneCompanyInfo, err error) {
	o := orm.NewOrm()
	v = &FortuneCompanyInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneCompanyInfo retrieves all FortuneCompanyInfo matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneCompanyInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneCompanyInfo))
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

	var l []FortuneCompanyInfo
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

// UpdateFortuneCompanyInfo updates FortuneCompanyInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneCompanyInfoById(m *FortuneCompanyInfo) (err error) {
	o := orm.NewOrm()
	v := FortuneCompanyInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneCompanyInfo deletes FortuneCompanyInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneCompanyInfo(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneCompanyInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneCompanyInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
