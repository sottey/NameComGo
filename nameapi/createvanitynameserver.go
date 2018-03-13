package nameapi

import (
    "encoding/json"
)

/// CreateVanitNameserver Models
type CreateVanitNameserverRequest struct {
    VanityServerData VanityNameserver `json:"vanityNameserver"`
    Url string
    Config Configuration
    Method string
}

// CreateVanitNameserver Accessors
func (h *CreateVanitNameserverRequest) GetConfig() Configuration {
    return h.Config
}
func (h *CreateVanitNameserverRequest) GetMethod() string {
    return h.Method
}
func (h *CreateVanitNameserverRequest) GetUrl() string {
    return h.Url
}

// CreateVanitNameserver Methods
func CreateVanityNameserver(c Configuration, vanityServer VanityNameserver) NameResponse {
    myCall := new(CreateVanitNameserverRequest)
    myCall.Url = c.BaseURL + "/domains/" + vanityServer.DomainName + "/vanity_nameservers"
    myCall.Config = c
    myCall.Method = "POST"
    myCall.VanityServerData = vanityServer
    
    body, err := json.Marshal(myCall.VanityServerData)
    Check("", err)
    
    return Execute(myCall, body)
}
