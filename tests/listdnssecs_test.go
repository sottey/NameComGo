package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestListDNSSECs(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.ListDNSSECs(c,"any2.org")
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