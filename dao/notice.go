package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNoticeById 数据库中查找uId对应的 notice record
func (dao *NoticeDao) GetNoticeById(uId uint) (*model.Notice, error) {
	var notice *model.Notice
	err := dao.DB.Model(&model.Notice{}).Where("id = ?", uId).First(&notice).Error
	return notice, err
}
