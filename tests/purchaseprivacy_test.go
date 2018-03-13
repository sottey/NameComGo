package examples

import (
    "../nameapi"
    "testing"
    "encoding/json"
)

func TestPurchasePrivacy(t *testing.T) {
    c  := nameapi.GetConfig("test")
    
    // Test Data
    domainName := "any4.org"
    price := 4.99
    var years int32 = 1

    nameResponse := nameapi.PurchasePrivacy(c, domainName, price, years)
    resp := []byte(nameResponse.Body)
    
    var ret nameapi.PurchasePrivacyResponse
    var jsonErr = json.Unmarshal(resp, &ret)
    nameapi.Check("Error in PurchasePrivacy(): ", jsonErr)
    
    if nameResponse.StatusCode != 200 || ret.DomainData.DomainName == "" {
        t.Errorf("Bad response from PurchasePrivacy(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}