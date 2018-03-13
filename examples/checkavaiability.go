package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func CheckAvailabilityExample() {
    fmt.Println("Executing CheckAvailability() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // test data, domains to check
    domains := []string {"amazon.org", "namecomgodomain.live"}
    
    // Make the call
    nameResponse := namecomgo.CheckAvailability(c, domains)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.SearchResults
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in CheckAvailability(): ", jsonErr)
        
            // Hydrate the response object with the response data
            for _, domain := range ret.Responses {
                fmt.Printf("\nDomain: %v\nSLD:%v\nTLD:%v\nPurchasable:%v\nPremium:%v\nPurchasePrice:%v\nPurchaseType:%v\nRenewalPrice:%v\n\n", domain.DomainName,domain.Sld, domain.Tld, domain.Purchasable, domain.Premium, domain.PurchasePrice, domain.PurchaseType,domain.RenewalPrice)
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}