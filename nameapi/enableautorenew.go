package nameapi

import (
    "fmt"
)

/// EnableAutorenew Models
type EnableAutoRenewRequest struct {
    Url string
    Config Configuration
    Method string
}

// UnlockDomain Accessors
func (h *EnableAutoRenewRequest) GetMethod() string {
    return h.Method
}
func (h *EnableAutoRenewRequest) GetConfig() Configuration {
    return h.Config
}
func (h *EnableAutoRenewRequest) GetUrl() string {
    return h.Url
}

// EnableAutoRenew Methods
func EnableAutoRenew(c Configuration, domainName string) NameResponse {
    myCall := new(EnableAutoRenewRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s:enableAutorenew", c.BaseURL, domainName)
    myCall.Config = c
    myCall.Method = "POST"
    
    return Execute(myCall, nil)
}
