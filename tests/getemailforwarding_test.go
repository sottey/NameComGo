package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestGetEmailForwarding(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    // Test data
    domainName := "any99.org"
    emailBox := "admin"
    
    nameResponse := nameapi.GetEmailForwarding(c,domainName, emailBox)
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.EmailForwarding
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in GetEmailForwarding(): ", jsonErr)
        
        if ret.DomainName == "" {
            t.Errorf("Incorrect GetEmailForwarding response! Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}
