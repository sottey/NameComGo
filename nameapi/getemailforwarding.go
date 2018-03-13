package nameapi

import (
    "fmt"
)

/// GetEmailForwarding Models
type GetEmailForwardingRequest struct {
    Url string
    Config Configuration
    Method string
}

// GetEmailForwarding Accessors
func (h *GetEmailForwardingRequest) GetMethod() string {
    return h.Method
}
func (h *GetEmailForwardingRequest) GetConfig() Configuration {
    return h.Config
}
func (h *GetEmailForwardingRequest) GetUrl() string {
    return h.Url
}

// GetEmailForwarding Methods
func GetEmailForwarding(c Configuration, domainName, emailBox string) NameResponse {
    myCall := new(GetEmailForwardingRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s/email/forwarding/%v", c.BaseURL, domainName, emailBox)
    myCall.Config = c
    myCall.Method = "GET"
    
	return Execute(myCall, nil)
}
