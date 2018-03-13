package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func UpdateEmailForwardingExample() {
    fmt.Println("Executing UpdateEmailForwarding() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any99.org"
    emailBox := "admin"
    emailTo := "admin@any1.org"
    
    // Make the call
    nameResponse := namecomgo.UpdateEmailForwarding(c, domainName, emailBox, emailTo)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.EmailForwarding
            
            // Hydrate the response obect with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in UpdateEmailForwarding(): ", jsonErr)
            
            // Display results
            if ret.DomainName != "" {
                fmt.Printf("UpdateEmailForwarding Returned for domain %v\n", ret.DomainName)
            } else {
                fmt.Printf("Empty UpdateEmailForwarding response!")
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}