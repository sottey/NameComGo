package nameapi

/// ListDomains Models
type ListDomainsRequest struct {
    PerPage int `json:"perPage"`
    Page int `json:"page"`
    Url string
    Config Configuration
    Method string
}

type ListDomainsResponse struct {
    Domains []Domain `json:"domains,omitempty"`
    NextPage int32 `json:"nextPage,omitempty"`
    LastPage int32 `json:"lastPage,omitempty"`
}

// Hello Accessors
func (h *ListDomainsRequest) GetConfig() Configuration {
    return h.Config
}
func (h *ListDomainsRequest) GetMethod() string {
    return h.Method
}
func (h *ListDomainsRequest) GetUrl() string {
    return h.Url
}

// ListDomains Methods
func ListDomains(c Configuration) NameResponse {
    myCall := new(ListDomainsRequest)
    myCall.Url = c.BaseURL + "/domains/"
    myCall.Config = c
    myCall.Method = "GET"
    
    return Execute(myCall, nil)
}
