package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func ListDomainsExample() {
    fmt.Println("Executing ListDomains() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    c.Debug = true
    
    // Create a response object
    var ret namecomgo.ListDomainsResponse
    
    // Make the call
    nameResponse := namecomgo.ListDomains(c)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
            // Hydrate the response object with the response data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in ListDomains(): ", jsonErr)
            
            // Display member variables from the response object
            for _, domain := range ret.Domains {
                
                // NOTE: Renewal and PrivacyEnabled are never returned from the API
                fmt.Printf("\nDomain: %v\nPrivacyEnabled:%v\nLocked:%v\nAutoRenew:%v\nExpireDate:%v\nCreateDate:%v\nRenewalPrice:%v\n\n", domain.DomainName,domain.PrivacyEnabled, domain.Locked, domain.AutorenewEnabled, domain.ExpireDate, domain.CreateDate,domain.RenewalPrice)
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}
