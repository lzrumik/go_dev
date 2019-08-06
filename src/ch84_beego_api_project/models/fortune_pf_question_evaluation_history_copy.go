package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortunePfQuestionEvaluationHistoryCopy struct {
	Id                          int       `orm:"column(id);pk" description:"测评ID"`
	UserId                      string    `orm:"column(user_id);size(36)" description:"用户ID"`
	Username                    string    `orm:"column(username);size(30)" description:"用户名"`
	FullName                    string    `orm:"column(full_name);size(20)" description:"用户姓名"`
	MobilePrefix                string    `orm:"column(mobile_prefix);size(10)" description:"手机号前缀 国家码"`
	Mobile                      string    `orm:"column(mobile);size(15)" description:"手机号"`
	Type                        uint8     `orm:"column(type)" description:"测评类型 1=个人 2=机构"`
	EvaluationGrade             uint16    `orm:"column(evaluation_grade)" description:"测评分数"`
	EvaluationLevel             string    `orm:"column(evaluation_level);size(10)" description:"测评等级"`
	EvaluationQuestionSnapshoot string    `orm:"column(evaluation_question_snapshoot)" description:"测评题选项列表快照"`
	SubscribeQuestionSnapshoot  string    `orm:"column(subscribe_question_snapshoot)" description:"预约题选项列表快照"`
	EvaluationResult            string    `orm:"column(evaluation_result);size(1000)" description:"测评题结果"`
	EvaluationExtResult         string    `orm:"column(evaluation_ext_result);size(500)" description:"预约题结果"`
	Remark                      string    `orm:"column(remark);size(100)" description:"备注"`
	DateEntered                 time.Time `orm:"column(date_entered);type(datetime)" description:"创建时间"`
	Deleted                     uint8     `orm:"column(deleted)" description:"是否删除"`
}

func (t *FortunePfQuestionEvaluationHistoryCopy) TableName() string {
	return "fortune_pf_question_evaluation_history_copy"
}

func init() {
	orm.RegisterModel(new(FortunePfQuestionEvaluationHistoryCopy))
}

// AddFortunePfQuestionEvaluationHistoryCopy insert a new FortunePfQuestionEvaluationHistoryCopy into database and returns
// last inserted Id on success.
func AddFortunePfQuestionEvaluationHistoryCopy(m *FortunePfQuestionEvaluationHistoryCopy) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortunePfQuestionEvaluationHistoryCopyById retrieves FortunePfQuestionEvaluationHistoryCopy by Id. Returns error if
// Id doesn't exist
func GetFortunePfQuestionEvaluationHistoryCopyById(id int) (v *FortunePfQuestionEvaluationHistoryCopy, err error) {
	o := orm.NewOrm()
	v = &FortunePfQuestionEvaluationHistoryCopy{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortunePfQuestionEvaluationHistoryCopy retrieves all FortunePfQuestionEvaluationHistoryCopy matches certain condition. Returns empty list if
// no records exist
func GetAllFortunePfQuestionEvaluationHistoryCopy(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortunePfQuestionEvaluationHistoryCopy))
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

	var l []FortunePfQuestionEvaluationHistoryCopy
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

// UpdateFortunePfQuestionEvaluationHistoryCopy updates FortunePfQuestionEvaluationHistoryCopy by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortunePfQuestionEvaluationHistoryCopyById(m *FortunePfQuestionEvaluationHistoryCopy) (err error) {
	o := orm.NewOrm()
	v := FortunePfQuestionEvaluationHistoryCopy{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortunePfQuestionEvaluationHistoryCopy deletes FortunePfQuestionEvaluationHistoryCopy by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortunePfQuestionEvaluationHistoryCopy(id int) (err error) {
	o := orm.NewOrm()
	v := FortunePfQuestionEvaluationHistoryCopy{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortunePfQuestionEvaluationHistoryCopy{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
