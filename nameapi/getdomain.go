package nameapi

/// GetDomain Models
type GetDomainRequest struct {
    DomainName string `json:"domainName"`
    Url string
    Config Configuration
    Method string
    
}

// GetDomain Accessors
func (h *GetDomainRequest) GetConfig() Configuration {
    return h.Config
}
func (h *GetDomainRequest) GetMethod() string {
    return h.Method
}
func (h *GetDomainRequest) GetUrl() string {
    return h.Url
}

// GetDomain Methods
func GetDomain(c Configuration, domainName string) NameResponse {
    myCall := new(GetDomainRequest)
    myCall.DomainName = domainName
    myCall.Url = c.BaseURL + "/domains/" + myCall.DomainName
    myCall.Config = c
    myCall.Method = "GET"
    
    return Execute(myCall, nil)
}
