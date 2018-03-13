package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestCreateDomain(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    domainName := nameapi.RandString(15) + ".org"
    
    nameResponse := nameapi.CreateDomain(c,nameapi.CreatePuchaseObject(domainName, 1, 12.99))
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.Domain
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in CreateDomain(): ", jsonErr)
    }
}