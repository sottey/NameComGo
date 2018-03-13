package nameapi

/// DeleteDNSSEC Models
type DeleteDNSSECRequest struct {
    DomainName string `json:"domainName"`
    Digest string `json:"digest"`
    Url string
    Config Configuration
    Method string
}

// DeleteDNSSEC Accessors
func (h *DeleteDNSSECRequest) GetConfig() Configuration {
    return h.Config
}
func (h *DeleteDNSSECRequest) GetMethod() string {
    return h.Method
}
func (h *DeleteDNSSECRequest) GetUrl() string {
    return h.Url
}

// DeleteDNSSEC Methods
func DeleteDNSSEC(c Configuration, domainName, digest string) NameResponse {
    myCall := new(DeleteDNSSECRequest)
    myCall.Url = c.BaseURL + "/domains/" + domainName + "/dnssec/" + digest
    myCall.Config = c
    myCall.Method = "DELETE"
    myCall.DomainName = domainName
    myCall.Digest = digest
    
    return Execute(myCall, nil)
}
