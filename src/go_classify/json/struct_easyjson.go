// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package json

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9f2eff5fDecodeGoClassifyRepay(in *jlexer.Lexer, out *RepayInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "invest_repay_tm":
			out.InvestRepayTm = string(in.String())
		case "invest_repay_amount":
			out.InvestRepayAmount = float64(in.Float64())
		case "loan_id":
			out.LoanID = string(in.String())
		case "invest_repay_status":
			out.InvestRepayStatus = string(in.String())
		case "invest_repay_due_date":
			out.InvestRepayDueDate = string(in.String())
		case "plan_repay_amount":
			out.PlanRepayAmount = float64(in.Float64())
		case "ca_cd":
			out.CaCd = int(in.Int())
		case "uid":
			out.UID = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGoClassifyRepay(out *jwriter.Writer, in RepayInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"invest_repay_tm\":"
		out.RawString(prefix[1:])
		out.String(string(in.InvestRepayTm))
	}
	{
		const prefix string = ",\"invest_repay_amount\":"
		out.RawString(prefix)
		out.Float64(float64(in.InvestRepayAmount))
	}
	{
		const prefix string = ",\"loan_id\":"
		out.RawString(prefix)
		out.String(string(in.LoanID))
	}
	{
		const prefix string = ",\"invest_repay_status\":"
		out.RawString(prefix)
		out.String(string(in.InvestRepayStatus))
	}
	{
		const prefix string = ",\"invest_repay_due_date\":"
		out.RawString(prefix)
		out.String(string(in.InvestRepayDueDate))
	}
	{
		const prefix string = ",\"plan_repay_amount\":"
		out.RawString(prefix)
		out.Float64(float64(in.PlanRepayAmount))
	}
	{
		const prefix string = ",\"ca_cd\":"
		out.RawString(prefix)
		out.Int(int(in.CaCd))
	}
	{
		const prefix string = ",\"uid\":"
		out.RawString(prefix)
		out.Int(int(in.UID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RepayInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGoClassifyRepay(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RepayInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGoClassifyRepay(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RepayInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGoClassifyRepay(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RepayInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGoClassifyRepay(l, v)
}