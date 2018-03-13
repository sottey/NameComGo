package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestCreateEmailForwarding(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.CreateEmailForwarding(c,"any2.org", "admin", "admintest@any1.org")
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.EmailForwarding
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in CreateEmailForwarding(): ", jsonErr)
        
        if ret.DomainName == "" {
            t.Errorf("Incorrect CreateEmailForwarding response! Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}