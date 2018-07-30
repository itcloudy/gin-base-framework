package models

import "time"

//BaseModel base model
type BaseModel struct {
	ID        int       `gorm:"ForeignKey:ID" gorm:"primary_key" form:"id" json:"id"` //主键ID
	CreatedAt time.Time `json:"created_at" form:"-"`                                  //记录创建时间
	UpdatedAt time.Time `json:"updated_at" form:"-"`                                  //记录更新时间
}
