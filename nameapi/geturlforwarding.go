package nameapi

import (
    "fmt"
)

/// GetURLForwarding Models
type GetURLForwardingRequest struct {
    DomainName string `json:"domainName"`
    PerPage int `json:"perPage"`
    Page int `json:"page"`
    Url string
    Config Configuration
    Method string
}

// GetURLForwarding Accessors
func (h *GetURLForwardingRequest) GetConfig() Configuration {
    return h.Config
}
func (h *GetURLForwardingRequest) GetMethod() string {
    return h.Method
}
func (h *GetURLForwardingRequest) GetUrl() string {
    return h.Url
}

// GetURLForwarding Methods
func GetURLForwarding(c Configuration, domainName, host string) NameResponse {
    myCall := new(GetURLForwardingRequest)
    myCall.Url = fmt.Sprintf( "%s/domains/%s/url/forwarding/%s", c.BaseURL, domainName, host)
    myCall.Config = c
    myCall.Method = "GET"
    
    return Execute(myCall, nil)
}
