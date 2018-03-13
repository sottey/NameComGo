package nameapi

import (
    "encoding/json"
    "fmt"
)

/// CheckAvailability Models

// CheckAvailability Request Object
type CheckAvailabilityRequest struct {
    // Domain object
    Domains []string `json:"domainNames"`
    
    // Request URL
    url string
    
    // Request Configuration object
    config Configuration
    
    // Request Method (GET|POST)
    method string
}

// CheckAvailability Accessors
func (h *CheckAvailabilityRequest) GetMethod() string {
    return h.method
}
func (h *CheckAvailabilityRequest) GetConfig() Configuration {
    return h.config
}
func (h *CheckAvailabilityRequest) GetUrl() string {
    return h.url
}


// CheckAvailability Call
func CheckAvailability(c Configuration, domainNames []string) NameResponse {

    myCall := new(CheckAvailabilityRequest)
    myCall.url = c.BaseURL + "/domains:checkAvailability"
    myCall.config = c
    myCall.method = "POST"
    myCall.Domains = domainNames

    body, err := json.Marshal(CheckAvailabilityRequest{Domains: domainNames})
    Check("In CheckAvailability: ", err)
    
    if (c.Debug) { fmt.Println(string(body)) }
    
    return Execute(myCall, body)
}
