package nameapi

import (
    "fmt"
)

/// GetRecord Models
type GetRecordRequest struct {
    id uint32
    Url string
    Config Configuration
    Method string
}

// GetRecord Accessors
func (h *GetRecordRequest) GetMethod() string {
    return h.Method
}
func (h *GetRecordRequest) GetConfig() Configuration {
    return h.Config
}
func (h *GetRecordRequest) GetUrl() string {
    return h.Url
}

// GetRecord Methods
func GetRecord(c Configuration, domainName string, id int32) NameResponse {
    myCall := new(GetRecordRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s/records/%d", c.BaseURL, domainName, id)
    myCall.Config = c
    myCall.Method = "GET"
    
    return Execute(myCall, nil)
}
