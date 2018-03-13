package examples

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestLockDomain(t *testing.T) {
    c  := nameapi.GetConfig("test")
    
    //Test data
    domainName := "Any4.org"
    
    nameResponse := nameapi.LockDomain(c, domainName)
    resp := []byte(nameResponse.Body)
    
    var ret nameapi.Domain
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in LockDomain(): ", jsonErr)
    
    if nameResponse.StatusCode != 200 || ret.DomainName == "" {
        t.Errorf("Bad response from LockDomain(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}