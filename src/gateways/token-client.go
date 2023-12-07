package gateways

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
	tokenService "github.com/schema-creator/graph-gateway/pkg/grpc/token-service/v1"
)

type tokenClient struct {
	client tokenService.TokenClient
}

type TokenClient interface {
	CreateToken(ctx context.Context, txn *newrelic.Transaction, arg *tokenService.CreateTokenRequest) (*tokenService.CreateTokenResponse, error)
	GetToken(ctx context.Context, txn *newrelic.Transaction, arg *tokenService.GetTokenRequest) (*tokenService.GetTokenResponse, error)
	DeleteToken(ctx context.Context, txn *newrelic.Transaction, arg *tokenService.DeleteTokenRequest) (*tokenService.DeleteTokenResponse, error)
}

func NewTokenClient(client tokenService.TokenClient) TokenClient {
	return &tokenClient{
		client: client,
	}
}

func (t *tokenClient) CreateToken(ctx context.Context, txn *newrelic.Transaction, arg *tokenService.CreateTokenRequest) (*tokenService.CreateTokenResponse, error) {
	defer txn.StartSegment("CreateToken-client").End()
	return t.client.CreateToken(ctx, arg)
}

func (t *tokenClient) GetToken(ctx context.Context, txn *newrelic.Transaction, arg *tokenService.GetTokenRequest) (*tokenService.GetTokenResponse, error) {
	defer txn.StartSegment("GetToken-client").End()
	return t.client.GetToken(ctx, arg)
}

func (t *tokenClient) DeleteToken(ctx context.Context, txn *newrelic.Transaction, arg *tokenService.DeleteTokenRequest) (*tokenService.DeleteTokenResponse, error) {
	defer txn.StartSegment("Delete-client").End()
	return t.client.DeleteToken(ctx, arg)
}
