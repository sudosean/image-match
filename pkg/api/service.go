package api

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
)

func init(){
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

// Service provides some "date capabilities" to your microservice
// In Go kit, you should model a service as an interface
type Service interface {
	Status(ctx context.Context) (string, error)
	GetAlgoInfo(ctx context.Context)  (Info, error)
	CreateTemplate(ctx context.Context, imageData string) (string, error)
	CompareList(ctx context.Context, templaste string, templateList []string) ([]Comparison, error)
}

type biometricService struct{}
// type algoService struct{}

// NewService makes a new Service.
func NewService() Service {
	return biometricService{}
}

// Status only tell us that our service is ok!
func (biometricService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}



// GetAlgoInfo will return a json object of  the algo
func (biometricService) GetAlgoInfo(ctx context.Context) (Info,error){
	 info := Info{
		 AlgorithmName: "Semantic Texton Forests",
		 AlgorithmVersion: "1.0.1",
		 AlgorithmType: "Face",
		 CompanyName: "MdTF",
		 TechnicalContactEmail: "john@mdtf.org",
		 RecommendedCPUs: 4,
		 RecommendedMem: 2048,
	 }
	 return info, nil
}
// create  template
func (biometricService) CreateTemplate(ctx context.Context, imageData string) (string, error){
	log.Println("image data",imageData)
	if len(imageData) == 0{
		return "{status:400, message: Bad Request. ImageData Attribute must be set}", errors.New("request error")
	}
	return "dGhpcyBzZW50ZW5jZSBpcyBhbiBleGFtcGxlIHRlbXBsYXRlLi4K", nil
}

// compare list
func (biometricService) CompareList(ctx context.Context, template string, templastelist []string) ([]Comparison, error){
	result := make([]Comparison,len(templastelist))
	for index := 0; index < len(templastelist); index++ {
		compare := Comparison{
			Score: 8734,
			NormalizedScore: 0.8734,
		}
		result = append(result, compare)
	}
	return result, nil
}

