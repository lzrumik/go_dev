package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneBankAuthInfo struct {
	Id                   int       `orm:"column(id);auto"`
	UserId               string    `orm:"column(user_id);size(36)" description:"用户userId"`
	Name                 string    `orm:"column(name);size(30)" description:"真实姓名"`
	IdNum                string    `orm:"column(id_num);size(20)" description:"身份证号"`
	BankNum              string    `orm:"column(bank_num);size(30)" description:"银行卡号"`
	Mobile               string    `orm:"column(mobile);size(18)" description:"银行预留手机号"`
	BankName             string    `orm:"column(bank_name);size(255)" description:"开户行"`
	BankProvince         string    `orm:"column(bank_province);size(255)" description:"省"`
	BankCity             string    `orm:"column(bank_city);size(255)" description:"市"`
	BankId               int       `orm:"column(bank_id)" description:"银行对应的id"`
	Platform             string    `orm:"column(platform);size(100)" description:"开户来源"`
	AddTime              time.Time `orm:"column(add_time);type(datetime)"`
	UpdateTime           time.Time `orm:"column(update_time);type(datetime)"`
	CopyId               int       `orm:"column(copy_id)" description:"复制的记录id"`
	Status               int8      `orm:"column(status)" description:"状态位,默认0认证完成,1更换中(新版本为审核中),2更换失败(新版本为审核失败),3认证中,4添加其他银行卡完成,5修改手机号"`
	ApproveFlowId        int       `orm:"column(approve_flow_id)" description:"审批流程ID"`
	ApproveTaskId        int       `orm:"column(approve_task_id)" description:"审批任务ID"`
	NowApproveProcessId  int       `orm:"column(now_approve_process_id)" description:"当前审批节点ID"`
	NowApproveHistoryId  int       `orm:"column(now_approve_history_id)" description:"当前任务审批节点ID"`
	AuditStatus          int8      `orm:"column(audit_status)" description:"为了兼容修改手机号，0通过,1审核中，2审核失败"`
	Reason               int8      `orm:"column(reason)" description:"拒绝原因"`
	Address              string    `orm:"column(address);size(100)" description:"通信地址"`
	Mark                 string    `orm:"column(mark);size(500)" description:"说明"`
	Type                 int8      `orm:"column(type)" description:"业务类型，1开户,2绑卡，3修改手机号"`
	Source               int8      `orm:"column(source)" description:"申请来源，1客户端，2后台,3pad端"`
	City                 int8      `orm:"column(city)" description:"提交时的城市"`
	BindSamId            string    `orm:"column(bind_sam_id);size(36)" description:"该条记录提交时的id"`
	FailRetry            int8      `orm:"column(fail_retry)" description:"记录失败后是否再提交"`
	CurrentBindSamId     string    `orm:"column(current_bind_sam_id);size(36)" description:"当前绑定客户经理ID"`
	CurrentBindSamCityId string    `orm:"column(current_bind_sam_city_id);size(36)" description:"当前绑定客户经理所属城市ID"`
	CreateUser           string    `orm:"column(create_user);size(255)" description:"创建人"`
	UpdateUser           string    `orm:"column(update_user);size(255)" description:"修改人"`
	Deleted              int8      `orm:"column(deleted)"`
}

func (t *FortuneBankAuthInfo) TableName() string {
	return "fortune_bank_auth_info"
}

func init() {
	orm.RegisterModel(new(FortuneBankAuthInfo))
}

// AddFortuneBankAuthInfo insert a new FortuneBankAuthInfo into database and returns
// last inserted Id on success.
func AddFortuneBankAuthInfo(m *FortuneBankAuthInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneBankAuthInfoById retrieves FortuneBankAuthInfo by Id. Returns error if
// Id doesn't exist
func GetFortuneBankAuthInfoById(id int) (v *FortuneBankAuthInfo, err error) {
	o := orm.NewOrm()
	v = &FortuneBankAuthInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneBankAuthInfo retrieves all FortuneBankAuthInfo matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneBankAuthInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneBankAuthInfo))
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

	var l []FortuneBankAuthInfo
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

// UpdateFortuneBankAuthInfo updates FortuneBankAuthInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneBankAuthInfoById(m *FortuneBankAuthInfo) (err error) {
	o := orm.NewOrm()
	v := FortuneBankAuthInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneBankAuthInfo deletes FortuneBankAuthInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneBankAuthInfo(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneBankAuthInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneBankAuthInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
