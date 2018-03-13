package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func GetEmailForwardingExample() {
    fmt.Println("Executing GetEmailForwarding() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any99.org"
    emailBox := "admin"
    
    // Make the call
    nameResponse := namecomgo.GetEmailForwarding(c, domainName, emailBox)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)

    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create response object
            var ret namecomgo.EmailForwarding
            
            // Hydrate the response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in GetEmailForwarding(): ", jsonErr)
            
            // Display results
            if ret.DomainName != "" {
                fmt.Printf("%v@%v forwards to %v\n", ret.EmailBox, ret.DomainName, ret.EmailTo)
            } else {
                fmt.Printf("Empty GetEmailForwarding response!")
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}