package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneVideoRecords struct {
	Id                   int       `orm:"column(id);auto"`
	VideoName            string    `orm:"column(video_name);size(255)" description:"双录文件名"`
	SubmitTime           time.Time `orm:"column(submit_time);type(datetime)" description:"提交时间"`
	UserId               string    `orm:"column(user_id);size(36)" description:"用户userId"`
	UserName             string    `orm:"column(user_name);size(50)" description:"用户姓名"`
	PhoneMobile          string    `orm:"column(phone_mobile);size(11)" description:"用户手机号"`
	LoanName             string    `orm:"column(loan_name);size(50)" description:"预约产品名称"`
	OrderNo              string    `orm:"column(order_no);size(50)" description:"订单编号"`
	Type                 int8      `orm:"column(type)" description:"来源: 1财富,2私募"`
	CreatedUser          string    `orm:"column(created_user);size(50)" description:"创建者"`
	ExtraInfo            string    `orm:"column(extra_info)" description:"额外信息"`
	CurrentApprovalLevel int8      `orm:"column(current_approval_level)" description:"财富当前审批级别"`
	ApproveFlowId        int       `orm:"column(approve_flow_id)" description:"审批流程ID"`
	ApproveTaskId        int       `orm:"column(approve_task_id)" description:"审批任务ID"`
	NowApproveProcessId  int       `orm:"column(now_approve_process_id)" description:"当前审批节点ID"`
	NowApproveHistoryId  int       `orm:"column(now_approve_history_id)" description:"当前任务审批节点ID"`
	City                 int8      `orm:"column(city)" description:"提交时的城市"`
	BindSamId            string    `orm:"column(bind_sam_id);size(36)" description:"当前提交时绑定的id"`
	Status               uint8     `orm:"column(status)" description:"审核状态，0默认，1待审核,2审核通过，3审核拒绝"`
	AddTime              time.Time `orm:"column(add_time);type(datetime)"`
	UpdateTime           time.Time `orm:"column(update_time);type(datetime)"`
	Deleted              uint8     `orm:"column(deleted)"`
}

func (t *FortuneVideoRecords) TableName() string {
	return "fortune_video_records"
}

func init() {
	orm.RegisterModel(new(FortuneVideoRecords))
}

// AddFortuneVideoRecords insert a new FortuneVideoRecords into database and returns
// last inserted Id on success.
func AddFortuneVideoRecords(m *FortuneVideoRecords) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneVideoRecordsById retrieves FortuneVideoRecords by Id. Returns error if
// Id doesn't exist
func GetFortuneVideoRecordsById(id int) (v *FortuneVideoRecords, err error) {
	o := orm.NewOrm()
	v = &FortuneVideoRecords{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneVideoRecords retrieves all FortuneVideoRecords matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneVideoRecords(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneVideoRecords))
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

	var l []FortuneVideoRecords
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

// UpdateFortuneVideoRecords updates FortuneVideoRecords by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneVideoRecordsById(m *FortuneVideoRecords) (err error) {
	o := orm.NewOrm()
	v := FortuneVideoRecords{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneVideoRecords deletes FortuneVideoRecords by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneVideoRecords(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneVideoRecords{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneVideoRecords{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
