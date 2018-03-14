package nameapi

/// CancelTransfer Models
type CancelTransferRequest struct {
    DomainName string `json:"domainName"`
    url string
    config Configuration
    method string
}

// CancelTransfer Accessors
func (h *CancelTransferRequest) GetConfig() Configuration {
    return h.config
}
func (h *CancelTransferRequest) GetMethod() string {
    return h.method
}
func (h *CancelTransferRequest) GetUrl() string {
    return h.url
}

// CancelTransfer Methods
func CancelTransfer(c Configuration, domainname string) NameResponse {
    myCall := new(CreateTransferRequest)
    myCall.url = c.BaseURL + "/transfers/" + domainname + ":cancel"
    myCall.config = c
    myCall.method = "POST"
    myCall.DomainName = domainname

    return Execute(myCall, nil)
}
