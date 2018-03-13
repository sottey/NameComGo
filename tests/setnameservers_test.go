package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestSetNameServers(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.SetNameServers(c,"any2.org", []string{"ns1.name.com", "ns2.name.com"})
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.Domain
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in SetNameServers(): ", jsonErr)
        
        if ret.DomainName == "" {
            t.Errorf("Empty SetNameServers response!")
        }
    }
}