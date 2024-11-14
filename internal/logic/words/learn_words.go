package words

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"star/internal/dao"
	"star/internal/model"
	"star/internal/model/entity"
)

// Rand 随机若干获取单词
func Rand(ctx context.Context, uid, limit uint) ([]entity.Words, error) {
	if limit <= 0 {
		limit = 50
	}
	var (
		list = make([]entity.Words, limit)
		err  error
	)
	db := dao.Words.Ctx(ctx)
	if uid > 0 {
		db = db.Where("uid", uid)
	}

	err = db.Limit(int(limit)).OrderRandom().Scan(&list)
	return list, err
}

// SetLevel 设置单词熟练度
func SetLevel(ctx context.Context, uid, id uint, level model.ProficiencyLevel) error {
	if level < 0 || level > 5 {
		return gerror.New("熟练度值不合法")
	}

	db := dao.Words.Ctx(ctx)
	if uid > 0 {
		db = db.Where("uid", uid)
	}

	_, err := db.Data("proficiency_level", level).Where("id", id).Update()
	return err
}
