package requisitions

import (
	"context"

	"github.com/weportfolio/go-nordigen/consts"
)

// Post creates a new requisition
func (c Client) Post(ctx context.Context, token string, requisitionRequestBody *RequisitionRequestBody) (*Requisition, error) {
	var requisition Requisition
	err := c.HTTP.Post(ctx, consts.RequisitionsPath+"/", consts.RequestHeadersWithAuth(token), requisitionRequestBody, &requisition)
	if err != nil {
		return nil, err
	}

	return &requisition, nil
}
