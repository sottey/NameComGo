package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestCreateRecord(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.CreateRecord(c, nameapi.CreateRecordObject(0, "any1.org", nameapi.RandString(6), "A", "10.0.0.11", 300))
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode !=200{
        t.Errorf("Bad response from CreateRecord(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.ListRecordsResponse
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in CreateRecords(): ", jsonErr)
    }
}