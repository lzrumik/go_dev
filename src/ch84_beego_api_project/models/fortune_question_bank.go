package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneQuestionBank struct {
	Id                  int       `orm:"column(id);pk" description:"题库ID"`
	EvaluateType        uint8     `orm:"column(evaluate_type)" description:"题类型 1=个人 2=机构"`
	QuestionDescription string    `orm:"column(question_description);size(800)" description:"问题描述"`
	QuestionType        uint8     `orm:"column(question_type)" description:"问题类型 1=测评 2=预约"`
	QuestionMultiSelect uint8     `orm:"column(question_multi_select)" description:"是否多选 0=否 1=是"`
	OptionList          string    `orm:"column(option_list)" description:"题选项列表"`
	DateEntered         time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	CreatedUserId       string    `orm:"column(created_user_id);size(36)" description:"创建人id"`
	CreatedName         string    `orm:"column(created_name);size(20)" description:"创建人姓名"`
	LastDateModified    time.Time `orm:"column(last_date_modified);type(datetime)" description:"最后修改时间"`
	LastModifiedUserId  string    `orm:"column(last_modified_user_id);size(36)" description:"最后修改人id"`
	LastModifiedName    string    `orm:"column(last_modified_name);size(20)" description:"最后修改人姓名"`
	Deleted             uint8     `orm:"column(deleted)" description:"是否删除"`
}

func (t *FortuneQuestionBank) TableName() string {
	return "fortune_question_bank"
}

func init() {
	orm.RegisterModel(new(FortuneQuestionBank))
}

// AddFortuneQuestionBank insert a new FortuneQuestionBank into database and returns
// last inserted Id on success.
func AddFortuneQuestionBank(m *FortuneQuestionBank) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneQuestionBankById retrieves FortuneQuestionBank by Id. Returns error if
// Id doesn't exist
func GetFortuneQuestionBankById(id int) (v *FortuneQuestionBank, err error) {
	o := orm.NewOrm()
	v = &FortuneQuestionBank{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneQuestionBank retrieves all FortuneQuestionBank matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneQuestionBank(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneQuestionBank))
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

	var l []FortuneQuestionBank
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

// UpdateFortuneQuestionBank updates FortuneQuestionBank by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneQuestionBankById(m *FortuneQuestionBank) (err error) {
	o := orm.NewOrm()
	v := FortuneQuestionBank{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneQuestionBank deletes FortuneQuestionBank by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneQuestionBank(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneQuestionBank{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneQuestionBank{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
