package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func CreateURLForwardingExample() {
    fmt.Println("Executing URLForwarding() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")

    // Creating the object used in the call via a utility class
    urlForwarding := namecomgo.CreateURLForwardingObject("any1.org",namecomgo.RandString(6) + ".any1.org", "http://www.test.org", "redirect")
    
    // Make the call
    nameResponse := namecomgo.CreateURLForwarding(c, urlForwarding)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create return object
            var ret namecomgo.URLForwarding
            
            // Hydrate return object with return data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in CreateURLForwardings(): ", jsonErr)
            
            // Display result
            fmt.Printf("\nDomainName:%v\nForwardsTo:%v\n", ret.DomainName, ret.ForwardsTo)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}