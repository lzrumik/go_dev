package models

import "time"

type FortuneRoleUser struct {
	RoleId       string    `orm:"column(role_id);size(36)" description:"角色ID"`
	UserId       string    `orm:"column(user_id);size(36);null" description:"用户ID"`
	DateEntered  time.Time `orm:"column(date_entered);type(datetime);null" description:"创建时间"`
	DateModified time.Time `orm:"column(date_modified);type(datetime);null" description:"修改时间"`
	Deleted      int8      `orm:"column(deleted)" description:"是否删除"`
}
