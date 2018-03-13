package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestGetDomain(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    var ret nameapi.Domain

    domainName := "any1.org"
    nameResponse := nameapi.GetDomain(c, domainName)
    resp := []byte(nameResponse.Body)
    
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in GetDomains(): ", jsonErr)
    
    if nameResponse.StatusCode != 200 || ret.DomainName == "" {
        t.Errorf("Bad response from GetDomains(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}