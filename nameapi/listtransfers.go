package nameapi

import (
    "fmt"
)

/// ListTransfers Models
type ListTransfersRequest struct {
    PerPage int32
    Page int32
    Url string
    Config Configuration
    Method string
}

type ListTransfersResponse struct {
    Transfers []Transfer `json:"transfers"`
}

// ListTransfers Accessors
func (h *ListTransfersRequest) GetMethod() string {
    return h.Method
}
func (h *ListTransfersRequest) GetConfig() Configuration {
    return h.Config
}
func (h *ListTransfersRequest) GetUrl() string {
    return h.Url
}

// ListTransfers Methods
func ListTransfers(c Configuration) NameResponse {
    myCall := new(ListTransfersRequest)
    myCall.Url = fmt.Sprintf("%s/transfers", c.BaseURL)
    myCall.Config = c
    myCall.Method = "GET"
    
    return Execute(myCall, nil)
}
