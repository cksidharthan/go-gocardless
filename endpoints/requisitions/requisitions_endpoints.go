package requisitions

import (
	"context"
	"strconv"

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

func (c Client) List(ctx context.Context, token string, requestParams *ListRequestParams) (*Requisitions, error) {
	var requisitions Requisitions

	endpointURL := consts.RequisitionsPath + "/"
	if requestParams != nil {
		if requestParams.Limit != 0 {
			endpointURL = endpointURL + "?" + strconv.Itoa(requestParams.Limit)
		}
		if requestParams.Offset != 0 {
			endpointURL = endpointURL + "&" + strconv.Itoa(requestParams.Offset)
		}
	}

	err := c.HTTP.Get(ctx, endpointURL, consts.RequestHeadersWithAuth(token), &requisitions)
	if err != nil {
		return nil, err
	}

	return &requisitions, nil
}

// Fetch retrieves a requisition by requisitionID
func (c Client) Fetch(ctx context.Context, token string, requisitionID string) (*Requisition, error) {
	var requisition Requisition
	err := c.HTTP.Get(ctx, consts.RequisitionsPath+"/"+requisitionID, consts.RequestHeadersWithAuth(token), &requisition)
	if err != nil {
		return nil, err
	}

	return &requisition, nil
}

// Delete deletes a requisition by requisitionID
func (c Client) Delete(ctx context.Context, token string, requisitionID string) error {
	err := c.HTTP.Delete(ctx, consts.RequisitionsPath+"/"+requisitionID, consts.RequestHeadersWithAuth(token), nil)
	if err != nil {
		return err
	}

	return nil
}
