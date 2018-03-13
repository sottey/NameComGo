package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)


func TestSearch(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    nameResponse := nameapi.Search(c, "sean", nil)
    resp := []byte(nameResponse.Body)
    
    var ret nameapi.SearchResults
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in Search(): ", jsonErr)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("Bad response from Search(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}