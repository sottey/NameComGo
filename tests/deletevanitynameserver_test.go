package apitests

import (
    "../nameapi"
    "testing"
)

func TestDeleteVanityNameserver(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    vanityServer := nameapi.CreateVanityNameserverObject("any1.org", "ww2.any1.org", []string {"123.1.1.10", "123.1.1.11"})
    nameResponse := nameapi.DeleteVanityNameserver(c,vanityServer)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}