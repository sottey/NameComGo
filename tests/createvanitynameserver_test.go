package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestCreateVanityNameserver(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    vanityServer := nameapi.CreateVanityNameserverObject("any1.org", "ww2.any1.org", []string {"123.1.1.10", "123.1.1.11"})
    nameResponse := nameapi.CreateVanityNameserver(c,vanityServer)
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.VanityNameserver
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in CreateVanityNameserver(): ", jsonErr)
        
        if ret.DomainName == "" {
            t.Errorf("\nNothing returned. Expected values, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
        }
    }
}