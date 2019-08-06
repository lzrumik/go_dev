package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneUserTools struct {
	Id                 int       `orm:"column(id);auto"`
	Username           string    `orm:"column(username);size(20)" description:"登录账号"`
	Password           string    `orm:"column(password);size(34)" description:"登录密码"`
	Name               string    `orm:"column(name);size(20)" description:"姓名"`
	Status             int8      `orm:"column(status)" description:"是否生效 0否 1是"`
	Email              string    `orm:"column(email);size(50)" description:"邮箱"`
	Tel                string    `orm:"column(tel);size(15)" description:"电话"`
	Number             string    `orm:"column(number);size(20)" description:"员工编号"`
	Gender             uint8     `orm:"column(gender)" description:"性别 1:男 2:女"`
	LastLogin          time.Time `orm:"column(last_login);type(datetime);null" description:"最近登录时间"`
	LastChangePwd      time.Time `orm:"column(last_change_pwd);type(datetime);null" description:"最近修改密码时间"`
	CreatedUserId      string    `orm:"column(created_user_id);size(36);null" description:"创建人ID"`
	ModifiedUserId     string    `orm:"column(modified_user_id);size(36);null" description:"修改人ID"`
	PNumber            string    `orm:"column(p_number);size(36)" description:"职位编号"`
	OverseasUserId     string    `orm:"column(overseas_user_id);size(50)" description:"海外客户ID"`
	Json               string    `orm:"column(json);size(1000)" description:"海外门店ID下关外客户相关信息"`
	OwnAreaId          uint      `orm:"column(own_area_id)" description:"所属大区"`
	OwnCityId          uint      `orm:"column(own_city_id)" description:"所属城市"`
	CityManageId       uint      `orm:"column(city_manage_id)" description:"城市经理职位ID"`
	CityManageStaffId  uint      `orm:"column(city_manage_staff_id)" description:"城市经理ID"`
	GroupManageId      uint      `orm:"column(group_manage_id)" description:"团队经理职位ID"`
	GroupManageStaffId uint      `orm:"column(group_manage_staff_id)" description:"团队经理ID"`
	CustomerManageId   uint      `orm:"column(customer_manage_id)" description:"客户经理职位ID"`
	Level              int8      `orm:"column(level)" description:"角色 1超级管理员 2管理员 3城市经理 4团队经理 5客户经理"`
	IsShow             int8      `orm:"column(is_show);null" description:"是否前端显示 1=显示  2=不显示"`
	Picture            string    `orm:"column(picture);size(200)" description:"管理员照片"`
	WorkRecord         string    `orm:"column(work_record);size(800)" description:"从业经历"`
	Speciality         string    `orm:"column(speciality);size(800)" description:"擅长领域"`
	Education          int8      `orm:"column(education)" description:"学历 1=无、2=高中、3=大专、4=本科、5=硕士、6=博士"`
	Certificate        string    `orm:"column(certificate);size(800)" description:"证书"`
	Intro              string    `orm:"column(intro);size(2000)" description:"个人简介"`
	JobMarketAddress   uint8     `orm:"column(job_market_address)" description:"职场所在地 0:无 1:北京 2:深圳 3:上海"`
	AscriptionCategory uint8     `orm:"column(ascription_category);null" description:"归属分类 0:无 1:财富管理 2:内部理财师"`
	DateEntered        time.Time `orm:"column(date_entered);type(datetime);auto_now_add" description:"创建时间"`
	DateModified       time.Time `orm:"column(date_modified);type(datetime);auto_now_add" description:"修改时间"`
	Deleted            int8      `orm:"column(deleted)" description:"是否删除 0否 1是"`
}

func (t *FortuneUserTools) TableName() string {
	return "fortune_user_tools"
}

func init() {
	orm.RegisterModel(new(FortuneUserTools))
}

// AddFortuneUserTools insert a new FortuneUserTools into database and returns
// last inserted Id on success.
func AddFortuneUserTools(m *FortuneUserTools) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneUserToolsById retrieves FortuneUserTools by Id. Returns error if
// Id doesn't exist
func GetFortuneUserToolsById(id int) (v *FortuneUserTools, err error) {
	o := orm.NewOrm()
	v = &FortuneUserTools{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneUserTools retrieves all FortuneUserTools matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneUserTools(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneUserTools))
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

	var l []FortuneUserTools
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

// UpdateFortuneUserTools updates FortuneUserTools by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneUserToolsById(m *FortuneUserTools) (err error) {
	o := orm.NewOrm()
	v := FortuneUserTools{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneUserTools deletes FortuneUserTools by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneUserTools(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneUserTools{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneUserTools{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
