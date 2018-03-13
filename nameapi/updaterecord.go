package nameapi

import (
    "fmt"
    "encoding/json"
)

/// UpdateRecord Models
type UpdateRecordRequest struct {
    RecordData Record
    Url string
    Config Configuration
    Method string
}

// ListRecords Accessors
func (h *UpdateRecordRequest) GetMethod() string {
    return h.Method
}
func (h *UpdateRecordRequest) GetConfig() Configuration {
    return h.Config
}
func (h *UpdateRecordRequest) GetUrl() string {
    return h.Url
}

// UpdateRecord Methods
func UpdateRecord(c Configuration, record Record) NameResponse {
    myCall := new(UpdateRecordRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s/records/%d", c.BaseURL, record.DomainName, record.ID)
    myCall.Config = c
    myCall.Method = "PUT"
    myCall.RecordData = record
    
    body, err := json.Marshal(myCall.RecordData)
    Check("", err)
    
    return Execute(myCall, body)
}
