package words

import (
	"context"

	v1 "star/api/words/v1"
	"star/internal/dao"
	"star/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (w *Words) Rand(ctx context.Context, uid, limit uint) ([]entity.Words, error) {
	if limit <= 0 {
		limit = 50
	}
	var (
		err  error
		cls  = dao.Words.Columns()
		orm  = dao.Words.Ctx(ctx)
		list = make([]entity.Words, limit)
	)
	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}
	err = orm.Limit(int(limit)).OrderRandom().Scan(&list)
	return list, err
}

func (w *Words) SetLevel(ctx context.Context, uid, id uint, level v1.ProficiencyLevel) error {
	if level < 0 || level > 5 {
		return gerror.New("熟练度值不合法")
	}

	var (
		cls = dao.Words.Columns()
		orm = dao.Words.Ctx(ctx)
	)
	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}

	_, err := orm.Data(cls.ProficiencyLevel, level).Where(cls.Id, id).Update()
	return err
}
