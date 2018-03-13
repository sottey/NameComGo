package nameapi

import (
    "fmt"
)

/// ListURLForwardings Models
type ListURLForwardingsRequest struct {
    DomainName string `json:"domainName"`
    PerPage int `json:"perPage"`
    Page int `json:"page"`
    Url string
    Config Configuration
    Method string
}

// ListURLForwardings Accessors
func (h *ListURLForwardingsRequest) GetConfig() Configuration {
    return h.Config
}
func (h *ListURLForwardingsRequest) GetMethod() string {
    return h.Method
}
func (h *ListURLForwardingsRequest) GetUrl() string {
    return h.Url
}

// ListURLForwardings Methods
func ListURLForwardings(c Configuration, domainName string) NameResponse {
    myCall := new(ListURLForwardingsRequest)
    myCall.Url = fmt.Sprintf( "%s/domains/%s/url/forwarding", c.BaseURL, domainName)
    myCall.Config = c
    myCall.Method = "GET"
    
    return Execute(myCall, nil)
}
