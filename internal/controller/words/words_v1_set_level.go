package words

import (
	"context"

	"star/api/words/v1"
	"star/internal/logic/words"
)

func (c *ControllerV1) SetLevel(ctx context.Context, req *v1.SetLevelReq) (res *v1.SetLevelRes, err error) {
	err = words.SetLevel(ctx, req.Id, req.Level)
	return nil, err
}
