package engine

//解析结果
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//请求 结构体
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

//空parser
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
