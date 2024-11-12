package words

import (
	"context"

	"star/internal/logic/users"
	"star/internal/logic/words"
	"star/internal/model"

	"star/api/words/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	uid, err := users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	err = words.Create(ctx, &model.WordInput{
		Uid:                uid,
		Word:               req.Word,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ChineseTranslation: req.ChineseTranslation,
		Pronunciation:      req.Pronunciation,
		ProficiencyLevel:   model.ProficiencyLevel(req.ProficiencyLevel),
	})
	return nil, err
}
