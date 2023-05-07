package nordigen

import (
	"context"
	"strconv"
)

// CreateRequisition creates a new requisition
func (c Client) CreateRequisition(ctx context.Context, token string, requisitionRequestBody *RequisitionRequestBody) (*Requisition, error) {
	var requisition Requisition
	err := c.HTTP.Post(ctx, RequisitionsPath+"/", RequestHeadersWithAuth(token), requisitionRequestBody, &requisition)
	if err != nil {
		return nil, err
	}

	return &requisition, nil
}

// ListRequisition lists all requisitions
func (c Client) ListRequisitions(ctx context.Context, token string, requestParams *ListRequisitionsParams) (*Requisitions, error) {
	var requisitions Requisitions

	endpointURL := RequisitionsPath + "/"
	if requestParams != nil {
		if requestParams.Limit != 0 {
			endpointURL = endpointURL + "?" + strconv.Itoa(requestParams.Limit)
		}
		if requestParams.Offset != 0 {
			endpointURL = endpointURL + "&" + strconv.Itoa(requestParams.Offset)
		}
	}

	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(token), &requisitions)
	if err != nil {
		return nil, err
	}

	return &requisitions, nil
}

// FetchRequisition retrieves a requisition by requisitionID
func (c Client) FetchRequisition(ctx context.Context, token string, requisitionID string) (*Requisition, error) {
	var requisition Requisition
	err := c.HTTP.Get(ctx, RequisitionsPath+"/"+requisitionID, RequestHeadersWithAuth(token), &requisition)
	if err != nil {
		return nil, err
	}

	return &requisition, nil
}

// DeleteRequisition deletes a requisition by requisitionID
func (c Client) DeleteRequisition(ctx context.Context, token string, requisitionID string) error {
	err := c.HTTP.Delete(ctx, RequisitionsPath+"/"+requisitionID, RequestHeadersWithAuth(token), nil)
	if err != nil {
		return err
	}

	return nil
}
