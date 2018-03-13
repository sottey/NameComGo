package nameapi

import (
    "encoding/json"
)

/// UpdateVanitNameserver Models
type UpdateVanityNameserverRequest struct {
    VanityServerData VanityNameserver `json:"vanityNameserver"`
    Url string
    Config Configuration
    Method string
}

// UpdateVanityNameserver Accessors
func (h *UpdateVanityNameserverRequest) GetConfig() Configuration {
    return h.Config
}
func (h *UpdateVanityNameserverRequest) GetMethod() string {
    return h.Method
}
func (h *UpdateVanityNameserverRequest) GetUrl() string {
    return h.Url
}

// UpdateVanityNameserver Methods
func UpdateVanityNameserver(c Configuration, vanityServer VanityNameserver) NameResponse {
    myCall := new(UpdateVanityNameserverRequest)
    myCall.Url = c.BaseURL + "/domains/" + vanityServer.DomainName + "/vanity_nameservers/" + vanityServer.Hostname
    myCall.Config = c
    myCall.Method = "PUT"
    myCall.VanityServerData = vanityServer
    
    body, err := json.Marshal(myCall.VanityServerData)
    Check("", err)
    
    return Execute(myCall, body)
}
