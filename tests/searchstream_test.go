package apitests

import (
    "../nameapi"
    "testing"
    "encoding/json"
)


func TestSearchStream(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    // Test data
    keyword:= "emmabowbemma"
    tldList := []string {"com", "net", "org", "live", "club", "news"}
    
    // Create a response object
    nameResponse := nameapi.Search(c, keyword, tldList, 5000, "seano")
    resp := []byte(nameResponse.Body)
    
    var ret nameapi.SearchResults
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in Search(): ", jsonErr)
    
    if nameResponse.StatusCode != 200 {
        t.Errorf("Bad response from Search(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}