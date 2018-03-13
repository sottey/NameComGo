package nameapi

import (
    "fmt"
)

/// UnlockDomain Models
type UnlockDomainRequest struct {
    Url string
    Config Configuration
    Method string
}

// UnlockDomain Accessors
func (h *UnlockDomainRequest) GetMethod() string {
    return h.Method
}
func (h *UnlockDomainRequest) GetConfig() Configuration {
    return h.Config
}
func (h *UnlockDomainRequest) GetUrl() string {
    return h.Url
}

// UnlockDomain Methods
func UnlockDomain(c Configuration, domainName string) NameResponse {
    myCall := new(UnlockDomainRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s:unlock", c.BaseURL, domainName)
    myCall.Config = c
    myCall.Method = "POST"
    
    return Execute(myCall, nil)
}
