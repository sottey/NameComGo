package nameapi

import (
    "fmt"
    "encoding/json"
)

/// UpdateEmailForwarding Models
type UpdateEmailForwardingRequest struct {
    EmailForwardingData EmailForwarding
    Url string
    Config Configuration
    Method string
}

// UpdateEmailForwarding Accessors
func (h *UpdateEmailForwardingRequest) GetMethod() string {
    return h.Method
}
func (h *UpdateEmailForwardingRequest) GetConfig() Configuration {
    return h.Config
}
func (h *UpdateEmailForwardingRequest) GetUrl() string {
    return h.Url
}

// UpdateEmailForwarding Methods
func UpdateEmailForwarding(c Configuration, domainName, emailBox, emailTo string) NameResponse {
    myCall := new(UpdateEmailForwardingRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s/email/forwarding/%v", c.BaseURL, domainName, emailBox)
    myCall.Config = c
    myCall.Method = "PUT"
    myCall.EmailForwardingData = EmailForwarding {DomainName: domainName, EmailBox: emailBox, EmailTo: emailTo}
    
    body, err := json.Marshal(myCall.EmailForwardingData)
    Check("Error handling UpdateEmailForwarding data", err)
    
	return Execute(myCall, body)
}
