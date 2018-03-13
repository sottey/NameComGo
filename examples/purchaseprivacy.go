package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func PurchasePrivacyExample() {
    fmt.Println("Executing PurchasePrivacy() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test Data
    domainName := "any4.org"
    price := 4.99
    var years int32 = 1
    
    // Make the call
    nameResponse := namecomgo.PurchasePrivacy(c, domainName, price, years)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create response object
            var ret namecomgo.PurchasePrivacyResponse
            
            // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in PurchasePrivacy(): ", jsonErr)
            
            // Display results
            fmt.Printf("\nPrivacy Purchased for: %v\nPrivacy:%v", ret.DomainData.DomainName, ret.DomainData.PrivacyEnabled)
        case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}