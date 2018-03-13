package examples

import (
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func DeleteEmailForwardingExample() {
    fmt.Println("Executing DeleteEmailForwarding() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any99.org"
    emailBox := "admin"
    
    // Make the call
    nameResponse := namecomgo.DeleteEmailForwarding(c, domainName, emailBox)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
            fmt.Printf("DeleteEmailForwarding Returned for domain %v\n", domainName)
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}