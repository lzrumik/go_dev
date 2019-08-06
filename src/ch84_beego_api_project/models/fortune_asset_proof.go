package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneAssetProof struct {
	Id                   int       `orm:"column(id);auto"`
	UserId               string    `orm:"column(user_id);size(36)" description:"用户user_id"`
	Assets               int8      `orm:"column(assets)" description:"1三年收入不低于50w , 2金额不低于300W，默认0"`
	IdPicStatus          int8      `orm:"column(id_pic_status)" description:"身份证图片你上传的状态"`
	AssetPicStatus       int8      `orm:"column(asset_pic_status)" description:"资产证明状态"`
	ExpireTime           int       `orm:"column(expire_time)" description:"过期时间"`
	AddTime              time.Time `orm:"column(add_time);type(datetime)"`
	UpdateTime           time.Time `orm:"column(update_time);type(datetime)"`
	Source               int8      `orm:"column(source)" description:"信息来源，1客户移动端,2rm后台,3pad移动端"`
	Status               int8      `orm:"column(status)" description:"状态 1待审核，2已认证，3认证失败,4认证过期"`
	City                 int8      `orm:"column(city)" description:"提交时的城市"`
	BindSamId            string    `orm:"column(bind_sam_id);size(36)" description:"绑定理财师id"`
	ApproveFlowId        int       `orm:"column(approve_flow_id)" description:"审批流程ID"`
	ApproveTaskId        int       `orm:"column(approve_task_id)" description:"审批任务ID"`
	NowApproveProcessId  int       `orm:"column(now_approve_process_id)" description:"当前审批节点ID"`
	NowApproveHistoryId  int       `orm:"column(now_approve_history_id)" description:"当前任务审批节点ID"`
	Reason               int8      `orm:"column(reason)" description:"拒绝理由"`
	Mark                 string    `orm:"column(mark);size(300)" description:"备注"`
	FailRetry            int8      `orm:"column(fail_retry);null" description:"0非重试,1重试"`
	CurrentBindSamId     string    `orm:"column(current_bind_sam_id);size(36)" description:"当前绑定客户经理ID"`
	CurrentBindSamCityId string    `orm:"column(current_bind_sam_city_id);size(36)" description:"当前绑定客户经理所属城市ID"`
	CreateUser           string    `orm:"column(create_user);size(255)" description:"创建人"`
	UpdateUser           string    `orm:"column(update_user);size(255)" description:"修改人"`
	Deleted              int8      `orm:"column(deleted)"`
}

func (t *FortuneAssetProof) TableName() string {
	return "fortune_asset_proof"
}

func init() {
	orm.RegisterModel(new(FortuneAssetProof))
}

// AddFortuneAssetProof insert a new FortuneAssetProof into database and returns
// last inserted Id on success.
func AddFortuneAssetProof(m *FortuneAssetProof) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneAssetProofById retrieves FortuneAssetProof by Id. Returns error if
// Id doesn't exist
func GetFortuneAssetProofById(id int) (v *FortuneAssetProof, err error) {
	o := orm.NewOrm()
	v = &FortuneAssetProof{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneAssetProof retrieves all FortuneAssetProof matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneAssetProof(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneAssetProof))
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

	var l []FortuneAssetProof
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

// UpdateFortuneAssetProof updates FortuneAssetProof by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneAssetProofById(m *FortuneAssetProof) (err error) {
	o := orm.NewOrm()
	v := FortuneAssetProof{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneAssetProof deletes FortuneAssetProof by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneAssetProof(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneAssetProof{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneAssetProof{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
