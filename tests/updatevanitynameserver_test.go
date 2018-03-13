package examples

import (
    "../nameapi"
    "testing"
)

func TestUpdateVanityNameserver(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    vanityServer := nameapi.CreateVanityNameserverObject("any99.org", "www.any99.org", []string {"123.0.0.1", "123.0.0.2"})
    nameResponse := nameapi.UpdateVanityNameserver(c, vanityServer)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("Bad response from UpdateVanityNameserver(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }

}