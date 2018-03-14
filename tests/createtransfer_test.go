package apitests

import (
    "../nameapi"
    "testing"
)

func TestCreateTransfer(t *testing.T) {
    c  := nameapi.GetConfig("test")
    c.Debug = false
    
    // Test data
    domain := "any99.org"
    authcode := "j5Y4U3f4S0N"
    purchasePrice := 12.99
    privacy := true
    promoCode := "seano"
    
    nameResponse := nameapi.CreateTransfer(c, domain, authcode, promoCode, privacy, purchasePrice)
    
    if nameResponse.StatusCode !=200{
        t.Errorf("Bad response from CreateTransfer(): Non 200 Status Code, Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}