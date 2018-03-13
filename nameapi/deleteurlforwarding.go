package nameapi

import (
    "fmt"
)

/// DeleteURLForwarding Models
type DeleteURLForwardingRequest struct {
    DomainName string `json:"domainName"`
    PerPage int `json:"perPage"`
    Page int `json:"page"`
    Url string
    Config Configuration
    Method string
}

// DeleteURLForwarding Accessors
func (h *DeleteURLForwardingRequest) GetConfig() Configuration {
    return h.Config
}
func (h *DeleteURLForwardingRequest) GetMethod() string {
    return h.Method
}
func (h *DeleteURLForwardingRequest) GetUrl() string {
    return h.Url
}

// DeleteURLForwarding Methods
func DeleteURLForwarding(c Configuration, domainName, host string) NameResponse {
    myCall := new(DeleteURLForwardingRequest)
    myCall.Url = fmt.Sprintf( "%s/domains/%s/url/forwarding/%s", c.BaseURL, domainName, host)
    myCall.Config = c
    myCall.Method = "DELETE"
    
    return Execute(myCall, nil)
}
