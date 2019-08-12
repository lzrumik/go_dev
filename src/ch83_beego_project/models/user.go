package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

var (
	db orm.Ormer
)

//由于model这个名字叫 UserInfo 那么操作的表其实 user_info
type UserInfo struct{
	Id int64
	Username string
	Password string
}

/**

CREATE TABLE `user_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(200) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(200) NOT NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

INSERT INTO `user_info` (`username`, `password`) VALUES ("ceshi123", "passwd1234");
INSERT INTO `user_info` (`username`, `password`) VALUES ("ceshi124", "passwd1234");
INSERT INTO `user_info` (`username`, `password`) VALUES ("ceshi125", "passwd1234");
INSERT INTO `user_info` (`username`, `password`) VALUES ("ceshi126", "passwd1234");
INSERT INTO `user_info` (`username`, `password`) VALUES ("ceshi127", "passwd1234");
INSERT INTO `user_info` (`username`, `password`) VALUES ("ceshi128", "passwd1234");
INSERT INTO `user_info` (`username`, `password`) VALUES ("ceshi129", "passwd1234");

 */
func init() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	orm.RegisterDataBase("default", "mysql", "beta-mysql:nopass.2@tcp(10.255.72.27:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))
	db = orm.NewOrm()
}

func AddUser(user_info *UserInfo)(int64,error){
	id,err := db.Insert(user_info)
	return id,err
}

func ReadUserInfo(users *[]UserInfo){
	qb, _:=orm.NewQueryBuilder("mysql")

	qb.Select("*").From("user_info")

	sql := qb.String()
	db.Raw(sql).QueryRows(users)
}

