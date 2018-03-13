package nameapi

import (
    "fmt"
)

/// DeleteRecord Models
type DeleteRecordRequest struct {
    ID int32
    Url string
    Config Configuration
    Method string
}

// ListRecords Accessors
func (h *DeleteRecordRequest) GetMethod() string {
    return h.Method
}
func (h *DeleteRecordRequest) GetConfig() Configuration {
    return h.Config
}
func (h *DeleteRecordRequest) GetUrl() string {
    return h.Url
}

// DeleteRecord Methods
func DeleteRecord(c Configuration, domainName string, id int32) NameResponse {
    myCall := new(DeleteRecordRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s/records/%d", c.BaseURL, domainName, id)
    myCall.Config = c
    myCall.Method = "DELETE"
    
    return Execute(myCall, nil)
}
