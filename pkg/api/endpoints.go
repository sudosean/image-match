package api

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	StatusEndpoint   		endpoint.Endpoint
	GetAlgoInfoEndpoint 	endpoint.Endpoint
	CreateTemplateEndpoint 	endpoint.Endpoint
	CompareListEndpoint		endpoint.Endpoint
}


// MakeGetAlgoInfoEndpoint returns the response from our service "getAlgoInfo"
func MakeGetAlgoInfoEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(algoInfoRequest) // we really just need the request, we don't use any value from it
		d, err := srv.GetAlgoInfo(ctx)
		if err != nil {
			return algoInfoResponse{d, err.Error()}, nil
		}
		return algoInfoResponse{d, ""}, nil
	}
}

// MakeCreateTemplateEndpoint
func MakeCreateTemplateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createTemplateRequest) // same here
		d, err := srv.CreateTemplate(ctx, req.ImageData)
		if err != nil {
			return createTemplateResponse{d, err.Error()}, nil
		}
		return createTemplateResponse{d, ""}, nil
	}
}
// MakeStatusEndpoint returns the response from our service "status"
func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(statusRequest) // same here
		s, err := srv.Status(ctx)
		if err != nil {
			return statusResponse{s}, err
		}
		return statusResponse{s}, nil
	}
}

// MakeCompareListEndpoint returns a list of compared items
func MakeCompareListEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(compareListRequest)
		list, err := srv.CompareList(ctx, req.Template, req.TemplateList)
		if err != nil {
			return compareListResponse{list, err.Error()}, nil
		}
		return compareListResponse{list,""}, nil
	}
}

// Status endpoint mapping
func (e Endpoints) Status(ctx context.Context) (string, error) {
	req := statusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	statusResp := resp.(statusResponse)
	return statusResp.Status, nil
}

// get Algo info mapping
func (e Endpoints) AlgoInfo(ctx context.Context) (Info, error){
	req := algoInfoRequest{}
	resp, err := e.GetAlgoInfoEndpoint(ctx, req)
	if err != nil {
		return  Info{} ,err // just returning empty struct for now with err
	}
	getAlgoInfoResp := resp.(algoInfoResponse)
	if getAlgoInfoResp.Err != "" {
		return Info{}, errors.New(getAlgoInfoResp.Err)
	}
	return getAlgoInfoResp.Info, nil
}

// create template mapping
func (e Endpoints) CreateTemplate(ctx context.Context, imageData string) (string, error){
	
	req := createTemplateRequest{ImageData: imageData}
	
	resp, err := e.CreateTemplateEndpoint(ctx, req)
	if err != nil {
		return "error", err
	}
	CreateTemplateResp := resp.(createTemplateResponse)
	if CreateTemplateResp.Err != "" {
		return "{status: 500 , message: error creating template}", errors.New(CreateTemplateResp.Err)
	}
	return CreateTemplateResp.Template, nil
}

// compare list mapping
func (e Endpoints) CompareList(ctx context.Context, template string, templateList []string) ([]Comparison, error){
	req := compareListRequest{Template: template, TemplateList: templateList}
	res, err := e.CompareListEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	CompareListResponse := res.(compareListResponse)
	if CompareListResponse.Err != "" {
		return nil, errors.New(CompareListResponse.Err)
	}
	return CompareListResponse.Comparison, nil
}