package words

import (
	"context"

	"star/api/words/v1"
	"star/internal/logic/words"
	"star/internal/model"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	query := &model.WordQuery{
		Word: req.Word,
		Page: req.Page,
		Size: req.Size,
	}
	list, total, err := words.List(ctx, query)
	if err != nil {
		return nil, err
	}
	return &v1.ListRes{
		List:  list,
		Total: total,
	}, nil
}
