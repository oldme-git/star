package words

import (
	"context"

	"star/internal/logic/words"

	"star/api/words/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = words.Delete(ctx, req.Id)
	return
}
