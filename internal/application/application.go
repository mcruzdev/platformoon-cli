package application

import (
	"log"
	"net/http"

	"github.com/mcruzdev/platformoon-cli/internal/httpclient"
)

type ApplicationRequest struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Technology string `json:"technology"`
}

func NewApplicationRequest(name, appType, technology string) *ApplicationRequest {
	return &ApplicationRequest{
		Name:       name,
		Technology: technology,
		Type:       appType,
	}
}

func CreateNewApplication(name, appType, technology string) {
	ar := NewApplicationRequest(name, appType, technology)
	r, err := httpclient.Post("http://localhost:8080/applications", ar)

	if err != nil {
		panic(any(err))
	}

	if r.StatusCode == http.StatusCreated {
		log.Printf("Application created: %s", r.Header["Location"])
	}
}
