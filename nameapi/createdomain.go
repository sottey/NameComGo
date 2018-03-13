package nameapi

import (
    "encoding/json"
)

/// CreateDomain Models
type CreateDomainRequest struct {
    PurchaseData PurchaseDetails `json:"domain"`
    PurchasePrice float64 `json:"purchasePrice"`
    PurchaseType string `json:"purchaseType"`
    Years int32 `json:"years"`
    TldRequirements map[string]string `json:"tldRequirements"`
    PromoCode string // Not yet implemented
    Url string
    Config Configuration
    Method string
}

// CreateDomain Accessors
func (h *CreateDomainRequest) GetMethod() string {
    return h.Method
}
func (h *CreateDomainRequest) GetPurchaseData() PurchaseDetails {
    return h.PurchaseData
}
func (h *CreateDomainRequest) GetConfig() Configuration {
    return h.Config
}
func (h *CreateDomainRequest) GetUrl() string {
    return h.Url
}


// CreateDomain Call
func CreateDomain(c Configuration, purchase PurchaseDetails) NameResponse {
    myCall := new(CreateDomainRequest)
    myCall.Url = c.BaseURL + "/domains"
    myCall.Config = c
    myCall.Method = "POST"
    
    myCall.PurchaseData = purchase

    body, err := json.Marshal(myCall.PurchaseData)
    
    Check("", err)

	return Execute(myCall, body)
}
