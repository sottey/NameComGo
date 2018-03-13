package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func CreateVanityNameserverExample() {
    fmt.Println("Executing CreateVanityNameserver() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    c.Debug = false
    
    // Test data
    vanityServer := namecomgo.CreateVanityNameserverObject("any99.org", "www.any99.org", []string {"123.0.0.1", "123.0.0.2"})
    
    
    // Make the call
    nameResponse := namecomgo.CreateVanityNameserver(c, vanityServer)

    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create response objectr
            var ret namecomgo.VanityNameserver
            
            // Hydrate response object with results
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in CreateVanityNameserver(): ", jsonErr)
            
            // Display results
            if ret.DomainName != "" {
                fmt.Printf("\nDomain: %v\nHostName: %v\nIP Count: %v\n\n", ret.DomainName, ret.Hostname, len(ret.IPs))
            } else {
                fmt.Printf("Empty CreateVanityNameserver response!")
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }

}