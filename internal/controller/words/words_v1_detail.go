package words

import (
	"context"

	"star/internal/logic/words"

	"star/api/words/v1"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	word, err := words.Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DetailRes{
		Words: word,
	}, nil
}
