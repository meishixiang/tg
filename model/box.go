package model

import (
	"boxDemo/drivers"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB = drivers.MysqlDb

type Box struct {
	BoxId      int `gorm:"primary_key;column:boxId"`
	Size       int
	Color      int
	Status     uint
	IsDeleted  uint `gorm:"column:isDeleted"`
	CreateId   uint `gorm:"column:createId"`
	CreateTime uint `gorm:"column:createTime"`
	UpdateTime uint `gorm:"column:updateTime"`
}

func (Box) TableName() string {
	return "box"
}

func (model *Box) GetBoxList(page int, pageSize int) ( []*Box, int64, error) {
	var total int64 = 0
	var boxList []*Box
	db.Find(&boxList).Count(&total)
	if err := db.Limit(page).Offset((page - 1) * pageSize).Find(&boxList).Error; err != nil {
		return nil, total, err
	}

	return boxList, total, nil
}
