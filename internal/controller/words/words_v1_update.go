package words

import (
	"context"

	"star/internal/logic/words"
	"star/internal/model"

	"star/api/words/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = words.Update(ctx, req.Id, &model.WordInput{
		Word:               req.Word,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ChineseTranslation: req.ChineseTranslation,
		Pronunciation:      req.Pronunciation,
		ProficiencyLevel:   model.ProficiencyLevel(req.ProficiencyLevel),
	})
	return nil, err
}
