package nameapi

import (
    "fmt"
)

/// ListRecords Models
type ListRecordsRequest struct {
    DomainName string
    Url string
    Config Configuration
    Method string
}

type ListRecordsResponse struct {
    Records []Record `json:"records"`
}

// ListRecords Accessors
func (h *ListRecordsRequest) GetMethod() string {
    return h.Method
}
func (h *ListRecordsRequest) GetConfig() Configuration {
    return h.Config
}
func (h *ListRecordsRequest) GetUrl() string {
    return h.Url
}

// ListRecords Methods
func ListRecords(c Configuration, domainName string) NameResponse {
    myCall := new(ListRecordsRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s/records", c.BaseURL, domainName)
    myCall.Config = c
    myCall.Method = "GET"
    
    return Execute(myCall, nil)
}
