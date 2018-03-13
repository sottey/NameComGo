package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func EnableAutoRenewExample() {
    fmt.Println("Executing EnableAutorenew() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "Any2.org"
    
    // Make the call
    nameResponse := namecomgo.EnableAutoRenew(c, domainName)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
            // Create return object
            var ret namecomgo.Domain
            
            // Hydrate return object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in DisableAutoRenew(): ", jsonErr)
    
            // Display returned data
            fmt.Printf("\nDomainName:%v\nAutoRenew:%v\n\n", ret.DomainName, ret.AutorenewEnabled)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}