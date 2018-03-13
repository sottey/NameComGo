package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func GetURLForwardingExample() {
    fmt.Println("Executing GetURLForwarding() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    
    // Test data
    domainName := "any2.org"
    host := "gahfoo.any2.org"
    
    // Make the call
    nameResponse := namecomgo.GetURLForwarding(c, domainName, host)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.URLForwarding
            
            // Hydrate the response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in GetURLForwarding(): ", jsonErr)
        
            // Display results
            fmt.Printf("\nDomainName:%v\nForwardsTo:%v\n\n", ret.Host, ret.ForwardsTo)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}
