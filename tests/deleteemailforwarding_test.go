package apitests

import (
    "../nameapi"
    "testing"
)

func TestDeleteEmailForwarding(t *testing.T) {
    c  := nameapi.GetConfig("test")
    
    // Test data
    domainName := "any99.org"
    emailBox := "admin"
    
    nameResponse := nameapi.DeleteEmailForwarding(c, domainName, emailBox)
    
    if nameResponse.StatusCode != 200 {
       t.Errorf("Bad response from DeleteEmailForwarding(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}