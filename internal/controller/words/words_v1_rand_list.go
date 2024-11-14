package words

import (
	"context"

	"star/api/words/v1"
	"star/internal/logic/words"
	"star/internal/model"
)

func (c *ControllerV1) RandList(ctx context.Context, req *v1.RandListReq) (res *v1.RandListRes, err error) {
	wordList, err := words.Rand(ctx, req.Limit)
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

	return &v1.RandListRes{
		List: list,
	}, nil
}
