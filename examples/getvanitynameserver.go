package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func GetVanityNameserverExample() {
    fmt.Println("Executing GetVanityNameserver() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    c.Debug = false
    
    // Test data
    domainName := "any1.org"
    hostname := "ww2.any1.org"
    
    // Make the call
    nameResponse := namecomgo.GetVanityNameserver(c, domainName, hostname)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)

    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.VanityNameserver
            
            // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in GetVanityNameserver(): ", jsonErr)
            
            // Display results
            if ret.DomainName == "" {
                fmt.Printf("No vanity nameserver for %v\n", hostname)
            } else {
                fmt.Printf("\nDomain: %v\nHostName:%v\nIP Count:%v\n\n", ret.DomainName, ret.Hostname, len(ret.IPs))
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}