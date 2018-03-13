package nameapi

import (
    "fmt"
    "encoding/json"
)

/// CreateRecord Models
type CreateRecordRequest struct {
    RecordData Record
    Url string
    Config Configuration
    Method string
}

// ListRecords Accessors
func (h *CreateRecordRequest) GetMethod() string {
    return h.Method
}
func (h *CreateRecordRequest) GetConfig() Configuration {
    return h.Config
}
func (h *CreateRecordRequest) GetUrl() string {
    return h.Url
}

// CreateRecord Methods
func CreateRecord(c Configuration, record Record) NameResponse {
    myCall := new(CreateRecordRequest)
    myCall.Url = fmt.Sprintf("%s/domains/%s/records", c.BaseURL, record.DomainName)
    myCall.Config = c
    myCall.Method = "POST"
    myCall.RecordData = record
    
    body, err := json.Marshal(myCall.RecordData)
    Check("", err)
    
	return Execute(myCall, body)
}
