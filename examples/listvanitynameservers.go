package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func ListVanityNameserversExample() {
    fmt.Println("Executing ListVanityNameservers() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    c.Debug = false
    
    // Test data
    domainName := "any1.org"
    
    // Make the call
    nameResponse := namecomgo.ListVanityNameservers(c, domainName)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)

    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.ListVanityNameserversResponse
            
            // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in ListVanityNameservers(): ", jsonErr)
            
            // Display results
            if len(ret.VanityServerData) == 0 {
                fmt.Printf("No vanity nameservers for %v\n", domainName)
            } else {
                for _, vanityserver := range ret.VanityServerData {
                    fmt.Printf("\nDomain: %v\nHostName:%v\n\n", vanityserver.DomainName, vanityserver.Hostname)
                    for _, ip := range vanityserver.IPs {
                        fmt.Printf("\nIP: %v\n\n", ip)
                    }
                }
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}