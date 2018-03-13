package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func CreateDomainExample() {
    fmt.Println("Executing CreateDomain() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    purchaseDetails := namecomgo.CreatePuchaseObject("any99.org", 1, 12.99)
    
    // Make the call
    nameResponse := namecomgo.CreateDomain(c, purchaseDetails)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create response object
            var ret namecomgo.PurchaseDetails
            
            // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in CreateDomain(): ", jsonErr)
            
            // Display results
            fmt.Printf("\nDomain: %v\nPrivacyEnabled:%v\nLocked:%v\nAutoRenew:%v\nExpireDate:%v\nCreateDate:%v\nRenewalPrice:%v\n\n", ret.DomainObj.DomainName,ret.DomainObj.PrivacyEnabled, ret.DomainObj.Locked, ret.DomainObj.AutorenewEnabled, ret.DomainObj.ExpireDate, ret.DomainObj.CreateDate,ret.DomainObj.RenewalPrice)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}
