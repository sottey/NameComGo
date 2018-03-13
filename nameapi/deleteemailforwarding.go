package nameapi

import (
    "fmt"
)

/// DeleteEmailForwarding Models
type DeleteEmailForwardingRequest struct {
    EmailForwardingData EmailForwarding
    Url string
    Config Configuration
    Method string
}

// DeleteEmailForwarding Accessors
func (h *DeleteEmailForwardingRequest) GetMethod() string {
    return h.Method
}
func (h *DeleteEmailForwardingRequest) GetConfig() Configuration {
    return h.Config
}
func (h *DeleteEmailForwardingRequest) GetUrl() string {
    return h.Url
}

// DeleteEmailForwarding Methods
func DeleteEmailForwarding(c Configuration, domainName, emailBox string) NameResponse {
    myCall := new(DeleteEmailForwardingRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s/email/forwarding/%v", c.BaseURL, domainName, emailBox)
    myCall.Config = c
    myCall.Method = "DELETE"
    
	return Execute(myCall, nil)
}
