package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneAssetProofAttachment struct {
	Id           int       `orm:"column(id);auto"`
	AssetProofId int       `orm:"column(asset_proof_id)" description:"合格投资者记录id"`
	PicType      int8      `orm:"column(pic_type)" description:"图片类型,1年收入50万,2资产不低于300万，3身份证"`
	PicName      string    `orm:"column(pic_name);size(255)" description:"图片名称"`
	PicUrl       string    `orm:"column(pic_url);size(300)" description:"上传图片的地址"`
	AddTime      time.Time `orm:"column(add_time);type(datetime)"`
	Deleted      int8      `orm:"column(deleted);null"`
}

func (t *FortuneAssetProofAttachment) TableName() string {
	return "fortune_asset_proof_attachment"
}

func init() {
	orm.RegisterModel(new(FortuneAssetProofAttachment))
}

// AddFortuneAssetProofAttachment insert a new FortuneAssetProofAttachment into database and returns
// last inserted Id on success.
func AddFortuneAssetProofAttachment(m *FortuneAssetProofAttachment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneAssetProofAttachmentById retrieves FortuneAssetProofAttachment by Id. Returns error if
// Id doesn't exist
func GetFortuneAssetProofAttachmentById(id int) (v *FortuneAssetProofAttachment, err error) {
	o := orm.NewOrm()
	v = &FortuneAssetProofAttachment{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneAssetProofAttachment retrieves all FortuneAssetProofAttachment matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneAssetProofAttachment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneAssetProofAttachment))
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

	var l []FortuneAssetProofAttachment
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

// UpdateFortuneAssetProofAttachment updates FortuneAssetProofAttachment by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneAssetProofAttachmentById(m *FortuneAssetProofAttachment) (err error) {
	o := orm.NewOrm()
	v := FortuneAssetProofAttachment{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneAssetProofAttachment deletes FortuneAssetProofAttachment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneAssetProofAttachment(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneAssetProofAttachment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneAssetProofAttachment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
