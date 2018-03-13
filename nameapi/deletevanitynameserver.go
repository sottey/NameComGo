package nameapi

import (
    "encoding/json"
)

/// DeleteVanityNameserver Models
type DeleteVanityNameserverRequest struct {
    VanityServerData VanityNameserver `json:"vanityNameserver"`
    Url string
    Config Configuration
    Method string
}

// CreateVanityNameserver Accessors
func (h *DeleteVanityNameserverRequest) GetConfig() Configuration {
    return h.Config
}
func (h *DeleteVanityNameserverRequest) GetMethod() string {
    return h.Method
}
func (h *DeleteVanityNameserverRequest) GetUrl() string {
    return h.Url
}

// DeleteVanityNameserver Methods
func DeleteVanityNameserver(c Configuration, vanityServer VanityNameserver) NameResponse {
    myCall := new(DeleteVanityNameserverRequest)
    myCall.Url = c.BaseURL + "/domains/" + vanityServer.DomainName + "/vanity_nameservers/" + vanityServer.Hostname
    myCall.Config = c
    myCall.Method = "DELETE"
    myCall.VanityServerData = vanityServer
    
    body, err := json.Marshal(myCall.VanityServerData)
    Check("", err)
    
    return Execute(myCall, body)
}
