package nameapi

import (
    "encoding/json"
)

/// SetNameServers Models
type SetNameServersRequest struct {
    DomainName string `json:"domainName"`
    NameServers []string `json:"nameservers"`
    Url string
    Config Configuration
    Method string
}

// SetNameServers Accessors
func (h *SetNameServersRequest) GetConfig() Configuration {
    return h.Config
}
func (h *SetNameServersRequest) GetMethod() string {
    return h.Method
}
func (h *SetNameServersRequest) GetUrl() string {
    return h.Url
}

// SetNameServers Methods
func SetNameServers(c Configuration, domainName string, nameservers []string) NameResponse {
    myCall := new(SetNameServersRequest)
    myCall.Url = c.BaseURL + "/domains/" + domainName + ":setNameservers"
    myCall.Config = c
    myCall.Method = "POST"
    myCall.DomainName = domainName
    myCall.NameServers = nameservers
    
    body, err := json.Marshal(myCall.NameServers)
    Check("", err)
    
    return Execute(myCall, body)
}
