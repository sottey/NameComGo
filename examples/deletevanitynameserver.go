package examples

import (
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func DeleteVanityNameserverExample() {
    fmt.Println("Executing DeleteVanityNameserver() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    c.Debug = false
    
    // Test data
    vanityServer := namecomgo.CreateVanityNameserverObject("any99.org", "www.any99.org", []string {"123.0.0.1", "123.0.0.2"})
    
    
    // Make the call
    nameResponse := namecomgo.DeleteVanityNameserver(c, vanityServer)

    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    fmt.Printf("Vanity Nameserver has been deleted\n")
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }

}