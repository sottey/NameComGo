package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func UpdateURLForwardingExample() {
    fmt.Println("Executing UpdateURLForwarding() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // TestData
    urlForwarding := namecomgo.CreateURLForwardingObject("any1.org", "www.any1.org", "http://test2.org", "redirect")
    
    // Make the call
    nameResponse := namecomgo.UpdateURLForwarding(c, urlForwarding)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
        	case 200:
        	// Create return object
            var ret namecomgo.URLForwarding
            
            // Hydrate object with return data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in UpdateURLForwardings(): ", jsonErr)
    
            // Display data
            fmt.Printf("\nDomainName:%v\nForwardsTo:%v", ret.Host, ret.ForwardsTo)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}