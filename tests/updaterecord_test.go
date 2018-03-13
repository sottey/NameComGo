package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestUpdateRecord(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.UpdateRecord(c, nameapi.CreateRecordObject(334094, "any1.org", nameapi.RandString(6), "A", "10.0.0.12", 300))
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode !=200{
        t.Errorf("Bad response from UpdateRecord(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.Record
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in UpdateRecords(): ", jsonErr)
    }
}