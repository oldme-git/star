package words

import (
	"context"

	"star/api/words/v1"
	"star/internal/logic/users"
	"star/internal/logic/words"
)

func (c *ControllerV1) SetLevel(ctx context.Context, req *v1.SetLevelReq) (res *v1.SetLevelRes, err error) {
	uid, err := users.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	err = words.SetLevel(ctx, uid, req.Id, req.Level)
	return nil, err
}
