package _9json_benchmark


type RepayInfo struct {
	InvestRepayTm      string  `json:"invest_repay_tm"`
	InvestRepayAmount  float64 `json:"invest_repay_amount"`
	LoanID             string  `json:"loan_id"`
	InvestRepayStatus  string  `json:"invest_repay_status"`
	InvestRepayDueDate string  `json:"invest_repay_due_date"`
	PlanRepayAmount    float64 `json:"plan_repay_amount"`
	CaCd               int     `json:"ca_cd"`
	UID                int     `json:"uid"`
}


