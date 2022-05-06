package handler

import (
	"context"
	"net/http"
)

type endpointService struct {
	service BusinessLogicService
}

type EndpointService interface {
	Tests(ctx context.Context) http.HandlerFunc
}

func NewEndpointService(service BusinessLogicService) EndpointService {
	return &endpointService{
		service: service,
	}
}

func (h *endpointService) Tests(ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := h.service.TestsService(ctx)
		if err != nil {
			HttpResponse(w, r, false, nil, err, http.StatusOK)
			return
		}
		HttpResponse(w, r, true, result, nil, http.StatusOK)
		return
	})
}
