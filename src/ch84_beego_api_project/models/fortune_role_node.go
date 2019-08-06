package models

import "time"

type FortuneRoleNode struct {
	RoleId      string    `orm:"column(role_id);size(36)" description:"角色ID"`
	NodeId      string    `orm:"column(node_id);size(36);null" description:"节点ID"`
	DateEntered time.Time `orm:"column(date_entered);type(datetime);null" description:"创建时间"`
}
