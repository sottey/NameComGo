package nameapi

import (
    "encoding/json"
)

/// PurchasePrivacy Models
type PurchasePrivacyRequest struct {
    DomainName string `json:"domainName"`
    PurchasePrice float64 `json:"purchasePrice"`
    Years int32 `json:"years"`
    url string
    config Configuration
    method string
    
}

type PurchasePrivacyResponse struct {
    DomainData Domain `json:"domain"`
    OrderID int32 `json:"order"`
    TotalPaid float64 `json:"totalPaid"`
}

// PurchasePrivacy Accessors
func (h *PurchasePrivacyRequest) GetConfig() Configuration {
    return h.config
}
func (h *PurchasePrivacyRequest) GetMethod() string {
    return h.method
}
func (h *PurchasePrivacyRequest) GetUrl() string {
    return h.url
}

// PurchasePrivacy Methods
func PurchasePrivacy(c Configuration, domainName string, purchasePrice float64, years int32) NameResponse {
    myCall := new(PurchasePrivacyRequest)
    myCall.DomainName = domainName
    myCall.PurchasePrice = purchasePrice
    myCall.Years = years
    myCall.url = c.BaseURL + "/domains/" + myCall.DomainName + ":purchasePrivacy"
    myCall.config = c
    myCall.method = "POST"
    
    body, err := json.Marshal(myCall)
    Check("", err)
    
    return Execute(myCall, body)
}
