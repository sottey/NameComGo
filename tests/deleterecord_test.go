package apitests

import (
    "../nameapi"
    "testing"
)

func TestDeleteRecord(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.DeleteRecord(c, "any1.org", 334363)
    
    if nameResponse.StatusCode != 200{
        t.Errorf("Bad response from DeleteRecord(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}