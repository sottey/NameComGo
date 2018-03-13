package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func SetNameServersExample() {
    fmt.Println("Executing SetNameServers() example...\n")

    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any99.org"
    nameservers := []string{"ns1.name.com", "ns2.name.com"}

    // Make the call
    nameResponse := namecomgo.SetNameServers(c, domainName, nameservers)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.Domain
            
            // Hydrate response object with result data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in SetNameServers(): ", jsonErr)
            
            // Display results
            if ret.DomainName != "" {
                fmt.Printf("SetNameServers Returned for domain %v\n", ret.DomainName)
            } else {
                fmt.Printf("Empty SetNameServers response!")
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }
}