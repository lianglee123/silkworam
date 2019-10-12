package model

import (
	"time"
	"hexcloud.cn/histore/oauth/utils"
)

type MyGormModel struct {
	Id        uint64     `json:"id,string" gorm:"type:int8;primary_key;"`
	CreatedAt time.Time  `json:"created_at" gorm:"index;type:timestamptz;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"index;type:timestamptz;"`
	DeletedAt *time.Time `json:"-" gorm:"index;type:timestamptz;"`
	CreatedBy uint64     `json:"created_by,string" grom:"type:int8;"`
	UpdatedBy uint64     `json:"updated_by,string"`
}

func (m *MyGormModel) BeforeCreate() error {
	if m.Id == 0 {
		m.Id, _ = utils.GetId()
	}
	return nil
}
