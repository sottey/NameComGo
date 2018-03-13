package examples

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func UpdateEmailForwardingExample(t *testing.T) {
    c  := nameapi.GetConfig("test")
    
    // Test data
    domainName := "any99.org"
    emailBox := "admin"
    emailTo := "admin@any1.org"
    
    // Make the call
    nameResponse := nameapi.UpdateEmailForwarding(c, domainName, emailBox, emailTo)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    var ret nameapi.EmailForwarding
    
    // Hydrate the response obect with returned data
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in UpdateEmailForwarding(): ", jsonErr)
    
    // Display results
    if nameResponse.StatusCode != 200 || ret.DomainName != "" {
        t.Errorf("UpdateEmailForwarding Returned for domain %v\n", ret.DomainName)
    }
}