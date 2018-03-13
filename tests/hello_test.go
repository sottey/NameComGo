package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestHello(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    var ret nameapi.HelloResponse
    nameResponse:= nameapi.Hello(c)
    resp := []byte(nameResponse.Body)
    
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in TestHello(): ", jsonErr)
    
    if nameResponse.StatusCode != 200 || ret.ServerName == "" {
        t.Errorf("Bad response from Hello(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}