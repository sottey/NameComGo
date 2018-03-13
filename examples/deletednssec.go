package examples

import (
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func DeleteDNSSECExample() {
    fmt.Println("Executing DeleteDNSSEC() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any99.org"
    digest := "E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766"
    
    // Make the call
    nameResponse := namecomgo.DeleteDNSSEC(c, domainName, digest)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    fmt.Println("Successfully deleted.")
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}