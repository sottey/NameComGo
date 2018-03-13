package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func ListDNSSECsExample() {
    fmt.Println("Executing ListDNSSECs() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")
    
    // Test data
    domainName := "any1.org"
    
    // Make the call
    nameResponse := namecomgo.ListDNSSECs(c, domainName)
    
    // Convert to bytes
    resp := []byte(nameResponse.Body)

    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create a response object
            var ret namecomgo.ListDNSSECsResponse
            
            // Hydrate response object with returned data
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in ListDNSSECs(): ", jsonErr)
            
            // Display results
            if len(ret.DNSSECData) == 0 {
                fmt.Printf("No records for %v\n", domainName)
            } else {
                for _, dnssec := range ret.DNSSECData {
                    fmt.Printf("\nDomain: %v\nKeyTag: %v\nAlgo: %v\nDigestType: %v\nDigest: %v\n\n", dnssec.DomainName, dnssec.KeyTag, dnssec.Algorithm, dnssec.DigestType, dnssec.Digest)
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