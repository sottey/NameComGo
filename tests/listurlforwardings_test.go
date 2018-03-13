package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestListURLForwardings(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.ListURLForwardings(c, "any2.org")
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.URLForwardingGroup
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in ListURLForwardings(): ", jsonErr)
        
        if len(ret.URLForwardings) < 1 {
            t.Errorf("\nObject is empty! Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}