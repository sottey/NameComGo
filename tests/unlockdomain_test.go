package examples

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestUnlockDomain(t *testing.T) {
    c  := nameapi.GetConfig("test")
    
    //Test data
    domainName := "Any4.org"
    
    nameResponse := nameapi.UnlockDomain(c, domainName)
    resp := []byte(nameResponse.Body)
    
    var ret nameapi.Domain
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in UnlockDomain(): ", jsonErr)
    
    if nameResponse.StatusCode != 200 || ret.DomainName == "" {
        t.Errorf("Bad response from UnlockDomain(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}