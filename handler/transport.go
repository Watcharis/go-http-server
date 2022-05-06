package handler

import (
	"context"
	"net/http"
)

type TransportService interface {
	Tests(ctx context.Context) http.Handler
}

type transportService struct {
	endpointService EndpointService
}

func NewTransportService(endpointService EndpointService) TransportService {
	return &transportService{
		endpointService: endpointService,
	}
}

func (t *transportService) Tests(ctx context.Context) http.Handler {
	return t.endpointService.Tests(ctx)
}
