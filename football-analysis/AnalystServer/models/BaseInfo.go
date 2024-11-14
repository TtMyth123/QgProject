package models

import "time"

type BaseInfo struct {
	Id        int64
	SysId     int64 `description:"后台用户ID"`
	MainId    int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime); description(创建时间)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime); description(更新时间)"`
}
