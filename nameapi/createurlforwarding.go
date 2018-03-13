package nameapi

import (
    "fmt"
    "encoding/json"
)

/// CreateURLForwarding Models
type CreateURLForwardingRequest struct {
    DomainName string `json:"domainName"`
    Host string `json:"Host"`
    ForwardsTo string `json:"forwardsTo"`
    Type string `json:"type"`
    Title string `json:"title"`
    Meta string `json:"meta"`
    Url string
    Config Configuration
    Method string
}

// CreateURLForwardings Accessors
func (h *CreateURLForwardingRequest) GetConfig() Configuration {
    return h.Config
}
func (h *CreateURLForwardingRequest) GetMethod() string {
    return h.Method
}
func (h *CreateURLForwardingRequest) GetUrl() string {
    return h.Url
}

// CreateURLForwarding Methods
func CreateURLForwarding(c Configuration, urlForwarding URLForwarding) NameResponse {
    myCall := new(CreateURLForwardingRequest)
    myCall.Url = fmt.Sprintf( "%s/domains/%s/url/forwarding", c.BaseURL, urlForwarding.DomainName)
    myCall.Config = c
    myCall.Method = "POST"
        
    body, err := json.Marshal(urlForwarding)
    Check("", err)
    
    return Execute(myCall, body)
}
