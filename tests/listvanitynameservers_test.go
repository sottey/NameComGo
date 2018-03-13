package examples

import (
    "../nameapi"
    "testing"
)

func TestListVanityNameservers(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    // Test data
    domainName := "any1.org"
    
    // Make the call
    nameResponse := nameapi.ListVanityNameservers(c, domainName)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("Bad response from ListVanityNameservers(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}