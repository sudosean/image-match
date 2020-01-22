package api

import (
	"context"
	"encoding/json"
	"net/http"
)

type statusRequest struct{}

type statusResponse struct {
	Status string `json:"status"`
}
// createTemplate req/resp schema
type createTemplateRequest struct{
	ImageData string `json:"imageData"` 
}
type createTemplateResponse struct{
	Template string `json:"template"`
	Err 	 string `json:"err,omitempty"`
}
// algo info req/resp
type algoInfoRequest struct{}
type algoInfoResponse struct {
	AlgorithmName  string `json:"AlgorithmName"`
	Err            string `json:"err,omitempty"`
}
	// AlgorithmVersion string  `json:"AlgorithmVersion"`
	// AlgorithmType string  `json:"AlgorithmType"`
	// CompanyName string `json:"CompanyName"`
	// TechnicalContactEmail string `json:"TechnicalContactEmail"`
	// RecommendedCPUs int `json:"RecommendedCPUs"`
	// RecommendedMem int `json:RecommendedMem"`

// compareList req/resp schema
type compareListRequest struct {
	Template      string   `json:"Template"`
	TemplateList []string   `json:"TemplateList"`
	
}

type compareListResponse struct {
	Comparison    []Comparison  `json:"Comparison"`
	Err  		  string 	        `json:"err,omitempty"`
}
// data type for compare list response
type Comparison struct {
	Score int `json:"Score"`
	NormalizedScore float64 `json:"NormalizedScore"`
}

// In the second part we will write "decoders" for our incoming requests

func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req statusRequest
	return req, nil
}

// decode AlgoInfoRequest
func decodeAlgoInfoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req algoInfoRequest
	return req, nil
}
// decode CreateTemplate
func decodeCreateTemplateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createTemplateRequest
	return req, nil
}

// decode CompareList
func decodeCompareListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req compareListRequest
	return req, nil
}
		
// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}