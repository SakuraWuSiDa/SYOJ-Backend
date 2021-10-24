package model

import "time"

type User struct {
	ID             int64       `json:"user_id"`
	Username       string      `json:"username" gorm:"type:varchar(64);column:username;NOT NULL"`
	HashedPassword string      `json:"hashed_password" gorm:"type:VARCHAR(255);column:hashed_password;NOT NULL"`
	Email          string      `json:"email" gorm:"type:VARCHAR(64);column:email;NOT NULL"`
	Avatar         string      `json:"avatar" gorm:"type:VARCHAR(128);column:avatar;NOT NULL"`
	Gender         int8        `json:"gender" gorm:"type:tinyint;column:gender;default:2;NOT NULL"`
	StudentID      string      `json:"student_id" gorm:"type:VARCHAR(64);column:student_id;"`
	Class          string      `json:"class" gorm:"type:VARCHAR(64);column:class;"`
	IsAdmin        int8        `json:"is_admin" gorm:"type:tinyint;column:is_admin;default:0;NOT NULL"`
	IsSuperAdmin   int8        `json:"is_super_admin" gorm:"type:tinyint;column:is_super_admin;default:0;NOT NULL"`
	CreatedAt      time.Time   `json:"created_at"  gorm:"column:created_at;NOT NULL"`
	Privilege      uint64      `json:"privilege" gorm:"column:privilege;NOT NULL;default:0"`
}
