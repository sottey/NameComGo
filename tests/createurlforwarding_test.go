package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestCreateURLForwarding(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    urlForwarding := nameapi.CreateURLForwardingObject("any2.org", nameapi.RandString(6) + ".any2.org", "http://www.test2.org", "redirect")
    
    nameResponse := nameapi.CreateURLForwarding(c, urlForwarding)
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.URLForwarding
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in CreateURLForwardings(): ", jsonErr)

        if ret.DomainName == "" {
            t.Errorf("Creating URL Forwarding failed, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}