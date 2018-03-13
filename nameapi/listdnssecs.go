package nameapi

/// ListDNSSECs Models
type ListDNSSECsRequest struct {
    DomainName string `json:"domainName"`
    Url string
    Config Configuration
    Method string
}

type ListDNSSECsResponse struct {
    DNSSECData []DNSSEC `json:"dnssec,omitempty"`
    NextPage int32 `json:"nextPage,omitempty"`
    LastPage int32 `json:"lastPage,omitempty"`
}

// ListDNSSECs Accessors
func (h *ListDNSSECsRequest) GetConfig() Configuration {
    return h.Config
}
func (h *ListDNSSECsRequest) GetMethod() string {
    return h.Method
}
func (h *ListDNSSECsRequest) GetUrl() string {
    return h.Url
}

// ListDNSSECs Methods
func ListDNSSECs(c Configuration, domainName string) NameResponse {
    myCall := new(ListDNSSECsRequest)
    myCall.Url = c.BaseURL + "/domains/" + domainName + "/dnssec"
    myCall.Config = c
    myCall.Method = "GET"
    myCall.DomainName = domainName
    
    return Execute(myCall, nil)
}
