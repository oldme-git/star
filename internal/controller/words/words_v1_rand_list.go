package words

import (
    "context"

    "star/api/words/v1"
)

func (c *ControllerV1) RandList(ctx context.Context, req *v1.RandListReq) (res *v1.RandListRes, err error) {
    uid, err := c.users.GetUid(ctx)
    if err != nil {
        return nil, err
    }

    wordList, err := c.words.Rand(ctx, uid, req.Limit)
    if err != nil {
        return nil, err
    }

    var list []v1.List
    for _, v := range wordList {
        list = append(list, v1.List{
            Id:               v.Id,
            Word:             v.Word,
            Definition:       v.Definition,
            ProficiencyLevel: v1.ProficiencyLevel(v.ProficiencyLevel),
        })
    }

    return &v1.RandListRes{
        List: list,
    }, nil
}
