package api

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a good little server
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

	r.Methods("GET").Path("/status").Handler(httptransport.NewServer(
		endpoints.StatusEndpoint,
		decodeStatusRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/v1/info").Handler(httptransport.NewServer(
		endpoints.GetAlgoInfoEndpoint,
		decodeAlgoInfoRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/v1/create-template").Handler(httptransport.NewServer(
		endpoints.CreateTemplateEndpoint,
		decodeCreateTemplateRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/v1/compare-list").Handler(httptransport.NewServer(
		endpoints.CompareListEndpoint,
		decodeCompareListRequest,
		encodeResponse,
	))


	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}