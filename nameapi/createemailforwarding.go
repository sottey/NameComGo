package nameapi

import (
    "fmt"
    "encoding/json"
)

/// CreateEmailForwarding Models
type CreateEmailForwardingRequest struct {
    EmailForwardingData EmailForwarding
    Url string
    Config Configuration
    Method string
}

// CreateEmailForwarding Accessors
func (h *CreateEmailForwardingRequest) GetMethod() string {
    return h.Method
}
func (h *CreateEmailForwardingRequest) GetConfig() Configuration {
    return h.Config
}
func (h *CreateEmailForwardingRequest) GetUrl() string {
    return h.Url
}

// CreateEmailForwarding Methods
func CreateEmailForwarding(c Configuration, domainName, emailBox, emailTo string) NameResponse {
    myCall := new(CreateEmailForwardingRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s/email/forwarding", c.BaseURL, domainName)
    myCall.Config = c
    myCall.Method = "POST"
    myCall.EmailForwardingData = EmailForwarding {DomainName: domainName, EmailBox: emailBox, EmailTo: emailTo}
    
    body, err := json.Marshal(myCall.EmailForwardingData)
    Check("", err)
    
	return Execute(myCall, body)
}
