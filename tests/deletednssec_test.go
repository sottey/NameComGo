package apitests

import (
    "../nameapi"
    "testing"
)

func TestDeleteDNSSEC(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.DeleteDNSSEC(c,"any2.org", "E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766")
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("\nNon 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}