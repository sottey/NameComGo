package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestGetRecord(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.GetRecord(c, "any1.org", 334094)
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode !=200{
        t.Errorf("Bad response from GetRecord(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.Record
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in GetRecord(): ", jsonErr)
    }
}