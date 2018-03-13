package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestCreateDNSSEC(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    dnssec := nameapi.CreateDNSSECObject("any1.org", 30909, 8, 2, "E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766")
    nameResponse := nameapi.CreateDNSSEC(c,dnssec)
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.DNSSEC
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in CreateDNSSEC(): ", jsonErr)
        
        if ret.DomainName == "" {
            t.Errorf("\nNothing returned. Expected values, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}