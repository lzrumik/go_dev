package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FortuneLoanBak struct {
	Id                          int       `orm:"column(id);pk"`
	Company                     string    `orm:"column(company);size(20)" description:"所属公司"`
	LoanId                      string    `orm:"column(loan_id);size(70)" description:"合同编号"`
	FrontendDisplayName         string    `orm:"column(frontend_display_name);size(50)" description:"前端展示名称"`
	Name                        string    `orm:"column(name);size(50)" description:"项目名称"`
	Type                        int8      `orm:"column(type)" description:"项目类型:0人民币,1美元,2私募股权,3海外置业,4家庭保障、财富传承"`
	Currency                    int8      `orm:"column(currency)" description:"币种，0人民币，1美元"`
	Deadline                    int16     `orm:"column(deadline)" description:"产品期限"`
	DeadlineType                int8      `orm:"column(deadline_type)" description:"产品期限类型:0天,1月,2年"`
	DeadlineRemark              string    `orm:"column(deadline_remark);size(300)" description:"产品期限备注"`
	DeadlineDay                 int       `orm:"column(deadline_day)" description:"产品期限用天标识"`
	Profit                      string    `orm:"column(profit);size(30)" description:"预期收益"`
	StartTime                   time.Time `orm:"column(start_time);type(date)" description:"起息日"`
	EndTime                     time.Time `orm:"column(end_time);type(date)" description:"到期日"`
	Collect                     string    `orm:"column(collect);size(50)" description:"募集账户"`
	BankName                    string    `orm:"column(bank_name);size(100)" description:"收款银行"`
	Account                     string    `orm:"column(account);size(50)" description:"募集账号"`
	MinMoney                    float64   `orm:"column(min_money);digits(15);decimals(2)" description:"起投金额"`
	MaxInvestAmount             float64   `orm:"column(max_invest_amount);digits(15);decimals(2)" description:"最大认购金额 0:不限制"`
	SubFee                      string    `orm:"column(sub_fee);size(500)" description:"认购费"`
	ManageFee                   string    `orm:"column(manage_fee);size(1000)" description:"管理费"`
	RmDiscount                  float64   `orm:"column(rm_discount);digits(8);decimals(2)" description:"RM佣金系数"`
	RedemptionFee               string    `orm:"column(redemption_fee);size(300)" description:"赎回费"`
	IncrMoney                   float64   `orm:"column(incr_money);digits(15);decimals(2)" description:"递增金额"`
	CollectMoney                float64   `orm:"column(collect_money);digits(15);decimals(2)" description:"募集总金额"`
	EstimatCollectMoney         float64   `orm:"column(estimat_collect_money);digits(15);decimals(2)" description:"预计募集总额"`
	TotalMoney                  float64   `orm:"column(total_money);digits(15);decimals(2)" description:"募集到的总金额"`
	FrozenMoney                 float64   `orm:"column(frozen_money);digits(15);decimals(2)" description:"冻结金额"`
	AchievementDiscount         float64   `orm:"column(achievement_discount);digits(15);decimals(2)" description:"业绩折扣"`
	ServiceDiscount             float64   `orm:"column(service_discount);digits(15);decimals(2)" description:"服务费"`
	CommissionDiscount          float64   `orm:"column(commission_discount);digits(15);decimals(2)" description:"佣金这算系数"`
	CityManagerRate             float64   `orm:"column(city_manager_rate);digits(15);decimals(2)" description:"城市负责人佣金系数"`
	CountrywideManagerRate      float64   `orm:"column(countrywide_manager_rate);digits(15);decimals(2)" description:"全国负责人佣金系数"`
	FxServiceManagerRate        float64   `orm:"column(fx_service_manager_rate);digits(15);decimals(2)" description:"凤新-增值业务经理佣金系数"`
	FxProductManagerRate        float64   `orm:"column(fx_product_manager_rate);digits(15);decimals(2)" description:"凤新-产品经理佣金系数"`
	FxPropertyConsultantRate    float64   `orm:"column(fx_property_consultant_rate);digits(15);decimals(2)" description:"凤新-置业顾问佣金系数"`
	CollectOrg                  string    `orm:"column(collect_org);size(50)" description:"募集机构"`
	ProfitType                  string    `orm:"column(profit_type);size(200)" description:"收益方式:0一次性还本息,1半年付息，到期还本2按季付息，到期还本3产品结束时一次性分配4一次性分配收益 非数字手动填写"`
	Guarantee                   string    `orm:"column(guarantee);size(300)" description:"投资保障"`
	ProfitDesc                  string    `orm:"column(profit_desc);size(300)" description:"收益率说明"`
	FundManager                 string    `orm:"column(fund_manager);size(30)" description:"基金管理人"`
	Publisher                   string    `orm:"column(publisher);size(100)" description:"发行人"`
	RepaymentSource             string    `orm:"column(repayment_source);size(50)" description:"还款来源"`
	Level                       string    `orm:"column(level);size(10)" description:"级别"`
	TrusteeshipOrg              string    `orm:"column(trusteeship_org);size(50)" description:"托管机构"`
	CostSettlementOrg           string    `orm:"column(cost_settlement_org);size(100)" description:"费用结算机构"`
	CompanyId                   string    `orm:"column(company_id);size(36)" description:"归属企业ID"`
	ContractNoPrefix            string    `orm:"column(contract_no_prefix);size(50)" description:"合同编号前缀"`
	ProductSummary              string    `orm:"column(product_summary);size(20)" description:"产品亮点概述"`
	Comment                     string    `orm:"column(comment);size(500)" description:"备注"`
	InterestBearingRule         uint8     `orm:"column(interest_bearing_rule)" description:"计息规则 1:按天计息 2:按月平均"`
	InterestBearingPeriod       uint8     `orm:"column(interest_bearing_period)" description:"计息周期 1:365天 2:360天"`
	Guarantor                   string    `orm:"column(guarantor);size(100)" description:"担保方"`
	Status                      int8      `orm:"column(status)" description:"状态,0预热,1在售,2已售完,3已到期"`
	Version                     int64     `orm:"column(version)" description:"版本号"`
	IsFrontendDisplay           uint8     `orm:"column(is_frontend_display)" description:"是否前端展示 0:是 1:否"`
	IsSupportElectronicContract uint8     `orm:"column(is_support_electronic_contract)" description:"是否支持电子合同 0:否 1:是"`
	ServiceCostRule             uint      `orm:"column(service_cost_rule)" description:"服务费结算规则"`
	CreatedUserId               string    `orm:"column(created_user_id);size(30)" description:"创建人"`
	DateEntered                 time.Time `orm:"column(date_entered);type(datetime)" description:"添加时间"`
	ModifiedUserId              string    `orm:"column(modified_user_id);size(30)" description:"更新操作人"`
	DateModified                time.Time `orm:"column(date_modified);type(datetime)" description:"更新时间"`
	Deleted                     int8      `orm:"column(deleted)" description:"删除 0 未删除,1删除"`
	VideoRecord                 int8      `orm:"column(video_record)" description:"是否双录，1是，0否"`
	OnlineOrder                 int8      `orm:"column(online_order)" description:"是否线上认购，1是，0否"`
}

func (t *FortuneLoanBak) TableName() string {
	return "fortune_loan_bak"
}

func init() {
	orm.RegisterModel(new(FortuneLoanBak))
}

// AddFortuneLoanBak insert a new FortuneLoanBak into database and returns
// last inserted Id on success.
func AddFortuneLoanBak(m *FortuneLoanBak) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFortuneLoanBakById retrieves FortuneLoanBak by Id. Returns error if
// Id doesn't exist
func GetFortuneLoanBakById(id int) (v *FortuneLoanBak, err error) {
	o := orm.NewOrm()
	v = &FortuneLoanBak{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFortuneLoanBak retrieves all FortuneLoanBak matches certain condition. Returns empty list if
// no records exist
func GetAllFortuneLoanBak(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FortuneLoanBak))
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

	var l []FortuneLoanBak
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

// UpdateFortuneLoanBak updates FortuneLoanBak by Id and returns error if
// the record to be updated doesn't exist
func UpdateFortuneLoanBakById(m *FortuneLoanBak) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanBak{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFortuneLoanBak deletes FortuneLoanBak by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFortuneLoanBak(id int) (err error) {
	o := orm.NewOrm()
	v := FortuneLoanBak{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FortuneLoanBak{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
