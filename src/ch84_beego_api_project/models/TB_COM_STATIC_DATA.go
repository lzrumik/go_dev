package models

type TBCOMSTATICDATA struct {
	CODETYPE       string `orm:"column(CODE_TYPE);size(255)" description:"编码类型"`
	CODENAME       string `orm:"column(CODE_NAME);size(255)"`
	CODEVALUE      string `orm:"column(CODE_VALUE)" description:"编码值"`
	CODEDESC       string `orm:"column(CODE_DESC)" description:"编码描述"`
	CODETYPEALIAS  string `orm:"column(CODE_TYPE_ALIAS);size(255);null" description:"编码别名,用作java的参数名称"`
	SORTID         int    `orm:"column(SORT_ID)" description:"排序ID"`
	STATE          int    `orm:"column(STATE)" description:"状态(1在用，2废弃"`
	EXTERNCODETYPE string `orm:"column(EXTERN_CODE_TYPE);size(255);null" description:"外部编码"`
	ENTID          int    `orm:"column(ENT_ID);null" description:"实体ID"`
}
