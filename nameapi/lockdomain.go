package nameapi

import (
    "fmt"
)

/// LockDomain Models
type LockDomainRequest struct {
    Url string
    Config Configuration
    Method string
}

// LockDomain Accessors
func (h *LockDomainRequest) GetMethod() string {
    return h.Method
}
func (h *LockDomainRequest) GetConfig() Configuration {
    return h.Config
}
func (h *LockDomainRequest) GetUrl() string {
    return h.Url
}

// LockDomain Methods
func LockDomain(c Configuration, domainName string) NameResponse {
    myCall := new(LockDomainRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s:lock", c.BaseURL, domainName)
    myCall.Config = c
    myCall.Method = "POST"

    return Execute(myCall, nil)
}
