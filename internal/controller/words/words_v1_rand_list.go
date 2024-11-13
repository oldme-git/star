package words

import (
	"context"

	"star/api/words/v1"
	"star/internal/logic/words"
)

func (c *ControllerV1) RandList(ctx context.Context, req *v1.RandListReq) (res *v1.RandListRes, err error) {
	list, err := words.Rand(ctx, req.Limit)
	if err != nil {
		return nil, err
	}
	return &v1.RandListRes{
		List: list,
	}, nil
}
