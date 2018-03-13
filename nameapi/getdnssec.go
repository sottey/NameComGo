package nameapi

/// CreateDNSSEC Models
type GetDNSSECRequest struct {
    DomainName string `json:"domainName"`
    Digest string `json:"digest"`
    Url string
    Config Configuration
    Method string
}

// GetDNSSECRequest Accessors
func (h *GetDNSSECRequest) GetConfig() Configuration {
    return h.Config
}
func (h *GetDNSSECRequest) GetMethod() string {
    return h.Method
}
func (h *GetDNSSECRequest) GetUrl() string {
    return h.Url
}

// GetDNSSEC Methods
func GetDNSSEC(c Configuration, domainName, digest string) NameResponse {
    myCall := new(GetDNSSECRequest)
    myCall.Url = c.BaseURL + "/domains/" + domainName + "/dnssec/" + digest
    myCall.Config = c
    myCall.Method = "GET"
    myCall.DomainName = domainName
    myCall.Digest = digest
    
    return Execute(myCall, nil)
}
