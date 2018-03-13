package nameapi

import (
    "fmt"
)

/// DisableAutorenew Models
type DisableAutoRenewRequest struct {
    Url string
    Config Configuration
    Method string
}

// DisableAutoRenew Accessors
func (h *DisableAutoRenewRequest) GetMethod() string {
    return h.Method
}
func (h *DisableAutoRenewRequest) GetConfig() Configuration {
    return h.Config
}
func (h *DisableAutoRenewRequest) GetUrl() string {
    return h.Url
}

// DisableAutoRenewRequest Methods
func DisableAutoRenew(c Configuration, domainName string) NameResponse {
    myCall := new(DisableAutoRenewRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s:disableAutorenew", c.BaseURL, domainName)
    myCall.Config = c
    myCall.Method = "POST"

    return Execute(myCall, nil)
}
