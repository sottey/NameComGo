package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func ListEmailForwardingsExample() {
    fmt.Println("Executing ListEmailForwardings() example...\n")
    
    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any99.org"
    
    // Make the call
    nameResponse := namecomgo.ListEmailForwardings(c, domainName)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)

    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create response object
            var ret namecomgo.ListEmailForwardingsResponse
            
            // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in ListEmailForwardings(): ", jsonErr)
            
            // Display results
            if len(ret.EmailForwarding) > 0 {
                fmt.Printf("ListEmailForwardings Returned:")
                for _, ef := range ret.EmailForwarding {
                    fmt.Printf("%v@%v forwards to %v\n", ef.EmailBox, ef.DomainName, ef.EmailTo)
                }
            } else {
                fmt.Printf("Empty ListEmailForwardings response!\n")
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}