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
