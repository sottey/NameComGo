package nameapi
    
import (
    "encoding/json"
)

/// Search Models
type SearchRequest struct {
    Timeout int32 `json:"timeout,omitempty"`
    Keyword string `json:"keyword,omitempty"`
    TLDFilter []string `json:"tldFilter,omitempty"`
    PromoCode string `json:"promoCode,omitempty"`
    url string
    config Configuration
    method string
    
}

type SearchResults struct {
    Responses []SearchResponse `json:"results"`
}

type SearchResponse struct {
    DomainName string `json:"domainName"`
    Sld string `json:"tld"`
    Tld string `json:"sld"`
    Purchasable bool `json:"purchasable"`
    Premium bool `json:"premium"`
    PurchasePrice float64 `json:"purchasePrice"`
    PurchaseType string `json:"purchaseType"`
    RenewalPrice float64 `json:"renewalPrice"`
}


// Search Accessors
func (h *SearchRequest) GetConfig() Configuration {
    return h.config
}
func (h *SearchRequest) GetMethod() string {
    return h.method
}
func (h *SearchRequest) GetUrl() string {
    return h.url
}

// Search Methods
func Search(c Configuration, keyword string, tldFilter []string) NameResponse {
    myCall := new(SearchRequest)
    myCall.url = c.BaseURL + "/domains:search"
    myCall.config = c
    myCall.method = "POST"
    myCall.Keyword = keyword

    body, err := json.Marshal(SearchRequest{Keyword: keyword, TLDFilter: tldFilter})
    Check("", err)

    return Execute(myCall, body)
}
