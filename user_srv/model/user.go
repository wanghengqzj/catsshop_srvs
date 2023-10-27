package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}

/*
1.密码：密文保存，不可反解
  对称加密
  非对称加密
  md5 信息摘要算法
*/

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11) comment '电话号码';not null"`
	Password string     `gorm:"type:varchar(100) comment '密码'"`
	NickName string     `gorm:"type:varchar(20) comment '昵称'"` //昵称
	Birthday *time.Time `gorm:"type:datetime comment '生日'"`    //设置为指针类型防止保存的时候出错
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女, male表示男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示普通用户, 2表示管理员'"`
}
