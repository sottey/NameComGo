package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestListDomains(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    var ret nameapi.ListDomainsResponse
    nameResponse := nameapi.ListDomains(c)
    resp := []byte(nameResponse.Body)
    
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in ListDomains(): ", jsonErr)
    
    if nameResponse.StatusCode != 200 || len(ret.Domains) < 1 {
        t.Errorf("Bad response from ListDomains(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}