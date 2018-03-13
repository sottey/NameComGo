package nameapi

import (
    "encoding/json"
)

/// CreateVanityNameserver Models
type CreateVanityNameserverRequest struct {
    VanityServerData VanityNameserver `json:"vanityNameserver"`
    Url string
    Config Configuration
    Method string
}

// CreateVanityNameserver Accessors
func (h *CreateVanityNameserverRequest) GetConfig() Configuration {
    return h.Config
}
func (h *CreateVanityNameserverRequest) GetMethod() string {
    return h.Method
}
func (h *CreateVanityNameserverRequest) GetUrl() string {
    return h.Url
}

// CreateVanityNameserver Methods
func CreateVanityNameserver(c Configuration, vanityServer VanityNameserver) NameResponse {
    myCall := new(CreateVanityNameserverRequest)
    myCall.Url = c.BaseURL + "/domains/" + vanityServer.DomainName + "/vanity_nameservers"
    myCall.Config = c
    myCall.Method = "POST"
    myCall.VanityServerData = vanityServer
    
    body, err := json.Marshal(myCall.VanityServerData)
    Check("", err)
    
    return Execute(myCall, body)
}
