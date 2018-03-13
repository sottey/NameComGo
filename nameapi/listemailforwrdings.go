package nameapi

import (
    "fmt"
)

/// ListEmailForwardings Models
type ListEmailForwardingsRequest struct {
    DomainName string `json:"domainName"`
    PerPage int32
    Page int32
    Url string
    Config Configuration
    Method string
}

type ListEmailForwardingsResponse struct {
    EmailForwarding []EmailForwarding `json:"emailForwarding"`
}

// ListEmailForwardings Accessors
func (h *ListEmailForwardingsRequest) GetMethod() string {
    return h.Method
}
func (h *ListEmailForwardingsRequest) GetConfig() Configuration {
    return h.Config
}
func (h *ListEmailForwardingsRequest) GetUrl() string {
    return h.Url
}

// ListEmailForwardings Methods
func ListEmailForwardings(c Configuration, domainName string) NameResponse {
    myCall := new(ListEmailForwardingsRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s/email/forwarding", c.BaseURL, domainName)
    myCall.Config = c
    myCall.Method = "GET"
    
    return Execute(myCall, nil)
}
