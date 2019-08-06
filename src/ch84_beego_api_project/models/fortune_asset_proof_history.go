package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneAssetProofHistory struct {
	Id           int       `orm:"column(id);auto"`
	AssetProofId int       `orm:"column(asset_proof_id)"`
	HistoryId    int       `orm:"column(history_id)" description:"审批历史id"`
	UserId       string    `orm:"column(user_id);size(36)" description:"用户user_id"`
	SubmitTime   time.Time `orm:"column(submit_time);type(datetime)" description:"提交审核的时间"`
	AuditName    string    `orm:"column(audit_name);size(50)" description:"审核人"`
	AddTime      time.Time `orm:"column(add_time);type(datetime)"`
	UpdateTime   time.Time `orm:"column(update_time);type(datetime)"`
	Status       int8      `orm:"column(status)" description:"状态 1审核通过,2审核拒绝"`
	Reason       int8      `orm:"column(reason)" description:"拒绝理由"`
	Mark         string    `orm:"column(mark);size(300)" description:"备注"`
	Deleted      int8      `orm:"column(deleted)"`
}

func (t *FortuneAssetProofHistory) TableName() string {
	return "fortune_asset_proof_history"
}

func init() {
	orm.RegisterModel(new(FortuneAssetProofHistory))
}

// AddFortuneAssetProofHistory insert a new FortuneAssetProofHistory into database and returns
// last inserted Id on success.
func AddFortuneAssetProofHistory(m *FortuneAssetProofHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneAssetProofHistoryById retrieves FortuneAssetProofHistory by Id. Returns error if
// Id doesn't exist
func GetFortuneAssetProofHistoryById(id int) (v *FortuneAssetProofHistory, err error) {
	o := orm.NewOrm()
	v = &FortuneAssetProofHistory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneAssetProofHistory retrieves all FortuneAssetProofHistory matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneAssetProofHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneAssetProofHistory))
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

	var l []FortuneAssetProofHistory
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

// UpdateFortuneAssetProofHistory updates FortuneAssetProofHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneAssetProofHistoryById(m *FortuneAssetProofHistory) (err error) {
	o := orm.NewOrm()
	v := FortuneAssetProofHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneAssetProofHistory deletes FortuneAssetProofHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneAssetProofHistory(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneAssetProofHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneAssetProofHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
