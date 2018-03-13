package examples

import (
    "../nameapi"
    "testing"
)

func TestGetVanityNameserver(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    // Test data
    domainName := "any1.org"
    hostname := "ww2.any1.org"
    
    // Make the call
    nameResponse := nameapi.GetVanityNameserver(c, domainName, hostname)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("Bad response from GetVanityNameserver(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}