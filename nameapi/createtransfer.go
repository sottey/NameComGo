package nameapi

import (
    "encoding/json"
)

/// CreateTransfer Models
type CreateTransferRequest struct {
    DomainName string `json:"domainName"`
    AuthCode string `json:"authCode"`
    PrivacyEnabled bool `json:"privacyEnabled"`
    PurchasePrice float64 `json:"purchasePrice"`
    PromoCode string `json:"promoCode"`
    url string
    config Configuration
    method string
}

type CreateTransferResponse struct {
    TransferData Transfer `json:"transfer"`
    OrderID int32 `json:"order"`
    TotalPaid float64 `json:"totalPaid"`
}

// CreateTransfer Accessors
func (h *CreateTransferRequest) GetConfig() Configuration {
    return h.config
}
func (h *CreateTransferRequest) GetMethod() string {
    return h.method
}
func (h *CreateTransferRequest) GetUrl() string {
    return h.url
}

// CreateTransfer Methods
func CreateTransfer(c Configuration, domainname, authcode, promocode string, privacy bool, purchaseprice float64) NameResponse {
    myCall := new(CreateTransferRequest)
    myCall.url = c.BaseURL + "/transfers"
    myCall.config = c
    myCall.method = "POST"
    myCall.DomainName = domainname
    myCall.AuthCode = authcode
    myCall.PromoCode = promocode
    myCall.PrivacyEnabled = privacy
    myCall.PurchasePrice = purchaseprice
    
    body, err := json.Marshal(myCall)
    Check("", err)
    
    return Execute(myCall, body)
}
