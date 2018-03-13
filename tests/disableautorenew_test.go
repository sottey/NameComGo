package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestDisableAutoRenew(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.DisableAutoRenew(c, "any1.org")
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("Bad response from DisableAutoRenew(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.Domain
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in DisableAutoRenew(): ", jsonErr)

        if ret.AutorenewEnabled == true {
            t.Errorf("Disabling Autorenew failed, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}