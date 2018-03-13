package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestDeleteURLForwarding(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.DeleteURLForwarding(c, "any1.org", "www.any1.org")
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.URLForwarding
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in UpdateURLForwardings(): ", jsonErr)
    }
}