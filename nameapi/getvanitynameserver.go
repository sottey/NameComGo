package nameapi

/// GetVanityNameserver Models
type GetVanityNameserverRequest struct {
    DomainName string `json:"domainName"`
    HostName string `json:"hostname"`
    PerPage int32 `json:"perPage,omitempty"`
    Page int32 `json:"page,omitempty"`
    Url string
    Config Configuration
    Method string
}

// GetVanityNameserver Accessors
func (h *GetVanityNameserverRequest) GetConfig() Configuration {
    return h.Config
}
func (h *GetVanityNameserverRequest) GetMethod() string {
    return h.Method
}
func (h *GetVanityNameserverRequest) GetUrl() string {
    return h.Url
}

// GetVanityNameserver Methods
func GetVanityNameserver(c Configuration, domainName, hostname string) NameResponse {
    myCall := new(GetVanityNameserverRequest)
    myCall.Url = c.BaseURL + "/domains/" + domainName + "/vanity_nameservers/" + hostname
    myCall.Config = c
    myCall.Method = "GET"
    myCall.DomainName = domainName
    myCall.HostName = hostname
    
    return Execute(myCall, nil)
}
