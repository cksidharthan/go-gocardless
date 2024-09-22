package gocardless

import (
	"context"
	"strconv"
)

// CreateRequisition creates a new requisition
func (c Client) CreateRequisition(ctx context.Context, requisitionRequestBody *RequisitionRequestBody) (*Requisition, error) {
	var requisition Requisition
	err := c.HTTP.Post(ctx, RequisitionsPath, RequestHeadersWithAuth(c.Token.Access), requisitionRequestBody, &requisition)
	if err != nil {
		return nil, err
	}

	return &requisition, nil
}

// ListRequisition lists all requisitions
func (c Client) ListRequisitions(ctx context.Context, requestParams *ListRequisitionsParams) (*Requisitions, error) {
	var requisitions Requisitions

	endpointURL := RequisitionsPath
	if requestParams != nil {
		if requestParams.Limit != 0 {
			endpointURL = endpointURL + "?" + strconv.Itoa(requestParams.Limit)
		}
		if requestParams.Offset != 0 {
			endpointURL = endpointURL + "&" + strconv.Itoa(requestParams.Offset)
		}
	}

	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(c.Token.Access), &requisitions)
	if err != nil {
		return nil, err
	}

	return &requisitions, nil
}

// FetchRequisition retrieves a requisition by requisitionID
func (c Client) FetchRequisition(ctx context.Context, requisitionID string) (*Requisition, error) {
	var requisition Requisition
	err := c.HTTP.Get(ctx, RequisitionsPath+requisitionID, RequestHeadersWithAuth(c.Token.Access), &requisition)
	if err != nil {
		return nil, err
	}

	return &requisition, nil
}

// DeleteRequisition deletes a requisition by requisitionID
func (c Client) DeleteRequisition(ctx context.Context, requisitionID string) error {
	err := c.HTTP.Delete(ctx, RequisitionsPath+requisitionID, RequestHeadersWithAuth(c.Token.Access), nil)
	if err != nil {
		return err
	}

	return nil
}
