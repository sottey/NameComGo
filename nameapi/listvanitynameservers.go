package nameapi

/// ListVanityNameservers Models
type ListVanityNameserversRequest struct {
    DomainName string `json:"domainName"`
    PerPage int32 `json:"perPage,omitempty"`
    Page int32 `json:"page,omitempty"`
    Url string
    Config Configuration
    Method string
}

type ListVanityNameserversResponse struct {
    VanityServerData []VanityNameserver `json:"vanityNameservers,omitempty"`
    NextPage int32 `json:"nextPage,omitempty"`
    LastPage int32 `json:"lastPage,omitempty"`
}

// ListVanityNameservers Accessors
func (h *ListVanityNameserversRequest) GetConfig() Configuration {
    return h.Config
}
func (h *ListVanityNameserversRequest) GetMethod() string {
    return h.Method
}
func (h *ListVanityNameserversRequest) GetUrl() string {
    return h.Url
}

// ListVanityNameservers Methods
func ListVanityNameservers(c Configuration, domainName string) NameResponse {
    myCall := new(ListVanityNameserversRequest)
    myCall.Url = c.BaseURL + "/domains/" + domainName + "/vanity_nameservers"
    myCall.Config = c
    myCall.Method = "GET"
    myCall.DomainName = domainName
    
    return Execute(myCall, nil)
}
