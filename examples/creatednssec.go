package examples

import (
    "encoding/json"
    "fmt"
    namecomgo "namecomgo/nameapi"
)

func CreateDNSSECExample() {
    fmt.Println("Executing CreateDNSSEC() example...\n")

    // Get the configuration details (set in config.go... username, token, etc.)
    c  := namecomgo.GetConfig("test")

    // Test data
    dnsSEC := namecomgo.CreateDNSSECObject("any99.org", 30909, 8, 2, "E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766")
    
    
    // Make the call
    nameResponse := namecomgo.CreateDNSSEC(c, dnsSEC)

    // Convert to bytes
    resp := []byte(nameResponse.Body)
    
    // Handle by statuscode from server
    switch nameResponse.StatusCode {
    	case 200:
    	    // Create response objectr
            var ret namecomgo.DNSSEC
            
            // Hydrate response object with results
            var jsonErr = json.Unmarshal(resp, &ret)
            namecomgo.Check("Error in CreateDNSSEC(): ", jsonErr)
            
            // Display results
            if ret.DomainName != "" {
                fmt.Printf("\nDomain: %v\nKeyTag: %v\nAlgo: %v\nDigestType: %v\nDigest: %v\n\n", ret.DomainName, ret.KeyTag, ret.Algorithm, ret.DigestType, ret.Digest)
            } else {
                fmt.Printf("Empty DNSSEC response!")
            }
    	case 404:
    		fmt.Printf("\n404 Status Code... Not Found.\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	case 500:
    		fmt.Printf("\n500 Status Code... Internal Error\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    	default:
    		fmt.Printf("\nUnhandled Status Code\n Status: %v, Message: %v", nameResponse.StatusCode, nameResponse.Body)
    }

}