package agreements

import (
	"context"
	"github.com/weportfolio/go-nordigen/consts"
)

func (c Client) Post(ctx context.Context, token string, agreementRequestBody AgreementRequestBody) (*Agreement, error) {
	var agreement Agreement
	err := c.HTTP.Post(ctx, consts.AgreementsEndusersPath, consts.RequestHeadersWithAuth(token), agreementRequestBody, &agreement)
	if err != nil {
		return nil, err
	}

	return &agreement, nil
}
