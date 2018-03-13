package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestListEmailForwardings(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.ListEmailForwardings(c,"any2.org")
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.ListEmailForwardingsResponse
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in ListEmailForwardings(): ", jsonErr)
        
        if len(ret.EmailForwarding) == 0 {
            t.Errorf("Incorrect ListEmailForwardings response! Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}
