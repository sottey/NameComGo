package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)


func TestGetAuthCodeForDomain(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    var ret nameapi.GetAuthCodeForDomainResponse
    nameResponse:= nameapi.GetAuthCodeForDomain(c, "asdfhlasjhdfjasjdkf.com")
    resp := []byte(nameResponse.Body)
    
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in GetAuthCodeForDomain(): ", jsonErr)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("Bad response from GetAuthCodeForDomain(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}
