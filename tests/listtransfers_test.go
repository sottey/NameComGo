package apitests

import (
    "../nameapi"
    "testing"
)

func TestListTransfers(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.ListTransfers(c,)
    
    if nameResponse.StatusCode !=200 {
        t.Errorf("Bad response from ListTransfers(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}