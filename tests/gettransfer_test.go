package examples

import (
    "../nameapi"
    "testing"
)

func TestGetTransfer(t *testing.T) {
    c  := nameapi.GetConfig("test")
    
    // Test data
    domainName := "any4.org"
    
    // Make the call
    nameResponse := nameapi.GetTransfer(c, domainName)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}