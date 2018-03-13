package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestListRecords(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.ListRecords(c, "any1.org")
    resp := []byte(nameResponse.Body)
    
    if nameResponse.StatusCode !=200 {
        t.Errorf("Bad response from CheckAvailability(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    } else {
        var ret nameapi.ListRecordsResponse
        var jsonErr = json.Unmarshal(resp, &ret)
        nameapi.Check("Error in ListRecords(): ", jsonErr)
    
        if len(ret.Records) < 1 {
            t.Errorf("No responses from ListRecords(). BodyText:%v", nameResponse.Body)
        }
    }
}