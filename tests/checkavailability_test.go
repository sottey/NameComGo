package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestCheckAvailability(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
 
    nameResponse := nameapi.CheckAvailability(c, []string {"any2.org", "asdfjhasdljfhasjdfhkjlasdhfjlkahsdf.live"} )
    resp := []byte(nameResponse.Body)
    
    var ret nameapi.SearchResults
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in CheckAvailability(): ", jsonErr)

    if nameResponse.StatusCode != 200 {
        t.Errorf("Bad response from CheckAvailability(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}