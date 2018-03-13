package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestGetURLForwarding(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.GetURLForwarding(c, "any1.org", "www.any1.org")
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.URLForwarding
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in GetURLForwarding(): ", jsonErr)
        
        if ret.DomainName == "" {
            t.Errorf("\nObject is empty! Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}