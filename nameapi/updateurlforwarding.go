package nameapi

import (
    "fmt"
    "encoding/json"
)

/// UpdateURLForwarding Models
type UpdateURLForwardingRequest struct {
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
type UpdateURLForwardingResponse struct {
    DomainName string `json:"domainName"`
    Host string `json:"host"`
    ForwardsTo string `json:"forwardsTo"`
    Type string `json:"type"`
    Title string `json:"title"`
    Meta string `json:"meta"`
    BodyText string
}

// UpdateURLForwardings Accessors
func (h *UpdateURLForwardingRequest) GetConfig() Configuration {
    return h.Config
}
func (h *UpdateURLForwardingRequest) GetMethod() string {
    return h.Method
}
func (h *UpdateURLForwardingRequest) GetUrl() string {
    return h.Url
}

// UpdateURLForwarding Methods
func UpdateURLForwarding(c Configuration, urlForwarding URLForwarding) NameResponse {
    myCall := new(UpdateURLForwardingRequest)
    myCall.Url = fmt.Sprintf( "%s/domains/%s/url/forwarding/%s", c.BaseURL, urlForwarding.DomainName,urlForwarding.Host)
    myCall.Config = c
    myCall.Method = "PUT"
        
    body, err := json.Marshal(urlForwarding)
    Check("", err)

    return Execute(myCall, body)
}
