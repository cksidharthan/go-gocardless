package agreements

import (
	"context"
	"strconv"

	"github.com/weportfolio/go-nordigen/consts"
)

// Post creates a new agreement for the enduser
func (c Client) Post(ctx context.Context, token string, agreementRequestBody AgreementRequestBody) (*Agreement, error) {
	var agreement Agreement
	err := c.HTTP.Post(ctx, consts.AgreementsEndusersPath, consts.RequestHeadersWithAuth(token), agreementRequestBody, &agreement)
	if err != nil {
		return nil, err
	}

	return &agreement, nil
}

// Fetch retrieves an agreement for the enduser by agreementID
func (c Client) Fetch(ctx context.Context, token string, agreementID string) (*Agreement, error) {
	var agreement Agreement
	err := c.HTTP.Get(ctx, consts.AgreementsEndusersPath+agreementID, consts.RequestHeadersWithAuth(token), &agreement)
	if err != nil {
		return nil, err
	}

	return &agreement, nil
}

// List returns a list of agreements for the enduser
func (c Client) List(ctx context.Context, token string, requestParams *ListRequestParams) (*Agreements, error) {
	var agreements Agreements

	endpointURL := consts.AgreementsEndusersPath
	if requestParams != nil {
		endpointURL = endpointURL + "?" + strconv.Itoa(requestParams.Limit)
	}
	if requestParams != nil {
		endpointURL = endpointURL + "&" + strconv.Itoa(requestParams.Offset)
	}

	err := c.HTTP.Get(ctx, endpointURL, consts.RequestHeadersWithAuth(token), &agreements)
	if err != nil {
		return nil, err
	}

	return &agreements, nil
}

// Delete deletes an agreement for the enduser by agreementID
func (c Client) Delete(ctx context.Context, token string, agreementID string) error {
	err := c.HTTP.Delete(ctx, consts.AgreementsEndusersPath+agreementID, consts.RequestHeadersWithAuth(token), nil)
	if err != nil {
		return err
	}

	return nil
}

// Update updates an agreement for the enduser by agreementID
func (c Client) Update(ctx context.Context, token string, agreementID string, updateRequestBody UpdateRequestBody) (*Agreement, error) {
	var agreement Agreement
	// TODO: Check if this is the correct way to update an agreement, The API doc wants to append a /accept to the end of the URL
	err := c.HTTP.Put(ctx, consts.AgreementsEndusersPath+agreementID, consts.RequestHeadersWithAuth(token), updateRequestBody, &agreement)
	if err != nil {
		return nil, err
	}

	return &agreement, nil
}
