package nameapi

import (
    "fmt"
)

/// GetTransfer Models
type GetTransferRequest struct {
    DomainName string
    Url string
    Config Configuration
    Method string
}

type GetTransferResponse struct {
    DomainName string `json:"domainName"`
    Email string `json:"email"`
    Status string `json:"status"`
}

// GetTransfer Accessors
func (h *GetTransferRequest) GetMethod() string {
    return h.Method
}
func (h *GetTransferRequest) GetConfig() Configuration {
    return h.Config
}
func (h *GetTransferRequest) GetUrl() string {
    return h.Url
}

// GetTransfer Methods
func GetTransfer(c Configuration, domainName string) NameResponse {
    myCall := new(GetRecordRequest)
    myCall.Url = fmt.Sprintf("%s/transfers/%s", c.BaseURL, domainName)
    myCall.Config = c
    myCall.Method = "GET"
    
    return Execute(myCall, nil)
}
