package nameapi

import (
    "encoding/json"
)

/// CreateDNSSEC Models
type CreateDNSSECRequest struct {
    DNSSECData DNSSEC `json:"dnssec"`
    Url string
    Config Configuration
    Method string
}

// CreateDNSSEC Accessors
func (h *CreateDNSSECRequest) GetConfig() Configuration {
    return h.Config
}
func (h *CreateDNSSECRequest) GetMethod() string {
    return h.Method
}
func (h *CreateDNSSECRequest) GetUrl() string {
    return h.Url
}

// CreateDNSSEC Methods
func CreateDNSSEC(c Configuration, dnsSEC DNSSEC) NameResponse {
    myCall := new(CreateDNSSECRequest)
    myCall.Url = c.BaseURL + "/domains/" + dnsSEC.DomainName + "/dnssec"
    myCall.Config = c
    myCall.Method = "POST"
    myCall.DNSSECData = dnsSEC
    
    body, err := json.Marshal(myCall.DNSSECData)
    Check("", err)
    
    return Execute(myCall, body)
}
