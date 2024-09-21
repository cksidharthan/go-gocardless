package gocardless

import (
	"context"
	"strconv"
)

// CreateAgreement creates a new agreement for the enduser
func (c Client) CreateAgreement(ctx context.Context, token string, agreementRequestBody AgreementRequestBody) (*Agreement, error) {
	var agreement Agreement
	err := c.HTTP.Post(ctx, AgreementsEndusersPath, RequestHeadersWithAuth(token), agreementRequestBody, &agreement)
	if err != nil {
		return nil, err
	}

	return &agreement, nil
}

// FetchAgreement retrieves an agreement for the enduser by agreementID
func (c Client) FetchAgreement(ctx context.Context, token string, agreementID string) (*Agreement, error) {
	var agreement Agreement
	err := c.HTTP.Get(ctx, AgreementsEndusersPath+agreementID, RequestHeadersWithAuth(token), &agreement)
	if err != nil {
		return nil, err
	}

	return &agreement, nil
}

// ListAgreements returns a list of agreements for the enduser
func (c Client) ListAgreements(ctx context.Context, token string, requestParams *ListAgreementsParams) (*Agreements, error) {
	var agreements Agreements

	endpointURL := AgreementsEndusersPath
	if requestParams != nil {
		if requestParams.Limit != 0 {
			endpointURL = endpointURL + "?" + strconv.Itoa(requestParams.Limit)
		}
		if requestParams.Offset != 0 {
			endpointURL = endpointURL + "&" + strconv.Itoa(requestParams.Offset)
		}
	}

	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(token), &agreements)
	if err != nil {
		return nil, err
	}

	return &agreements, nil
}

// DeleteAgreement deletes an agreement for the enduser by agreementID
func (c Client) DeleteAgreement(ctx context.Context, token string, agreementID string) error {
	err := c.HTTP.Delete(ctx, AgreementsEndusersPath+agreementID, RequestHeadersWithAuth(token), nil)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAgreement updates an agreement for the enduser by agreementID
func (c Client) UpdateAgreement(ctx context.Context, token string, agreementID string, updateRequestBody UpdateRequestBody) (*Agreement, error) {
	var agreement Agreement
	// TODO: Check if this is the correct way to update an agreement, The API doc wants to append a /accept to the end of the URL
	err := c.HTTP.Put(ctx, AgreementsEndusersPath+agreementID, RequestHeadersWithAuth(token), updateRequestBody, &agreement)
	if err != nil {
		return nil, err
	}

	return &agreement, nil
}
