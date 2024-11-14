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
	wordList, total, err := words.List(ctx, query)
	if err != nil {
		return nil, err
	}

	var list []v1.List
	for _, v := range wordList {
		list = append(list, v1.List{
			Id:               v.Id,
			Word:             v.Word,
			Definition:       v.Definition,
			ProficiencyLevel: model.ProficiencyLevel(v.ProficiencyLevel),
		})
	}

	return &v1.ListRes{
		List:  list,
		Total: total,
	}, nil
}
