package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestEnableAutoRenew(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.EnableAutoRenew(c, "any1.org")
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("Bad response from EnableAutoRenew(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.Domain
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in EnableAutoRenew(): ", jsonErr)

        if ret.AutorenewEnabled != true {
            t.Errorf("Enabling Autorenew failed, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}