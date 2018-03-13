package nameapi

import (
    "encoding/json"
)

/// SetContacts Models
type SetContactsRequest struct {
    DomainName string `json:"domainName"`
    ContactData ContactsWrapper `json:"contacts"`
    Url string
    Config Configuration
    Method string
}

// SetContacts Accessors
func (h *SetContactsRequest) GetConfig() Configuration {
    return h.Config
}
func (h *SetContactsRequest) GetMethod() string {
    return h.Method
}
func (h *SetContactsRequest) GetUrl() string {
    return h.Url
}

// SetContactsRequest Methods
func SetContacts(c Configuration, domainName string, contacts ContactsWrapper) NameResponse {
    myCall := new(SetContactsRequest)
    myCall.Url = c.BaseURL + "/domains/" + domainName + ":setContacts"
    myCall.Config = c
    myCall.Method = "POST"
    myCall.DomainName = domainName
    myCall.ContactData = contacts
    
    body, err := json.Marshal(myCall.ContactData)
    Check("", err)
    
    return Execute(myCall, body)
}
