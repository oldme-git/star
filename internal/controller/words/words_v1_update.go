package words

import (
    "context"

    "star/api/words/v1"
    "star/internal/logic/words"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
    err = c.words.Update(ctx, req.Id, words.UpdateInput{
        Word:               req.Word,
        Definition:         req.Definition,
        ExampleSentence:    req.ExampleSentence,
        ChineseTranslation: req.ChineseTranslation,
        Pronunciation:      req.Pronunciation,
        ProficiencyLevel:   req.ProficiencyLevel,
    })
    return nil, err
}
