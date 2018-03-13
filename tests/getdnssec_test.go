package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestGetDNSSEC(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    // Test data
    domainName := "any99.org"
    digest := "E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766"
    
    nameResponse := nameapi.GetDNSSEC(c,domainName, digest)
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.ListDNSSECsResponse
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in ListDNSSECs(): ", jsonErr)
        
        if len(ret.DNSSECData) < 1 {
            t.Errorf("\nNothing returned. Expected values, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}